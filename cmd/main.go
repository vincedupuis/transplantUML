package main

import (
	"flag"
	"fmt"
	"os"
	"transplant-scxml/internal"
)

func main() {
	inputFileName := flag.String("i", "input.scxml", "Input SCXML filename")
	tmplFileName := flag.String("t", "template.tmpl", "Template filename")
	outputFileName := flag.String("o", "output.txt", "Output filename")
	flag.Parse()

	scxml := readFile(inputFileName)
	stateMachine := convertScxmlToStateMachineModel(scxml)

	tmplFile := readFile(tmplFileName)
	output := runTemplate(stateMachine, tmplFile)

	writeFile(outputFileName, output)

	fmt.Printf("Output file successfully written to: %s\n", *outputFileName)
}

func convertScxmlToStateMachineModel(xml string) *internal.StateMachine {
	stateMachine, err := internal.ScxmlToStateMachine(xml)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error extracting SCXML from file: %s", err)
		os.Exit(1)
	}
	return stateMachine
}

func runTemplate(sm *internal.StateMachine, tmplFile string) string {
	content, err := internal.RunTemplate(sm, tmplFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error running template: %s", err)
		os.Exit(1)
	}
	return content
}

func readFile(filename *string) string {
	content, err := os.ReadFile(*filename)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error reading file: %s", err)
		os.Exit(1)
	}
	return string(content)
}

func writeFile(outputFileName *string, content string) {
	err := os.WriteFile(*outputFileName, []byte(content), 0644)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error writing file: %s", err)
		os.Exit(1)
	}
}