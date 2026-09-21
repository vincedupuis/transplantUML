# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

`tpuml` is a Go CLI that converts state machine documents. Any supported input format is parsed into one
format-neutral model (`internal/model`), which is then rendered either through a Go `text/template` (text outputs:
PlantUML, code, docs) or a built-in emitter (structured outputs: JSON). The bundled `assets/puml.tmpl` (embedded in
the binary, used when `-t` is omitted) produces PlantUML.

## Commands

```bash
make build                          # produces ./bin/tpuml (see the Makefile for run/test/fmt/vet/clean)
go build ./...
go build ./cmd/tpuml                # produces ./tpuml
go vet ./...
go test ./...
go test ./internal/scxml -run TestEdgeCases          # one test
go test ./internal/render -run TestPlantUMLGolden/edge

# End-to-end
./tpuml -i example/coffee-machine.scxml                 # PlantUML to stdout
./tpuml -i example/coffee-machine.scxml -F json         # model as JSON
```

Golden files for the PlantUML template live in `internal/render/testdata/*.puml`; regenerate one with
`go run ./cmd/tpuml -i <input> -o internal/render/testdata/<name>.puml` after checking the diff is intended.

## Architecture

Pipeline: `cmd/tpuml/main.go` (`run() error`) → `format.ParserFor(name).Parse` → `*model.StateMachine` →
`sm.Validate()` → `render.Render(sm, tmpl)` **or** `format.EmitterFor(name).Emit(sm)` → output.

- **`internal/model`** — the contract everything else depends on. Flat `States` with `Parent` links (top level is
  `""`), and flat `Transitions`. Pseudo-states are real `State`s distinguished by `Kind` (`normal`, `parallel`,
  `final`, `history-shallow`, `history-deep`); a history state's default transition is an ordinary `Transition`
  whose `Source` is the history state. `Initial` is a property (`StateMachine.Initial`, `State.Initial`), not a
  synthetic transition. The model must stay a *superset* of every input format — do not shape it around one
  template. `Validate()` is the single place structural rules live; extend it when the model grows.
- **`internal/format`** — `Parser` / `Emitter` interfaces plus the name→implementation tables and extension
  detection. To add a format: implement it in its own package under `internal/`, register it here, add its
  extension. Package names avoid stdlib clashes (`jsonsm`, not `json`).
- **`internal/scxml`** — recursive `etree` walk over `<state>/<parallel>/<final>/<history>` children (direct
  children only, so `<initial>`'s inner `<transition>` is not mistaken for a real transition). Executable content
  is flattened to strings; unknown elements are kept as raw XML.
- **`internal/jsonsm`** — the model's own JSON shape (struct tags in `model`). Parser uses
  `DisallowUnknownFields`; round-trip equality with the SCXML parser is tested.
- **`internal/render`** — registers sprig plus project helpers (`include`, `prefix`, `surround`, `joinNonEmpty`)
  and the model accessors as template functions. `joinNonEmpty` exists because sprig's `join` has the signature
  `join sep list`; don't shadow sprig names.
- **`assets/puml.tmpl`** — every emitted line starts with `\n` so nested blocks compose via
  `include ... | trimPrefix "\n" | indent 4`. Finals are drawn on incoming arrows as `[*]`, history as
  `parent[H]`/`parent[H*]`, parallel regions are the *bodies* of the region states separated by `--`. A compound
  state's own outgoing transitions are emitted after its closing `}` so PlantUML doesn't nest the targets.

## Conventions

- README.md documents the model, template functions and CLI flags; update it when any of those change.
- Test fixtures: `example/coffee-machine.scxml` (simple, flat) and `internal/scxml/testdata/edge.scxml`
  (parallel, `<initial>` element, deep history, final, onentry, multi-target). Add new SCXML features to `edge.scxml`
  and its golden.
- `.gitattributes` forces LF line endings.
