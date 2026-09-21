// Package format is the registry of input parsers and built-in output emitters.
// Adding a format means implementing Parser and/or Emitter in its own package
// and registering it in the tables below.
package format

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/vincedupuis/transplantUML/internal/jsonsm"
	"github.com/vincedupuis/transplantUML/internal/model"
	"github.com/vincedupuis/transplantUML/internal/scxml"
)

// Parser turns a source document into the model.
type Parser interface {
	Parse(src []byte) (*model.StateMachine, error)
}

// Emitter serializes the model back into one of the supported document
// formats. Every format tpuml can read it can also write, so parsers and
// emitters come in pairs; free-form text output (diagrams, code, docs) is the
// template mechanism's job instead, see package render.
type Emitter interface {
	Emit(sm *model.StateMachine) ([]byte, error)
}

var parsers = map[string]Parser{
	"scxml": scxml.Parser{},
	"json":  jsonsm.Parser{},
}

var emitters = map[string]Emitter{
	"scxml": scxml.Emitter{},
	"json":  jsonsm.Emitter{},
}

// extensions maps a lower-case file extension to a parser name.
var extensions = map[string]string{
	".scxml": "scxml",
	".xml":   "scxml",
	".json":  "json",
}

func ParserFor(name string) (Parser, error) {
	if p, ok := parsers[name]; ok {
		return p, nil
	}
	return nil, fmt.Errorf("unknown input format %q (known: %s)", name, strings.Join(ParserNames(), ", "))
}

func EmitterFor(name string) (Emitter, error) {
	if e, ok := emitters[name]; ok {
		return e, nil
	}
	return nil, fmt.Errorf("unknown output format %q (known: %s)", name, strings.Join(EmitterNames(), ", "))
}

// Detect guesses the input format from a filename's extension; "" if unknown.
func Detect(filename string) string {
	return extensions[strings.ToLower(filepath.Ext(filename))]
}

func ParserNames() []string  { return sortedKeys(parsers) }
func EmitterNames() []string { return sortedKeys(emitters) }

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
