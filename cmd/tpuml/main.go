// Command tpuml converts state machine documents: it parses an input file
// into a format-neutral model and renders that model either through a Go
// template or with a built-in emitter.
package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/vincedupuis/transplantUML/assets"
	"github.com/vincedupuis/transplantUML/internal/format"
	"github.com/vincedupuis/transplantUML/internal/model"
	"github.com/vincedupuis/transplantUML/internal/render"
)

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "tpuml:", err)
		os.Exit(1)
	}
}

func run(args []string, stdout, stderr io.Writer) error {
	cmd := newRootCmd(stdout, stderr)
	if len(args) == 0 {
		// Bare "tpuml": show the usage rather than complaining about -i.
		args = []string{"--help"}
	}
	cmd.SetArgs(args)
	return cmd.Execute()
}

type options struct {
	input        string
	inputFormat  string
	tmplFile     string
	outputFormat string
	output       string
}

func newRootCmd(stdout, stderr io.Writer) *cobra.Command {
	var opts options
	cmd := &cobra.Command{
		Use:   "tpuml -i input [-f format] [-t template.tmpl | -F format] [-o output]",
		Short: "Convert state machine documents",
		Long: "tpuml parses a state machine document into a format-neutral model and renders\n" +
			"that model either through a Go template (PlantUML by default) or with a\n" +
			"built-in emitter.",
		Args:                  cobra.NoArgs,
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		SilenceErrors:         true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return convert(opts, stdout, stderr)
		},
	}
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)

	f := cmd.Flags()
	f.StringVarP(&opts.input, "input", "i", "", "input file (required)")
	f.StringVarP(&opts.inputFormat, "input-format", "f", "", "input format, one of: "+strings.Join(format.ParserNames(), ", ")+" (default: from the input file extension)")
	f.StringVarP(&opts.tmplFile, "template", "t", "", "Go template file to render the model with (default: built-in PlantUML template)")
	f.StringVarP(&opts.outputFormat, "output-format", "F", "", "emit a built-in output format instead of a template, one of: "+strings.Join(format.EmitterNames(), ", "))
	f.StringVarP(&opts.output, "output", "o", "", "output file (default: stdout)")

	cmd.MarkFlagRequired("input")
	cmd.MarkFlagsMutuallyExclusive("template", "output-format")
	return cmd
}

func convert(opts options, stdout, stderr io.Writer) error {
	name := opts.inputFormat
	if name == "" {
		if name = format.Detect(opts.input); name == "" {
			return fmt.Errorf("cannot infer the format of %q; use -f", opts.input)
		}
	}
	parser, err := format.ParserFor(name)
	if err != nil {
		return err
	}
	src, err := os.ReadFile(opts.input)
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}
	sm, warnings, err := parser.Parse(src)
	if err != nil {
		return fmt.Errorf("%s: %w", opts.input, err)
	}
	report(stderr, warnings)
	if err := sm.Validate(); err != nil {
		return fmt.Errorf("%s: invalid state machine:\n%w", opts.input, err)
	}

	var out []byte
	switch {
	case opts.outputFormat != "":
		emitter, err := format.EmitterFor(opts.outputFormat)
		if err != nil {
			return err
		}
		if out, warnings, err = emitter.Emit(sm); err != nil {
			return err
		}
	default:
		tmpl := assets.PlantUML
		if opts.tmplFile != "" {
			data, err := os.ReadFile(opts.tmplFile)
			if err != nil {
				return fmt.Errorf("reading template: %w", err)
			}
			tmpl = string(data)
		}
		text, w, err := render.Render(sm, tmpl)
		if err != nil {
			return err
		}
		out, warnings = []byte(text), w
	}
	report(stderr, warnings)

	if opts.output == "" {
		_, err = stdout.Write(out)
		return err
	}
	if err := os.WriteFile(opts.output, out, 0o644); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}
	fmt.Fprintln(stderr, "output written to", opts.output)
	return nil
}

// report prints the warnings a parser, emitter or template raised: what the
// input held that the model or the output has no place for. They never fail
// the conversion.
func report(stderr io.Writer, warnings model.Warnings) {
	for _, w := range warnings {
		fmt.Fprintln(stderr, "tpuml: warning:", w)
	}
}
