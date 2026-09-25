// Command fsm converts state machine documents: it parses an input file
// into a format-neutral model and renders that model either through a Go
// template or with a built-in emitter.
package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/vincedupuis/transplantUML/assets"
	"github.com/vincedupuis/transplantUML/internal/format"
	"github.com/vincedupuis/transplantUML/internal/model"
	"github.com/vincedupuis/transplantUML/internal/render"
)

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "fsm:", err)
		os.Exit(1)
	}
}

func run(args []string, stdout, stderr io.Writer) error {
	cmd := newRootCmd(stdout, stderr)
	if len(args) == 0 {
		// Bare "fsm": show the usage rather than complaining about -i.
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
		Use:   "fsm -i input [-f format] [-t template | -F format] [-o output]",
		Short: "Convert state machine documents",
		Long: "fsm parses a state machine document into a format-neutral model and renders\n" +
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
	f.StringVarP(&opts.tmplFile, "template", "t", "", "Go template file to render the model with, or the name of a built-in one: "+strings.Join(assets.TemplateNames(), ", ")+" (default: puml)")
	f.StringVarP(&opts.outputFormat, "output-format", "F", "", "emit a built-in output format instead of a template, one of: "+strings.Join(format.EmitterNames(), ", "))
	f.StringVarP(&opts.output, "output", "o", "", "output file, or the folder for a template that writes several files (default: stdout)")

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
		tmpl, err := template(opts.tmplFile)
		if err != nil {
			return err
		}
		text, w, err := render.Render(sm, tmpl)
		if err != nil {
			return err
		}
		report(stderr, w)
		files, err := render.Files(text)
		if err != nil {
			return err
		}
		if files != nil {
			return writeFiles(opts.output, files, stderr)
		}
		out, warnings = []byte(text), nil
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

// writeFiles writes the files a template produced into dir, creating it.
func writeFiles(dir string, files []render.File, stderr io.Writer) error {
	if dir == "" {
		return fmt.Errorf("the template writes %d files; name a folder for them with -o", len(files))
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}
	for _, f := range files {
		path := filepath.Join(dir, f.Name)
		if err := os.WriteFile(path, []byte(f.Content), 0o644); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
		fmt.Fprintln(stderr, "output written to", path)
	}
	return nil
}

// template returns the template -t names: a file when one exists at that
// path, a built-in template otherwise, PlantUML when -t is not given.
func template(name string) (string, error) {
	if name == "" {
		return assets.PlantUML, nil
	}
	data, err := os.ReadFile(name)
	if err == nil {
		return string(data), nil
	}
	if builtin, ok := assets.Templates[name]; ok && errors.Is(err, fs.ErrNotExist) {
		return builtin, nil
	}
	return "", fmt.Errorf("reading template: %w (built-in templates: %s)", err, strings.Join(assets.TemplateNames(), ", "))
}

// report prints the warnings a parser, emitter or template raised: what the
// input held that the model or the output has no place for. They never fail
// the conversion.
func report(stderr io.Writer, warnings model.Warnings) {
	for _, w := range warnings {
		fmt.Fprintln(stderr, "fsm: warning:", w)
	}
}
