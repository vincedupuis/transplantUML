# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

`tpuml` is a Go CLI that converts state machine documents. Any supported input format is parsed into one
format-neutral model (`internal/model`) of a UML state machine, which is then written back out either by a built-in
emitter (document formats: SCXML, JSON) or through a Go `text/template` (free-form text: PlantUML, code, docs).
Template rendering is output-only. The bundled `assets/puml.gotmpl` (embedded in the binary, used when `-t` is
omitted) produces PlantUML.

The user keeps one source document and generates outputs from it; round-tripping is *not* a goal. The goal is to
cover as much of UML as possible on the input side and, on the output side, to write what the format can express
and **warn** (never silently drop) about the rest. Parsers, emitters and `render.Render` all return
`model.Warnings`; `cmd/tpuml` prints them to stderr as `tpuml: warning: …`.

## Commands

```bash
make build                          # produces ./bin/tpuml (see the Makefile for run/test/fmt/vet/clean)
go build ./...
go build ./cmd/tpuml                # produces ./tpuml
make vet                            # go vet -unreachable=false ./... (the generated ANTLR parser trips that check)
go test ./...
go test ./internal/scxml -run TestEdgeCases          # one test
go test ./internal/render -run TestPlantUMLGolden/edge
make plantuml                       # download the PlantUML jar into bin/ so TestPlantUMLSyntax runs (make test picks it up)
make generate                       # download the ANTLR jar into bin/ and regenerate internal/fsm/parser from fsm.g4

# End-to-end
./tpuml -i example/coffee-machine.scxml                 # PlantUML to stdout
./tpuml -i example/coffee-machine.scxml -F json         # model as JSON
```

Golden files for the PlantUML template live in `internal/render/testdata/*.puml`; regenerate one with
`go run ./cmd/tpuml -i <input> -o internal/render/testdata/<name>.puml` after checking the diff is intended.

## Architecture

Pipeline: `cmd/tpuml/main.go` (`run() error`) → `format.ParserFor(name).Parse` → `*model.StateMachine` +
warnings → `sm.Validate()` → `render.Render(sm, tmpl)` **or** `format.EmitterFor(name).Emit(sm)` → output +
warnings.

- **`internal/model`** — the contract everything else depends on, shaped after UML rather than any one format.
  Flat `States` with `Parent` links (top level is `""`), and flat `Transitions`. Pseudo-states are real `State`s
  distinguished by `Kind` (`normal`, `parallel`, `final`, `terminate`, `history-shallow`, `history-deep`, and the
  connectors `choice`, `junction`, `fork`, `join`, `entry-point`, `exit-point`); a history state's default
  transition is an ordinary `Transition` whose `Source` is the history state. `Initial` is a property
  (`StateMachine.Initial`, `State.Initial`), not a synthetic transition. States also carry `Do`, `Defer`,
  `Submachine`, `Invariant`, `Variables`, `Stereotype`, `Note`; transitions carry `After` (time trigger), `Kind`
  (`""`/external, local, internal) and `Note`. `Validate()` is the single place structural rules live; extend it
  when the model grows. `Ancestors`/`CommonAncestor`/`ScopeOf` exist for templates that must place a transition
  in a scope. The README's coverage table lists what each format does with each concept; keep it in step.
- **`internal/format`** — `Parser` / `Emitter` interfaces plus the name→implementation tables and extension
  detection. To add a format: implement `Parser` and/or `Emitter` in its own package under `internal/`, register
  each in its table, add its extension (`TestExtensionsHaveParsers` checks every extension names a parser). Package
  names avoid stdlib clashes
  (`jsonsm`, not `json`).
- **`internal/scxml`** — `Parser` is a recursive `etree` walk over `<state>/<parallel>/<final>/<history>` children
  (direct children only, so `<initial>`'s inner `<transition>` is not mistaken for a real transition). Executable
  content is flattened to strings; unknown elements are kept as raw XML, normalized by `rawXML` (unindented,
  canonical escaping) so the string is stable. What SCXML has no element for comes from the `tpuml` extension
  namespace (`ExtNamespace`; matched by URI, not prefix): `tpuml:kind`, `tpuml:defer`, `tpuml:invariant`,
  `tpuml:stereotype`, `<tpuml:note>`. Two idioms are recognised without markup: transient states as choice/fork
  (`connectorKind`) and `<send delay>`+`<cancel>` as a time trigger (`timers`). Anything else unknown under a
  state or the root raises a parser warning. `Emitter` (`emit.go`) rebuilds the tree from `Parent` links, writes
  the initial child as an attribute, turns action strings back into `<script>` bodies — except those that are
  XML, which are re-inserted as elements — declares the extension namespace only when used, and warns for join,
  terminate, local, defer and free-text do activities. States it cannot reach from the top level are an error.
- **`internal/fsm`** — tpuml's own DSL (`fsm name { state s { on ev [guard] / actions goto target } }`), an
  ANTLR4 grammar in `fsm.g4`. `parser/` is generated from it (`make generate`, Go target with `-visitor
  -no-listener`) and committed so the build needs no Java; never edit it by hand, and regenerate it after any
  grammar change. `goto` targets are `.` (self), `final`, `H` (history) or a path (`/a/b` absolute, `../b`
  relative). The `Parser` that walks the parse tree into the model is not written yet, so the language is not
  registered in `internal/format` and only reaches as far as the generated parser.
- **`internal/jsonsm`** — the model's own JSON shape (struct tags in `model`). Parser uses
  `DisallowUnknownFields`; round-trip equality with the SCXML parser is tested.
- **`internal/render`** — registers sprig plus project helpers (`include`, `prefix`, `surround`, `joinNonEmpty`,
  `warn`) and the model accessors as template functions. `warn` records a warning and returns `""`; `Render`
  returns the collected warnings. `joinNonEmpty` exists because sprig's `join` has the signature `join sep list`;
  don't shadow sprig names.
- **`assets/puml.gotmpl`** — every emitted line starts with `\n` so nested blocks compose via
  `include ... | trimPrefix "\n" | indent 4`. Each scope (top level, compound body, region) declares its states
  first, then `scopeTransitions` draws the transitions whose `ScopeOf` is that scope — PlantUML creates a state
  where it first sees the name, so a forward reference into a nested state would create a stray copy, and arrows
  inside a parallel region are only accepted inside that region's braces. Every model state is declared by name
  (`id` aliases names with characters PlantUML misreads, e.g. `state "a.b" as a_b`): finals/terminate as
  `<<end>>` (their entry/exit go into the note, PlantUML ignores their description lines), history as
  `<<history>>`/`<<history*>>`, connectors as `<<choice>>`/`<<fork>>`/`<<join>>`/`<<entryPoint>>`/`<<exitPoint>>`
  (junction as `<<start>>`, the filled circle); only the initial state is drawn as `[*]`. A user stereotype is
  kept as `<<s>>` and shown as `«s»` in the label since PlantUML does not print it. Parallel regions are the
  *bodies* of the region states separated by `--` (a leaf or orthogonal region is drawn as the state itself; a
  compound region's own transitions, behaviours and note warn). PlantUML refuses arrows across the boundary of
  any region but the first (`confined`), so those warn instead. Verified with the PlantUML jar: `-syntax` accepts
  unknown stereotypes and silently ignores unsupported constructs, so test new notations by rendering
  (`-tpng`, with `-Playout=smetana` if Graphviz is missing) and looking at the image.

## Conventions

- README.md documents the model, template functions and CLI flags; update it when any of those change.
- Test fixtures: `example/coffee-machine.scxml` (simple, flat), `internal/scxml/testdata/edge.scxml`
  (parallel, `<initial>` element, deep history, final, onentry, multi-target) and `internal/scxml/testdata/uml.scxml`
  (every UML concept: connectors, terminate, submachine, do, defer, invariant, variables, time trigger, local and
  internal transitions, notes, stereotype; also exercises every warning). Add new concepts to `uml.scxml` and its
  goldens (`internal/render/testdata/uml.puml`, `internal/scxml/testdata/uml.emitted.scxml`) and to the warning
  lists in `TestPlantUMLWarnings`, `TestEmitWarnings` and `TestUMLConcepts`.
- `example/` has one document per group of concepts (see the README table) with its rendered `.puml` beside it.
  `goldens()` in `internal/render/render_test.go` globs the directory, so every `example/*.scxml|json` must have
  a matching `.puml` (regenerate with `go run ./cmd/tpuml -i example/<name> -o example/<name>.puml`), is checked
  by PlantUML's syntax test, and every `.scxml` there is validated against the W3C schema.
- `.gitattributes` forces LF line endings.
