// Command tpuml converts state machine documents: it parses an input file
// into a format-neutral model and renders that model either through a Go
// template or with a built-in emitter.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/vincedupuis/transplantUML/assets"
	"github.com/vincedupuis/transplantUML/internal/format"
	"github.com/vincedupuis/transplantUML/internal/render"
)

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			fmt.Fprintln(os.Stderr, "tpuml:", err)
		}
		os.Exit(1)
	}
}

func run(args []string, stdout, stderr io.Writer) error {
	fs := flag.NewFlagSet("tpuml", flag.ContinueOnError)
	fs.SetOutput(stderr)
	input := fs.String("i", "", "input file (required)")
	inputFormat := fs.String("f", "", "input format, one of: "+strings.Join(format.ParserNames(), ", ")+" (default: from the input file extension)")
	tmplFile := fs.String("t", "", "Go template file to render the model with (default: built-in PlantUML template)")
	outputFormat := fs.String("F", "", "emit a built-in output format instead of a template, one of: "+strings.Join(format.EmitterNames(), ", "))
	output := fs.String("o", "", "output file (default: stdout)")
	export := fs.String("e", "", "also write the parsed model as JSON to this file")
	fs.Usage = func() {
		fmt.Fprintln(stderr, "usage: tpuml -i input [-f format] [-t template.tmpl | -F format] [-o output] [-e model.json]")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *input == "" {
		fs.Usage()
		return errors.New("-i is required")
	}
	if *tmplFile != "" && *outputFormat != "" {
		return errors.New("-t and -F are mutually exclusive")
	}

	name := *inputFormat
	if name == "" {
		if name = format.Detect(*input); name == "" {
			return fmt.Errorf("cannot infer the format of %q; use -f", *input)
		}
	}
	parser, err := format.ParserFor(name)
	if err != nil {
		return err
	}
	src, err := os.ReadFile(*input)
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}
	sm, err := parser.Parse(src)
	if err != nil {
		return fmt.Errorf("%s: %w", *input, err)
	}
	if err := sm.Validate(); err != nil {
		return fmt.Errorf("%s: invalid state machine:\n%w", *input, err)
	}

	if *export != "" {
		emitter, _ := format.EmitterFor("json")
		data, err := emitter.Emit(sm)
		if err != nil {
			return err
		}
		if err := os.WriteFile(*export, data, 0o644); err != nil {
			return fmt.Errorf("writing export: %w", err)
		}
		fmt.Fprintln(stderr, "model exported to", *export)
	}

	var out []byte
	switch {
	case *outputFormat != "":
		emitter, err := format.EmitterFor(*outputFormat)
		if err != nil {
			return err
		}
		if out, err = emitter.Emit(sm); err != nil {
			return err
		}
	default:
		tmpl := assets.PlantUML
		if *tmplFile != "" {
			data, err := os.ReadFile(*tmplFile)
			if err != nil {
				return fmt.Errorf("reading template: %w", err)
			}
			tmpl = string(data)
		}
		text, err := render.Render(sm, tmpl)
		if err != nil {
			return err
		}
		out = []byte(text)
	}

	if *output == "" {
		_, err = stdout.Write(out)
		return err
	}
	if err := os.WriteFile(*output, out, 0o644); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}
	fmt.Fprintln(stderr, "output written to", *output)
	return nil
}
