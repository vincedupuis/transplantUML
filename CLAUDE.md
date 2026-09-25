# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

`fsm` is a Go CLI that converts state machine documents. Any supported input format is parsed into one
format-neutral model (`internal/model`) of a UML state machine, which is then written back out either by a built-in
emitter (document formats: SCXML, JSON) or through a Go `text/template` (free-form text: PlantUML, code, docs).
Template rendering is output-only. The bundled templates are embedded in the binary and `-t` takes their name when
no file has it: `assets/puml.gotmpl` (`puml`, used when `-t` is omitted) produces PlantUML, `assets/sml.gotmpl`
(`sml`) a C++20 state machine on Boost.SML as four files. A template that writes several files starts each with
the `file` function; `render.Files` splits the output at those marks and `-o` then names a folder.

The user keeps one source document and generates outputs from it; round-tripping is *not* a goal. The goal is to
cover as much of UML as possible on the input side and, on the output side, to write what the format can express
and **warn** (never silently drop) about the rest. Parsers, emitters and `render.Render` all return
`model.Warnings`; `cmd/fsm` prints them to stderr as `fsm: warning: …`.

## Commands

```bash
make build                          # produces ./bin/fsm (see the Makefile for run/test/fmt/vet/clean)
go build ./...
go build ./cmd/fsm                  # produces ./fsm
make vet                            # go vet -unreachable=false ./... (the generated ANTLR parser trips that check)
go test ./...
go test ./internal/scxml -run TestEdgeCases          # one test
go test ./internal/render -run TestPlantUMLGolden/edge
make plantuml                       # download the PlantUML jar into bin/ so TestPlantUMLSyntax runs (make test picks it up)
make sml                            # download the Boost.SML header into bin/ so TestSMLCompiles runs (needs a C++20 compiler)
make generate                       # download the ANTLR jar into bin/ and regenerate internal/fsm/parser from fsm.g4

# End-to-end
./fsm -i example/coffee-machine.scxml                   # PlantUML to stdout
./fsm -i example/coffee-machine.scxml -F json           # model as JSON
```

Golden files for the PlantUML template live in `internal/render/testdata/*.puml`; regenerate one with
`go run ./cmd/fsm -i <input> -o internal/render/testdata/<name>.puml` after checking the diff is intended. Those for
the SML template are folders, `internal/render/testdata/sml/<name>/` (listed in `smlGoldens`), regenerated with
`-t sml -o internal/render/testdata/sml/<name>`.

## Architecture

Pipeline: `cmd/fsm/main.go` (`run() error`) → `format.ParserFor(name).Parse` → `*model.StateMachine` +
warnings → `sm.Validate()` → `render.Render(sm, tmpl)` **or** `format.EmitterFor(name).Emit(sm)` → output +
warnings.

- **`internal/model`** — the contract everything else depends on, shaped after UML rather than any one format.
  Flat `States` with `Parent` links (top level is `""`), and flat `Transitions`. Pseudo-states are real `State`s
  distinguished by `Kind` (`normal`, `parallel`, `final`, `terminate`, `history-shallow`, `history-deep`, and the
  connectors `choice`, `junction`, `fork`, `join`, `entry-point`, `exit-point`); a history state's default
  transition is an ordinary `Transition` whose `Source` is the history state. `Initial` is a property
  (`StateMachine.Initial`, `State.Initial`, with the transition's effect in `InitialActions`), not a synthetic
  transition. An entry or exit point whose parent is a submachine state is UML's
  connection point reference to the referenced machine's point of that name (`IsReference`). States also carry `Do`, `Defer`,
  `Submachine`, `Invariant`, `Variables`, `Stereotype`, `Note`, and a note on each thing they list (`EntryNote`,
  `ExitNote`, `DoNote`, `InvariantNote`, `DeferNotes` by event); transitions carry `After` (time trigger), `Kind`
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
  `tpuml:stereotype`, `<tpuml:note>` (on a state, `about="entry|exit|do|invariant|defer"` says what it
  describes). Three idioms are recognised without markup: transient states as choice/fork (`connectorKind`),
  `<send delay>`+`<cancel>` as a time trigger (`timers`), and `done.state`/`done.invoke` as the completion event
  (`completes`); the emitter writes completion transitions back on those events. Anything else unknown under a state or the root raises a
  parser warning. `Emitter` (`emit.go`) rebuilds the tree from `Parent` links, writes
  the initial child as an attribute, turns action strings back into `<script>` bodies — except those that are
  XML, which are re-inserted as elements — writes a fork's transitions as one multi-target transition and an
  `else` branch last with no `cond` (engines take the first enabled transition), declares the extension
  namespace only when used, and warns for join, terminate, local, defer and free-text do activities, and leaves out a
  submachine state's entry and exit points with a warning. States it cannot reach from the top level are an error.
- **`internal/fsm`** — the command's own DSL (`fsm name { state s { on ev [guard] / actions goto target } }`), an
  ANTLR4 grammar in `fsm.g4`. `parser/` is generated from it (`make generate`, Go target with `-visitor
  -no-listener`) and committed so the build needs no Java; never edit it by hand, and regenerate it after any
  grammar change. Kinds use UML's names: `state`, `parallel state` holding `region`s, `submachine` (named after
  the machine it refers to, so its body holds clauses and the entry and exit points it references: a
  connection point reference, whose entry point has no `goto`), and the pseudostates `choice`/`junction` (branches,
  `[else]` becomes `Cond` "else"), `fork`, `join`, `entry point`/`exit point`, declared without `state`. A clause
  with no trigger is a completion transition. `goto` targets are a state name, `final`, `terminate`, or history as
  `H`/`H*`, alone or after a state name (`s.H*`); a line `H <<s>> / actions goto target` inside a state or region
  annotates that history and optionally gives it its default transition, and `final state [name]` / `terminate
  state [name]` declare a named final or terminate or, without a name, annotate the one `goto final` / `goto
  terminate` reaches in that scope; `goto local <target>`
  makes a local transition, whose target must be inside the source, and a clause without `goto` is internal; one name
  reaches any state because the names are one namespace for the whole machine, as they are in the model. A state
  may be marked `initial` (`initial state s { … }`), as may a choice or a junction since UML lets the
  initial transition lead to one, naming its parent's starting child or, at the top level, the
  machine's; a line `initial / actions goto s` says the same and is the only way to give the initial
  transition an effect (`InitialActions`). `build.go` holds `Parser` and the `builder` that walks the parse tree into the
  model: `declare` creates every state first, in document order, then `walk` resolves the transitions, because a
  `goto` may name a state declared further down. The unnamed final, terminate and history states have no name
  to declare, so `synthesize` creates them on first use, from a `goto` or an unnamed declaration, as
  `<scope>.final`, `<scope>.terminate`, `<state>.H` and `<state>.H-deep` (an Identifier holds letters, digits and `_` only, so no document can declare those names
  itself). The builder reports only what the model cannot hold or would misplace — two `initial` children in one
  scope would collapse into one `State.Initial`, a repeated state name would collapse the builder's own index,
  the machine has nowhere to put a top-level `on` clause, two unnamed finals, terminates or `H` lines in one scope
  would fold into one state, and a history or final state made directly inside a
  parallel state would become a region — and `model.Validate()` does the rest.
- **`internal/jsonsm`** — the model's own JSON shape (struct tags in `model`). Parser uses
  `DisallowUnknownFields`; round-trip equality with the SCXML parser is tested.
- **`internal/render`** — registers sprig plus project helpers (`include`, `prefix`, `surround`, `joinNonEmpty`,
  `warn`, `file`) and the model accessors as template functions. `warn` records a warning and returns `""`; `Render`
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
- **`assets/sml.gotmpl`** — Boost.SML, as `<Name>FsmEvents.h` (an interface, one pure virtual `void` method per
  event), `<Name>FsmActions.h` (an interface, a `bool … const` per guard and a `void` per action, which the user
  implements), `<Name>Fsm.h` (the state machine class, implementing the events interface, built from the actions,
  its SML machine behind a `unique_ptr<Machine>`) and `<Name>Fsm.cpp`, the only file including SML. No namespace in
  the public files; in the `.cpp` everything sits in an anonymous namespace, events in `namespace event`. A guard or
  action name becomes a lambda calling it on the actions interface, which SML injects. Everything from SML is
  qualified (`sml::event`, `sml::X`, …) with only `sml::literals` and the guard/action operators imported, because
  `using namespace sml` makes event names like `back` ambiguous. Each compound or orthogonal state is a struct
  (`HasTable`), an orthogonal state's regions are flattened into its table (`TableOf`), and a transition is lifted by
  `Lift` into the innermost table holding both ends, warning when that moves an end. Those table functions live in
  `internal/render/tables.go`; the template keeps the rendering and every `warn`. Rows are built with sentinels
  (`⟨g:…⟩`, `⟨a:…⟩`, `⟨e:…⟩`, `⟨s:…⟩`) so `body` can declare only the lambdas a table uses and fully qualify the
  event namespace or a struct a lambda would hide. The tables are rendered first and the interfaces derived from the
  lambdas they declare. Behaviour verified by compiling and running generated code against SML: guards on one source
  are tried in table order, an anonymous transition leaving a composite state waits for it to reach `X`, and an
  orthogonal state completes when every region has. `TestSMLCompiles` compiles every golden input with a stub
  implementation of the actions.

## Conventions

- README.md documents the model, template functions and CLI flags; update it when any of those change.
- A concept added to `uml.scxml` also goes into `internal/render/testdata/sml/uml/` and `TestSMLWarnings`.
- Test fixtures: `example/coffee-machine.scxml` (simple, flat), `internal/scxml/testdata/edge.scxml`
  (parallel, `<initial>` element, deep history, final, onentry, multi-target) and `internal/scxml/testdata/uml.scxml`
  (every UML concept: connectors, terminate, submachine, do, defer, invariant, variables, time trigger, local and
  internal transitions, notes, stereotype; also exercises every warning). Add new concepts to `uml.scxml` and its
  goldens (`internal/render/testdata/uml.puml`, `internal/scxml/testdata/uml.emitted.scxml`) and to the warning
  lists in `TestPlantUMLWarnings`, `TestEmitWarnings` and `TestUMLConcepts`.
- `example/` has one document per group of concepts (see the README table) with its rendered `.puml` beside it.
  `goldens()` in `internal/render/render_test.go` globs the directory, so every `example/*.scxml|json` must have
  a matching `.puml` (regenerate with `go run ./cmd/fsm -i example/<name> -o example/<name>.puml`), is checked
  by PlantUML's syntax test, and every `.scxml` there is validated against the W3C schema.
- `.gitattributes` forces LF line endings.
