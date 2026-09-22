// Package jsonsm reads and writes the model as JSON. The JSON shape is the
// model itself (see the json struct tags in package model), so it doubles as
// a lossless interchange format and as a debugging view of what a parser
// produced.
package jsonsm

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/vincedupuis/transplantUML/internal/model"
)

type Parser struct{}

// Parse never warns: the JSON shape is the model, so it holds everything.
func (Parser) Parse(src []byte) (*model.StateMachine, model.Warnings, error) {
	var sm model.StateMachine
	dec := json.NewDecoder(bytes.NewReader(src))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&sm); err != nil {
		return nil, nil, fmt.Errorf("parsing JSON: %w", err)
	}
	// Hand-written JSON may leave these out; keep the model uniform.
	if sm.States == nil {
		sm.States = make([]*model.State, 0)
	}
	if sm.Transitions == nil {
		sm.Transitions = make([]*model.Transition, 0)
	}
	for _, s := range sm.States {
		if s.Kind == "" {
			s.Kind = model.Normal
		}
	}
	return &sm, nil, nil
}

type Emitter struct{}

// Emit never warns: JSON holds the whole model.
func (Emitter) Emit(sm *model.StateMachine) ([]byte, model.Warnings, error) {
	out, err := json.MarshalIndent(sm, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("encoding JSON: %w", err)
	}
	return append(out, '\n'), nil, nil
}
