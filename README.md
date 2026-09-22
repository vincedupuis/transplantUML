# transplantUML

**transplantUML** (`tpuml`) converts state machine documents. It parses an input file into a format-neutral model
of a UML state machine, then writes that model back out: with a built-in emitter, in any of the supported document
formats (SCXML, JSON — every format `tpuml` reads it can also write), or through a Go template — so you can produce
PlantUML, source code, documentation, or any other text. Template rendering is the only one-way direction.

The idea is to keep **one** source document and generate every other representation from it. Each output writes
what it can and **warns** about what it cannot: an SCXML engine has no deferred events, PlantUML has no terminate
symbol, and so on. Warnings go to stderr and never fail the conversion.

## Features

- **UML model**: simple, compound, orthogonal (parallel) and submachine states; initial, final, terminate, shallow and
  deep history, choice, junction, fork, join, entry point and exit point pseudo-states; entry/do/exit behaviours,
  deferred events, state invariants, and variables; transitions with an event or time trigger, guard, effect and
  kind (external, local, internal), multi-target and targetless; notes and stereotypes.
- **Any text output** through the Go templating engine plus the [sprig](https://masterminds.github.io/sprig/) function
  library, with a built-in PlantUML template.
- **Every format both ways**: `scxml` and `json` are each accepted as input (`-f`) and produced as output (`-F`),
  so `tpuml` converts between them in either direction.
- **Nothing dropped silently**: the parser warns about input it has no place for, and each output warns about
  model features it can only approximate.
- **Validation**: dangling targets, unknown parents, duplicate ids, and similar mistakes are reported before anything
  is rendered.

## UML coverage

| UML concept                            | Model                                        | SCXML                                                   | PlantUML (built-in template)                        |
|----------------------------------------|----------------------------------------------|---------------------------------------------------------|-----------------------------------------------------|
| Simple, compound state                 | `normal`                                     | `<state>`                                               | `state X`, `state X { }`                            |
| Orthogonal state, regions              | `parallel`, children are the regions         | `<parallel>`                                            | regions separated by `--`                           |
| Submachine state                       | `Submachine`                                 | `<invoke src>`                                          | `state "X: ref" as X`                               |
| Initial                                | `Initial` on the machine / the state         | `initial` attr, `<initial>`                             | `[*] -->`                                           |
| Final                                  | `final`                                      | `<final>`                                               | `state X <<end>>`                                   |
| Terminate                              | `terminate`                                  | `<final>` ⚠                                             | `state X <<end>>` ⚠                                 |
| Shallow / deep history                 | `history-shallow`, `history-deep`            | `<history>`                                             | `<<history>>`, `<<history*>>`                       |
| Choice                                 | `choice`                                     | transient state with guarded eventless transitions      | `<<choice>>`                                        |
| Junction                               | `junction`                                   | transient state, `tpuml:kind="junction"`                | filled circle (`<<start>>`)                         |
| Fork                                   | `fork`                                       | transient state with one multi-target transition        | `<<fork>>`                                          |
| Join                                   | `join`                                       | transient state, `tpuml:kind="join"` ⚠                  | `<<join>>`                                          |
| Entry / exit point                     | `entry-point`, `exit-point`                  | transient state inside the compound, `tpuml:kind`       | `<<entryPoint>>`, `<<exitPoint>>`                   |
| Entry / exit behaviour                 | `OnEntry`, `OnExit`                          | `<onentry>`, `<onexit>`                                 | `X : entry / …`, `X : exit / …`                     |
| Do activity                            | `Do`                                         | `<invoke>`                                              | `X : do / …`                                        |
| Deferred events                        | `Defer`                                      | `tpuml:defer` ⚠                                         | `X : ev / defer`                                    |
| State invariant                        | `Invariant`                                  | `tpuml:invariant`                                       | `X : [ cond ]`                                      |
| Variables                              | `Variables` on the machine / the state       | `<datamodel>`                                           | `legend` / `X : name = value`                       |
| Trigger, guard, effect                 | `Event`, `Cond`, `Actions`                   | `event`, `cond`, executable content                     | `A --> B : ev [ g ] / act`                          |
| Time trigger `after(5s)`               | `After`                                      | `<send delay>` in `<onentry>`, `<cancel>` in `<onexit>` | `after(5s)`                                         |
| Completion transition                  | no `Event`, no `After`                       | eventless `<transition>`                                | unlabelled arrow                                    |
| External / local / internal transition | `Kind`                                       | `type="internal"`; `tpuml:kind="local"` ⚠               | `X : ev / act` for internal; local drawn external ⚠ |
| Note                                   | `Note` on the machine, a state, a transition | `<tpuml:note>`                                          | `note`, `note on link`                              |
| Stereotype                             | `Stereotype`                                 | `tpuml:stereotype`                                      | `state "«s»\nX" as X <<s>>`                         |

⚠ = written as an approximation, with a warning. Not modelled: signal vs. call events, change events (`when(…)` —
use a guard on a completion transition), protocol state machines.

## Requirements

- Go 1.26 or later (see `go.mod`).

## Installation

```bash
go install github.com/vincedupuis/transplantUML/cmd/tpuml@latest
```

or from a clone:

```bash
git clone https://github.com/vincedupuis/transplantUML.git
cd transplantUML
make build          # produces ./bin/tpuml, or: go build ./cmd/tpuml
```

The Makefile also has `run` (`make run ARGS="-i example/coffee-machine.scxml"`), `test`, `fmt`, `vet`, `clean` and
`generate`, which downloads the ANTLR tool and regenerates the parser of the `fsm` language
([`internal/fsm/fsm.g4`](internal/fsm/fsm.g4)) from its grammar; the generated code is committed, so a plain build
needs only Go.

### Tests

`make test` (or `go test ./...`) runs the unit, golden, and round-trip tests. Two tests additionally check the
output against something other than the project itself and skip when their tool is missing:

- **W3C schema** — every SCXML fixture and every document the SCXML emitter writes is validated with `xmllint`
  against the vendored SCXML 1.0 schema in `internal/scxml/testdata/schema/`. `xmllint` ships with macOS and most
  Linux distributions.
- **PlantUML syntax** — the rendered PlantUML is parsed by PlantUML itself (`-syntax`, no Graphviz needed). It uses
  `plantuml` from `PATH` or the jar named by `PLANTUML_JAR`; `make plantuml` downloads the jar into `bin/` and
  `make test` picks it up from there. Needs `java`.

## Usage

```
tpuml -i input [-f format] [-t template.gotmpl | -F format] [-o output]
```

| Flag                    | Meaning                                                                                               |
|-------------------------|-------------------------------------------------------------------------------------------------------|
| `-i`, `--input`         | Input file (required).                                                                                |
| `-f`, `--input-format`  | Input format: `scxml`, `json`. Default: inferred from the extension (`.scxml`/`.xml`, `.json`).       |
| `-t`, `--template`      | Go template file to render with. Default: the built-in PlantUML template (`assets/puml.gotmpl`).      |
| `-F`, `--output-format` | Write a document format instead of running a template: `scxml`, `json`. Mutually exclusive with `-t`. |
| `-o`, `--output`        | Output file. Default: stdout.                                                                         |
| `-h`, `--help`          | Show usage.                                                                                           |

Each flag has a long form; `-F`/`--output-format` is distinct from `-f`/`--input-format` (flags are case-sensitive).

Warnings are printed to stderr as `tpuml: warning: …`, one per line, and do not change the exit status:

```
$ tpuml -i machine.scxml -F scxml > out.scxml
tpuml: warning: state "work": SCXML has no deferred events; written as tpuml:defer, which engines ignore
```

### Examples

```bash
# SCXML -> PlantUML with the built-in template
tpuml -i example/coffee-machine.scxml -o coffee.puml

# SCXML -> your own template
tpuml -i example/coffee-machine.scxml -t my-template.gotmpl -o coffee.md

# SCXML -> JSON model, then JSON -> PlantUML (round trip)
tpuml -i example/coffee-machine.scxml -F json -o coffee.json
tpuml -i coffee.json -o coffee.puml

# JSON -> SCXML, and SCXML -> normalized SCXML
tpuml -i coffee.json -F scxml -o coffee.scxml
tpuml -i example/coffee-machine.scxml -F scxml
```

`coffee.puml` can be visualized with [PlantUML Online](https://plantuml.online).

### Example documents

[`example/`](example/) holds one document per group of UML concepts, each with the PlantUML `tpuml` renders from
it next to it (`<name>.puml`, kept up to date by the tests):

| Document                                                 | Demonstrates                                                                                                |
|----------------------------------------------------------|-------------------------------------------------------------------------------------------------------------|
| [`coffee-machine.scxml`](example/coffee-machine.scxml)   | the basics: flat states and event-triggered transitions                                                     |
| [`traffic-light.scxml`](example/traffic-light.scxml)     | time triggers `after(30s)`, variables, entry/exit actions, guarded and internal transitions, notes          |
| [`order.scxml`](example/order.scxml)                     | choice (recognised without markup), junction, submachine state, do activity, stereotype, terminate          |
| [`media-player.scxml`](example/media-player.scxml)       | compound state, deep history, entry/exit points, deferred events, invariant, local and internal transitions |
| [`washing-machine.scxml`](example/washing-machine.scxml) | orthogonal regions, fork and join — and the warnings PlantUML's region limitation produces                  |
| [`thermostat.json`](example/thermostat.json)             | the same concepts written directly in the model's JSON shape                                                |

Run any of them with `tpuml -i example/<name>` and compare with the `.puml` beside it; add `-F scxml` or `-F json`
to see the other formats.

## Model

The model below is what parsers produce and what templates receive as `.`. Hierarchy is expressed with `Parent`
(the top level is `""`), and pseudo-states are ordinary states distinguished by `Kind`.

```go
package main

// StateKind: "normal", "parallel", "final", "terminate", "history-shallow", "history-deep",
// "choice", "junction", "fork", "join", "entry-point", "exit-point"
type StateKind string

// TransitionKind: "external" (or ""), "local", "internal"
type TransitionKind string

type StateMachine struct {
  Name        string
  Initial     string     // top-level initial state
  Variables   []Variable // context attributes
  Note        string
  States      []*State
  Transitions []*Transition
}

type Variable struct {
  Name  string
  Value string // initial value expression
}

type State struct {
  Name       string
  Parent     string // "" = top level
  Kind       StateKind
  Initial    string   // for compound states
  OnEntry    []string // entry behaviour
  OnExit     []string // exit behaviour
  Do         []string // do activity
  Defer      []string // deferred events
  Submachine string   // referenced machine, for submachine states
  Invariant  string
  Variables  []Variable
  Stereotype string
  Note       string
}

type Transition struct {
  Source  string
  Targets []string       // empty = targetless
  Event   string         // trigger; none and no After = completion transition
  After   string         // time trigger, e.g. "5s"
  Cond    string         // guard
  Actions []string       // effect
  Kind    TransitionKind // "" = external
  Note    string
}
```

`State` has the predicates `IsNormal`, `IsParallel`, `IsFinal`, `IsTerminate`, `IsHistory`, `IsDeepHistory`,
`IsConnector` (choice, junction, fork, join, entry/exit point), `IsPseudo` (anything but normal and parallel) and
`IsBoundary` (history and entry/exit points, which must be nested). `Transition` has `IsExternal`, `IsLocal`,
`IsInternal` and `Trigger()`, which returns the event or `after(delay)`.

The JSON emitted by `-F json` is this structure with camelCase keys (`onEntry`, `targets`, …); empty
optional fields are omitted and `kind` defaults to `normal` when reading. JSON holds the whole model, so it never
warns in either direction.

## Writing templates

Templates use Go's [`text/template`](https://pkg.go.dev/text/template) syntax. All
[sprig](https://masterminds.github.io/sprig/) functions are available, plus:

| Function                     | Returns                                                                  |
|------------------------------|--------------------------------------------------------------------------|
| `States`, `Transitions`      | every state / transition, in model order                                 |
| `State name`                 | the `*State`, or nil                                                     |
| `Children parent`            | direct children of a state (`""` for the top level)                      |
| `RootStates`                 | same as `Children ""`                                                    |
| `HistoryOf parent`           | the history pseudo-states declared under a state                         |
| `InitialOf name`             | initial child of a state, or of the machine for `""`                     |
| `Ancestors name`             | parent, grandparent, … of a state, nearest first                         |
| `CommonAncestor name...`     | innermost state containing all the named states (`""` = top level)       |
| `ScopeOf transition`         | innermost state containing a transition's source and targets             |
| `OutgoingTransitions source` | transitions leaving a state                                              |
| `IncomingTransitions target` | transitions entering a state                                             |
| `include "name" data`        | like `template`, but returns a string so it can be piped (`\| indent 4`) |
| `prefix p s`                 | `p + s`, or `""` if `s` is empty                                         |
| `surround p s q`             | `p + s + q`, or `""` if `s` is empty                                     |
| `joinNonEmpty sep s...`      | joins the strings, skipping empty ones                                   |
| `warn format args...`        | records a warning for the user (printf-style) and returns `""`           |

Call `warn` wherever the template leaves something out, so the user learns what the output does not show.

[`assets/puml.gotmpl`](assets/puml.gotmpl) is the reference template: it shows how to recurse through compound states,
draw parallel regions, and render pseudo-states.

### The built-in PlantUML template

It draws everything in the coverage table above. Things to know:

- Every state is declared under its own name, so notes and behaviours can attach to any of them: final states are
  `<<end>>`, history states `<<history>>` / `<<history*>>`, and `[*]` only marks the initial state. Names PlantUML
  would misread (a dot nests, a dash breaks an arrow) get an alias: `state "a.b" as a_b`.
- A user stereotype is kept as `<<stereotype>>`, which PlantUML uses for skinparams but does not print, and shown
  as `«stereotype»` above the name. Pseudo-states already carry their kind as stereotype, so theirs warns.
- Each scope (top level, compound state, region) declares its states before its transitions, and a transition is
  drawn in the innermost scope containing both its ends. PlantUML needs this: an arrow to a nested state that has
  not been declared yet would create a copy at the wrong level.
- PlantUML refuses arrows across the boundary of any parallel region but the first; such transitions are skipped
  with a warning. Regions are anonymous in PlantUML, so a compound region's own transitions, behaviours and note
  cannot be drawn either.
- A junction is drawn as a filled circle using `<<start>>`, since PlantUML has no junction stereotype. Terminate has
  no symbol either and is drawn as a final state. PlantUML ignores description lines on a final state, so its entry
  and exit behaviours are written into its note.
- Local transitions are drawn as external ones, and a note on an internal transition is not drawn. Both warn.
- The machine name becomes the diagram `title`; real line breaks in actions and behaviours become `\n`.

## SCXML

### Native elements

`<state>`, `<parallel>`, `<final>`, `<history type="shallow|deep">`, `initial` attribute and `<initial>` element,
`<transition>` (`event`, `cond`, multiple `target`s, `type="internal"`), `<onentry>`, `<onexit>`, `<datamodel>` with
`<data>` (on the machine and on states, as variables) and `<invoke>`. Executable content is kept as text: `<script>`
verbatim, `<log>`/`<assign>`/`<raise>`/`<send>`/`<cancel>` in a short readable form, anything else as its XML.
Elements with no place in the model (`<donedata>`, a top-level `<script>`, …) are dropped with a warning.

`<invoke>` is a **submachine state** when it invokes another SCXML document (`type` absent or `scxml`, with a `src`)
and a **do activity** otherwise: `invoke(src)` or `invoke(src, type)` for a plain element, its XML when it has
`<param>`, `<finalize>` or `<content>` children.

### Idioms recognised as UML

- A `<state>` with no content of its own (no children, actions, invokes, or data) and only eventless transitions is
  a **choice** when it has several transitions and at least one guard, and a **fork** when it has one transition
  with several targets. Add `tpuml:kind="normal"` to keep such a state as a state.
- A `<send event="E" delay="D">` in `<onentry>` whose event `E` is consumed by a `<transition>` of the same state
  is a **time trigger**: the transition gets `After: D` and the `<send>` (and a matching `<cancel>` in `<onexit>`)
  are not treated as actions.

### The tpuml extension vocabulary

SCXML permits attributes and elements from other namespaces everywhere, and engines ignore them. `tpuml` uses the
namespace `https://github.com/vincedupuis/transplantUML` (any prefix; `tpuml` below) for what SCXML cannot say:

| Extension                                                                    | On                                   | Meaning                                   |
|------------------------------------------------------------------------------|--------------------------------------|-------------------------------------------|
| `tpuml:kind="junction\|join\|entry-point\|exit-point\|choice\|fork\|normal"` | `<state>`                            | the pseudo-state kind (a transient state) |
| `tpuml:kind="terminate"`                                                     | `<final>`                            | a terminate pseudo-state                  |
| `tpuml:kind="local"`                                                         | `<transition>`                       | a local transition                        |
| `tpuml:defer="e1 e2"`                                                        | `<state>`                            | deferred events                           |
| `tpuml:invariant="expr"`                                                     | `<state>`                            | state invariant                           |
| `tpuml:stereotype="name"`                                                    | `<state>`                            | stereotype                                |
| `<tpuml:note>text</tpuml:note>`                                              | `<scxml>`, `<state>`, `<transition>` | a note                                    |

See [`internal/scxml/testdata/uml.scxml`](internal/scxml/testdata/uml.scxml) for a document using all of them.

### Writing SCXML (`-F scxml`)

Everything the model holds is written, natively where SCXML has the element and through the extension vocabulary
otherwise; the namespace is declared only when it is used. The document validates against the W3C schema. It is
equivalent, not byte-identical, to the one it came from:

- the initial child is always written as an `initial` attribute, never as an `<initial>` element;
- executable content becomes `<script>` bodies holding the text the parser produced, except for elements it kept as
  XML, which are written back as themselves;
- a time trigger becomes `<send event="after.D" delay="D" id="…">` in `<onentry>`, the matching `<cancel>` in
  `<onexit>`, and a transition on that event;
- comments and anything the parser dropped are not preserved.

Warnings name what SCXML can only approximate: join (the first region to reach it leaves the parallel state),
terminate (a `<final>`, which runs exit actions), local transitions (written external), deferred events (ignored by
engines), and free-text do activities (written as `<invoke>` content).
