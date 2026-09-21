# transplantUML

**transplantUML** (`tpuml`) converts state machine documents. It parses an input file (SCXML today, more formats
planned) into a format-neutral model, then renders that model either through a Go template — so you can produce
PlantUML, source code, documentation, or any other text — or with a built-in emitter for structured formats such as JSON.

## Features

- **Lossless model**: nested states, parallel regions, history (shallow/deep) and final states, entry/exit actions,
  transition guards/actions, internal and multi-target transitions are all preserved.
- **Any text output** through the Go templating engine plus the [sprig](https://masterminds.github.io/sprig/) function
  library, with a built-in PlantUML template.
- **Structured output** with built-in emitters (JSON), and JSON is also accepted as input so it doubles as an
  interchange format.
- **Validation**: dangling targets, unknown parents, duplicate ids and similar mistakes are reported before anything
  is rendered.

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

The Makefile also has `run` (`make run ARGS="-i example/coffee-machine.scxml"`), `test`, `fmt`, `vet` and `clean`.

## Usage

```
tpuml -i input [-f format] [-t template.tmpl | -F format] [-o output]
```

| Flag | Meaning |
|------|---------|
| `-i`, `--input` | Input file (required). |
| `-f`, `--input-format` | Input format: `scxml`, `json`. Default: inferred from the extension (`.scxml`/`.xml`, `.json`). |
| `-t`, `--template` | Go template file to render with. Default: the built-in PlantUML template (`assets/puml.tmpl`). |
| `-F`, `--output-format` | Emit a built-in structured format instead of running a template: `json`. Mutually exclusive with `-t`. |
| `-o`, `--output` | Output file. Default: stdout. |
| `-h`, `--help` | Show usage. |

Each flag has a long form; `-F`/`--output-format` is distinct from `-f`/`--input-format` (flags are case-sensitive).

### Examples

```bash
# SCXML -> PlantUML with the built-in template
tpuml -i example/coffee-machine.scxml -o coffee.puml

# SCXML -> your own template
tpuml -i example/coffee-machine.scxml -t my-template.tmpl -o coffee.md

# SCXML -> JSON model, then JSON -> PlantUML (round trip)
tpuml -i example/coffee-machine.scxml -F json -o coffee.json
tpuml -i coffee.json -o coffee.puml
```

`coffee.puml` can be visualized with [PlantUML Online](https://plantuml.online).

## Model

The model below is what parsers produce and what templates receive as `.`. Hierarchy is expressed with `Parent`
(the top level is `""`), and pseudo-states are ordinary states distinguished by `Kind`.

```go
type StateKind string // "normal", "parallel", "final", "history-shallow", "history-deep"

type StateMachine struct {
    Name        string
    Initial     string        // top-level initial state
    States      []*State
    Transitions []*Transition
}

type State struct {
    Name    string
    Parent  string    // "" = top level
    Kind    StateKind
    Initial string    // for compound states
    OnEntry []string
    OnExit  []string
}

type Transition struct {
    Source   string
    Targets  []string // empty = targetless
    Event    string
    Cond     string
    Actions  []string
    Internal bool
}
```

`State` has the predicates `IsNormal`, `IsParallel`, `IsFinal`, `IsHistory`, `IsDeepHistory` and `IsPseudo`
(history or final).

The JSON emitted by `-F json` / `-e` is this structure with camelCase keys (`onEntry`, `targets`, …); empty
optional fields are omitted and `kind` defaults to `normal` when reading.

## Writing templates

Templates use Go's [`text/template`](https://pkg.go.dev/text/template) syntax. All
[sprig](https://masterminds.github.io/sprig/) functions are available, plus:

| Function | Returns |
|----------|---------|
| `State name` | the `*State`, or nil |
| `Children parent` | direct children of a state (`""` for the top level) |
| `RootStates` | same as `Children ""` |
| `HistoryOf parent` | the history pseudo-states declared under a state |
| `InitialOf name` | initial child of a state, or of the machine for `""` |
| `OutgoingTransitions source` | transitions leaving a state |
| `IncomingTransitions target` | transitions entering a state |
| `include "name" data` | like `template`, but returns a string so it can be piped (`\| indent 4`) |
| `prefix p s` | `p + s`, or `""` if `s` is empty |
| `surround p s q` | `p + s + q`, or `""` if `s` is empty |
| `joinNonEmpty sep s...` | joins the strings, skipping empty ones |

[`assets/puml.tmpl`](assets/puml.tmpl) is the reference template: it shows how to recurse through compound states,
draw parallel regions, and render pseudo-states.

## Supported SCXML

`<state>`, `<parallel>`, `<final>`, `<history type="shallow|deep">`, `initial` attribute and `<initial>` element,
`<transition>` (`event`, `cond`, multiple `target`s, `type="internal"`), `<onentry>`, `<onexit>`. Executable content
is kept as text: `<script>` verbatim, `<log>`/`<assign>`/`<raise>`/`<send>`/`<cancel>` in a short readable form,
anything else as its XML. `<datamodel>`, `<invoke>` and `<donedata>` are ignored.
