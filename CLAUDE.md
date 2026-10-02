# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

`fsm` is a Go CLI that converts state machine documents. Any supported input format is parsed into one
format-neutral model (`internal/model`) of a UML state machine, which is then written back out either by a built-in
emitter (document formats: SCXML, JSON) or through a Go `text/template` (free-form text: PlantUML, code, docs).
Template rendering is output-only. The bundled templates are embedded in the binary and `-t` takes their name when
no file has it: `assets/puml.gotmpl` (`puml`, used when `-t` is omitted) produces PlantUML, `assets/sml.gotmpl`
(`sml`) a C++20 state machine on Boost.SML as five files. A template that writes several files starts each with
the `file` function; `render.Files` splits the output at those marks and `-o` then names a folder.

The user keeps one source document and generates outputs from it; round-tripping is *not* a goal. The goal is to
cover as much of UML as possible on the input side and, on the output side, to write what the format can express.
What would break the machine's behaviour in the destination is an **error**, never a warning. A **warning** only
marks a concept with no 1:1 counterpart that a workaround makes behave the same (PlantUML, a diagram, warns and
draws a note saying what is meant). Conditional cases are decided per case: SML's fork, a state's entry point, a
transition into a composite state warn when entering by default reaches the same states (`DefaultEntry`), and fail
otherwise. Parsers, emitters and `render.Render` all return `model.Warnings` and fail with every error they found
(`errors.Join`); `cmd/fsm` prints the warnings to stderr as `fsm: warning: …`, then the error. Guards, actions, do activities,
invariants and named delays hold names only (`model.IsName`, `model.IsCondition`: names with `not`/`and`/`or`/
parentheses, `model.IsDelay`: a duration or a name), which generated code calls as functions; anything else is an
error from `Validate`, never a warning.

## Commands

```bash
make build                          # produces ./bin/fsm (see the Makefile for run/test/fmt/vet/clean)
go build ./...
go build ./cmd/fsm                  # produces ./fsm
make vet                            # go vet ./...
go test ./...
go test ./internal/scxml -run TestEdgeCases          # one test
go test ./internal/render -run TestPlantUMLGolden/edge
make plantuml                       # download the PlantUML jar into bin/ so TestPlantUMLSyntax runs (make test picks it up)
make sml                            # download the Boost.SML header into bin/ so TestSMLCompiles runs (needs a C++20 compiler)

# End-to-end
./fsm -i example/coffee-machine.scxml                   # PlantUML to stdout
./fsm -i example/coffee-machine.scxml -F json           # model as JSON
```

Golden files for the PlantUML template live in `internal/render/testdata/*.puml`; regenerate one with
`go run ./cmd/fsm -i <input> -o internal/render/testdata/<name>.puml` after checking the diff is intended. Those for
the SML template are folders, `internal/render/testdata/sml/<name>/` (listed in `smlGoldens`), regenerated with
`-t sml -o internal/render/testdata/sml/<name>` from the input beside the folder (`kiosk.json`, `shop.json`,
`uml.scxml`: the examples and `uml.scxml` without what Boost.SML cannot run) or `example/support.json`.

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
  (direct children only, so `<initial>`'s inner `<transition>` is not mistaken for a real transition). An action is
  a `<script>` holding a name (`actions`); any other executable content is a parser error, and so is an
  `<invoke>` that is neither a submachine nor a do activity named by its `src`. What SCXML has no element for comes from the `tpuml` extension
  namespace (`ExtNamespace`; matched by URI, not prefix): `tpuml:kind`, `tpuml:defer`, `tpuml:invariant`,
  `tpuml:stereotype`, `<tpuml:note>` (on a state, `about="entry|exit|do|invariant|defer"` says what it
  describes). Three idioms are recognised without markup: transient states as choice/fork (`connectorKind`),
  `<send delay>`+`<cancel>` as a time trigger (`timers`), and `done.state`/`done.invoke` as the completion event
  (`completes`); the emitter writes completion transitions back on those events. `type="internal"` with a target
  is a local transition when every target lies inside the source, external otherwise (`localTransitions`). An
  unknown SCXML element under a state or the root is a parser error (`unsupported`); one of another namespace,
  which engines ignore, a warning. `Emitter` (`emit.go`) rebuilds the tree from `Parent` links, writes
  the initial child as an attribute, writes each action as a `<script>` holding its name and each do activity as
  `<invoke type="tpuml:do" src>`, writes a fork's transitions as one multi-target transition and an
  `else` branch last with no `cond` (engines take the first enabled transition), writes a local transition as
  `type="internal"`, and declares the extension namespace only when used. It warns for a top-level terminate
  (a `<final>`) and for the machine's initial effect (a transient state `tpuml:kind="initial"` the machine starts in,
  which the parser folds back into `InitialActions`), and fails (`fail`) on join, defer, a nested terminate, a completion
  that would not wait for a do activity, and a submachine state's entry and exit points. States it cannot reach
  from the top level are an error.
- **`internal/jsonsm`** — the model's own JSON shape (struct tags in `model`). Parser uses
  `DisallowUnknownFields`; round-trip equality with the SCXML parser is tested.
- **`internal/render`** — registers sprig plus project helpers (`include`, `prefix`, `surround`, `joinNonEmpty`,
  `warn`, `error`, `file`) and the model accessors as template functions. `warn` and `error` record a warning or an
  error and return `""`; `Render` returns the collected warnings and fails with every error. `joinNonEmpty` exists because sprig's `join` has the signature `join sep list`;
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
  compound region's own transitions, behaviours and note warn and go into the parallel state's note). PlantUML
  refuses arrows across the boundary of any region but the first (`confined`), so those warn and the note on their
  source lists them (`undrawn`); terminate, a pseudostate's stereotype and a note on an internal transition likewise
  go into a note. Verified with the PlantUML jar: `-syntax` accepts
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
  `Lift` into the innermost table holding both ends; moving an end warns when the behaviour stays the same (a
  completion lifted through `X`, a region's transition, `DefaultEntry`, `AlwaysActive`) and is an error otherwise. Those table functions live in
  `internal/render/tables.go`; the template keeps the rendering and every `warn`. Rows are built with sentinels
  (`⟨g:…⟩`, `⟨a:…⟩`, `⟨e:…⟩`, `⟨s:…⟩`) so `body` can declare only the lambdas a table uses and fully qualify the
  event namespace or a struct a lambda would hide. The tables are rendered first and the interfaces derived from the
  lambdas they declare. Behaviour verified by compiling and running generated code against SML: guards on one source
  are tried in table order, an anonymous transition leaving a composite state waits for it to reach `X`, and an
  orthogonal state completes when every region has. `TestSMLCompiles` compiles every SML golden input with a stub
  implementation of the actions and a stand-in for each machine a submachine state runs.
  Every machine waits in `internal::stopped` until `enterFsm(Entry)` (the initial state or one of its entry points)
  and ends in `X` or one of its exit points, which `report` tells the `<Name>FsmListener`; `stopFsm` sends `internal::stop`, which every state of the
  machine's table answers by going back to `stopped` (SML exits active inner states and regions first), then
  rebuilds the SML machine (`terminateFsm`, which runs no exit behaviour), and the names `enterFsm`/`stopFsm`/
  `terminateFsm` keep them apart from event methods. A terminate state is `internal::terminated`, whose entry sets
  the injected `status.terminated`; `report` checks it after the event, so a terminate at any depth, or in a
  submachine (whose listener sets the flag), ends the machine and tells `onTerminated`. A submachine state runs the other
  generated machine, which the constructor takes by reference: its entry and exit behaviours call `enterFsm` and
  `stopFsm` through the `submachines` struct SML injects, and a listener per submachine state turns what that machine
  reports into internal events, queued by `Machine::process` while an event is being processed.
  `TestSMLSubmachineRuns` links two machines and checks the order of their actions.
  A time trigger has a timer per source state and delay (`timerList`; a compound region's timers belong to its
  orthogonal state). The state's entry starts it and its exit cancels it through `FsmTimers.h`, one interface
  written identically for every machine, which the user implements once. When it fires, the machine processes
  `internal::<state>_after_<delay>`. The callback holds a `weak_ptr` to its slot in the `timers` dependency, so a
  late fire, or one after the machine is rebuilt, does nothing. A named delay is a method of the actions
  interface. `TestSMLTimersRun` checks the starts, cancels and stale fires.
  A state invariant calls `invariant<Name>()` (a `bool … const` on the actions interface) for each of its names. The
  state's entry and exit behaviours set and clear its flag in the `invariants` dependency, and `checkInvariants`,
  run after every processed event, tells the listener `onInvariantViolated(state)` for each flagged state whose
  invariant is false. `TestSMLInvariantsRun` checks when it is told.

## Conventions

- README.md documents the model, template functions and CLI flags; update it when any of those change.
- A concept added to `uml.scxml` also goes into the per-output copies when that output can run it:
  `internal/scxml/testdata/uml-scxml.scxml` (golden `uml-scxml.emitted.scxml`) and
  `internal/render/testdata/sml/uml.scxml` (golden folder `sml/uml/`). What an output cannot run goes into its
  error test instead (`TestEmitBehaviourErrors`, `TestSMLErrors`).
- Test fixtures: `example/coffee-machine.scxml` (simple, flat), `internal/scxml/testdata/edge.scxml`
  (parallel, `<initial>` element, deep history, final, onentry, multi-target) and `internal/scxml/testdata/uml.scxml`
  (every UML concept: connectors, terminate, submachine, do, defer, invariant, variables, time trigger, local and
  internal transitions, notes, stereotype; also exercises every PlantUML warning). `internal/scxml/testdata/expressions.scxml`
  holds the executable content that is an error (`TestExpressionErrors`). Add new concepts to `uml.scxml` and its
  golden (`internal/render/testdata/uml.puml`), to the per-output copies above, and to the lists in
  `TestPlantUMLWarnings`, `TestEmitWarnings`, `TestSMLWarnings`, `TestUMLConcepts` and the error tests.
- `example/` has one document per group of concepts (see the README table) with its rendered `.puml` beside it.
  `goldens()` in `internal/render/render_test.go` globs the directory, so every `example/*.scxml|json` must have
  a matching `.puml` (regenerate with `go run ./cmd/fsm -i example/<name> -o example/<name>.puml`), is checked
  by PlantUML's syntax test, and every `.scxml` there is validated against the W3C schema.
- `.gitattributes` forces LF line endings.
