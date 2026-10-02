# transplantUML

**transplantUML** (`fsm`) converts state machine documents. It parses an input file into a format-neutral model
of a UML state machine, then writes that model back out: with a built-in emitter, in any of the supported document
formats (SCXML, JSON), or through a Go template — so you can produce PlantUML, source code, documentation, or any
other text.

The idea is to keep **one** source document and generate every other representation from it.
Each output writes what it can.
Where a format has no counterpart for a concept but a workaround behaves the same, the output **warns**: PlantUML has no terminate symbol, so it draws a final state and a note.
Where the output would behave differently, the conversion **fails**: an SCXML engine has no deferred events, so it would handle them at once.
Warnings go to stderr and never fail the conversion.

## Features

- **UML model**: simple, compound, orthogonal (parallel) and submachine states; initial, final, terminate, shallow and
  deep history, choice, junction, fork, join, entry point and exit point pseudo-states; entry/do/exit behaviours,
  deferred events, state invariants, and variables; transitions with an event or time trigger, guard, effect and
  kind (external, local, internal), multi-target and targetless; notes and stereotypes.
- **Any text output** through the Go templating engine plus the [sprig](https://masterminds.github.io/sprig/) function
  library, with two built-in templates: PlantUML, and C++ for [Boost.SML](https://github.com/boost-ext/sml).
- **Formats**: `scxml` and `json` are each accepted as input (`-f`) and produced as output (`-F`).
- **Nothing dropped silently**: input the model has no place for is an error, and so is a model feature an output
  cannot run; an output that runs one through a workaround warns.
- **Validation**: dangling targets, unknown parents, duplicate ids, and similar mistakes are reported before anything
  is rendered.

## UML coverage

| UML concept                            | Model                                                 | SCXML                                                                | PlantUML (built-in template)                        | Boost.SML (`-t sml`)                          |
|----------------------------------------|-------------------------------------------------------|----------------------------------------------------------------------|-----------------------------------------------------|-----------------------------------------------|
| Simple, compound state                 | `normal`                                              | `<state>`                                                            | `state X`, `state X { }`                            | `"X"_s`, a struct per compound state          |
| Orthogonal state, regions              | `parallel`, children are the regions                  | `<parallel>`                                                         | regions separated by `--`                           | one table, a `*` initial state per region     |
| Submachine state                       | `Submachine`                                          | `<invoke src>`                                                       | `state "X: ref" as X`                               | runs that machine's generated class           |
| Initial                                | `Initial` on the machine / the state                  | `initial` attr, `<initial>`                                          | `[*] -->`                                           | `*` on the initial state                      |
| Effect of the initial transition       | `InitialActions` on the machine / the state           | `<initial><transition>`; the machine's in a transient state ⚠        | `[*] --> X : / act`                                 | on `enterFsm`; transient in a state           |
| Final                                  | `final`                                               | `<final>`                                                            | `state X <<end>>`                                   | `sml::X`                                      |
| Terminate                              | `terminate`                                           | `<final>` at the top level ⚠; else ✗                                 | `state X <<end>>` ⚠                                 | ends the machine, `onTerminated`              |
| Shallow / deep history                 | `history-shallow`, `history-deep`                     | `<history>`                                                          | `<<history>>`, `<<history*>>`                       | `(sml::H)` on the initial state; ⚠ or ✗       |
| Choice                                 | `choice`                                              | transient state with guarded eventless transitions                   | `<<choice>>`                                        | state with guarded anonymous transitions      |
| Junction                               | `junction`                                            | transient state, `tpuml:kind="junction"`                             | filled circle (`<<start>>`)                         | state with guarded anonymous transitions      |
| Fork                                   | `fork`                                                | transient state with one multi-target transition                     | `<<fork>>`                                          | enters the orthogonal state ⚠ or ✗            |
| Join                                   | `join`                                                | ✗                                                                    | `<<join>>`                                          | regions end in `X`, then a completion ⚠ or ✗  |
| Entry / exit point                     | `entry-point`, `exit-point`                           | transient state inside the compound, `tpuml:kind`                    | `<<entryPoint>>`, `<<exitPoint>>`                   | the machine's: `Entry`, listener; else ⚠ or ✗ |
| Connection point reference             | entry / exit point whose parent is a submachine state | ✗                                                                    | `<<entryPoint>>`, `<<exitPoint>>` on its border     | that machine's `Entry` and listener           |
| Entry / exit behaviour                 | `OnEntry`, `OnExit`                                   | `<onentry>`, `<onexit>`                                              | `X : entry / …`, `X : exit / …`                     | `+ sml::on_entry<sml::_> / …`                 |
| Do activity                            | `Do`                                                  | `<invoke type="tpuml:do" src>`                                       | `X : do / …`                                        | ✗                                             |
| Deferred events                        | `Defer`                                               | ✗                                                                    | `X : ev / defer`                                    | `/ sml::defer`, with `defer_queue`            |
| State invariant                        | `Invariant`                                           | `tpuml:invariant`                                                    | `X : [ cond ]`                                      | `invariant<Name>()` checked after each event  |
| Variables                              | `Variables` on the machine / the state                | `<datamodel>`                                                        | `legend` / `X : name = value`                       | comment; the actions hold them                |
| Trigger, guard, effect                 | `Event`, `Cond`, `Actions`                            | `event`, `cond`, `<script>` holding a name                           | `A --> B : ev [ g ] / act`                          | event method, guard and action methods        |
| Time trigger `after(5s)`               | `After`                                               | `<send delay(expr)>` in `<onentry>`, `<cancel>` on exit              | `after(5s)`                                         | a timer through `FsmTimers`                   |
| Completion transition                  | no `Event`, no `After`                                | eventless, or on `done.state` / `done.invoke`; after a do activity ✗ | unlabelled arrow                                    | anonymous transition                          |
| External / local / internal transition | `Kind`                                                | `type="internal"`, with a target for local                           | `X : ev / act` for internal; local drawn external ⚠ | no target for internal; local ✗               |
| Note                                   | `Note` on the machine, a state, a transition          | `<tpuml:note>`                                                       | `note`, `note on link`                              | `//` comment                                  |
| Note on behaviour/invariant/deferral   | `EntryNote`, `DoNote`, …, `DeferNotes`                | `<tpuml:note about>` on the state                                    | lines of the state's `note`                         | `//` comment before its row                   |
| Stereotype                             | `Stereotype`                                          | `tpuml:stereotype`                                                   | `state "«s»\nX" as X <<s>>`                         | `// «s»` comment                              |

⚠ = written through a workaround that behaves the same, with a warning.
In PlantUML, ⚠ means drawn differently, with a warning and a note in the diagram saying what is meant.
✗ = an error: the output would not behave as the model does, so nothing is written.
Boost.SML resumes a history whenever it enters the state, so entering such a state other than through its history is an error.
A deep history in Boost.SML is an error when the state holds nested states.
Not modelled: signal vs. call events, change events (`when(…)` —
use a guard on a completion transition), protocol state machines.

### Guards and actions are names

Guards, actions, do activities, invariants and named delays name functions, which the generated code calls.
They hold no expressions, so every output language can use them.

- An action or a do activity is a name: letters, digits and `_`, not starting with a digit, such as `addLine`.
- A guard or an invariant is names joined by `not`, `and`, `or` and parentheses, such as
  `card and not (blocked or expired)`.
- A guard may also be `else`, on a choice or junction branch.
- A time trigger's delay is a duration, such as `5s`, `1.5s` or `250ms`, or a name, such as `after(authTimeout)`.

Anything else is an error, whatever the input format: `retries > 3`, `x = 1`, `log('hi')`, `delayexpr="t * 2"`.
In SCXML, an action is a `<script>` holding a name, and any other executable content is an error.

## Requirements

- Go 1.26 or later (see `go.mod`).

## Installation

```bash
go install github.com/vincedupuis/transplantUML/cmd/fsm@latest
```

or from a clone:

```bash
git clone https://github.com/vincedupuis/transplantUML.git
cd transplantUML
make build          # produces ./bin/fsm, or: go build ./cmd/fsm
```

The Makefile also has `run` (`make run ARGS="-i example/coffee-machine.scxml"`), `test`, `fmt`, `vet` and `clean`.

### Tests

`make test` (or `go test ./...`) runs the unit, golden, and round-trip tests. Two tests additionally check the
output against something other than the project itself and skip when their tool is missing:

- **W3C schema** — every SCXML fixture and every document the SCXML emitter writes is validated with `xmllint`
  against the vendored SCXML 1.0 schema in `internal/scxml/testdata/schema/`. `xmllint` ships with macOS and most
  Linux distributions.
- **PlantUML syntax** — the rendered PlantUML is parsed by PlantUML itself (`-syntax`, no Graphviz needed). It uses
  `plantuml` from `PATH` or the jar named by `PLANTUML_JAR`; `make plantuml` downloads the jar into `bin/` and
  `make test` picks it up from there. Needs `java`.
- **Boost.SML compilation** — every document's Boost.SML rendering is compiled with `-Wall -Wextra -Werror`.
  The test compiles it with a program that implements the actions with stubs and sends every event, so SML itself
  checks the tables.
  It uses the header in the directory named by `SML_INCLUDE`; `make sml` downloads it into `bin/` and `make test`
  picks it up from there.
  Needs a C++20 compiler: `CXX`, or `clang++` or `g++` from `PATH`.

## Usage

```
fsm -i input [-f format] [-t template | -F format] [-o output]
```

| Flag                    | Meaning                                                                                               |
|-------------------------|-------------------------------------------------------------------------------------------------------|
| `-i`, `--input`         | Input file (required).                                                                                |
| `-f`, `--input-format`  | Input format: `scxml`, `json`. Default: inferred from the extension (`.scxml`/`.xml`, `.json`).       |
| `-t`, `--template`      | Go template file to render with, or a built-in one: `puml` (the default), `sml`. A file wins.         |
| `-F`, `--output-format` | Write a document format instead of running a template: `scxml`, `json`. Mutually exclusive with `-t`. |
| `-o`, `--output`        | Output file, or the folder for a template that writes several files (`sml`). Default: stdout.         |
| `-h`, `--help`          | Show usage.                                                                                           |

Each flag has a long form; `-F`/`--output-format` is distinct from `-f`/`--input-format` (flags are case-sensitive).

Warnings are printed to stderr as `fsm: warning: …`, one per line, and do not change the exit status:

```
$ fsm -i machine.scxml -F scxml > out.scxml
fsm: warning: state "stop": SCXML has no terminate; written as a <final> state at the top level, which ends the machine
```

What the output cannot run is an error, reported after the warnings, and nothing is written:

```
$ fsm -i machine.scxml -F scxml > out.scxml
fsm: state "work": SCXML has no deferred events
```

### Examples

```bash
# SCXML -> PlantUML with the built-in template
fsm -i example/coffee-machine.scxml -o coffee.puml

# JSON -> C++ state machine on Boost.SML, five files in gen/
fsm -i example/support.json -t sml -o gen/

# SCXML -> your own template
fsm -i example/coffee-machine.scxml -t my-template.gotmpl -o coffee.md

# SCXML -> JSON model, then JSON -> PlantUML (round trip)
fsm -i example/coffee-machine.scxml -F json -o coffee.json
fsm -i coffee.json -o coffee.puml

# JSON -> SCXML, and SCXML -> normalized SCXML
fsm -i coffee.json -F scxml -o coffee.scxml
fsm -i example/coffee-machine.scxml -F scxml
```

`coffee.puml` can be visualized with [PlantUML Online](https://plantuml.online).

### Example documents

[`example/`](example/) holds one document per group of UML concepts, each with the PlantUML `fsm` renders from
it next to it (`<name>.puml`, kept up to date by the tests):

| Document                                                 | Demonstrates                                                                                                          |
|----------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| [`coffee-machine.scxml`](example/coffee-machine.scxml)   | the basics: flat states and event-triggered transitions                                                               |
| [`traffic-light.scxml`](example/traffic-light.scxml)     | time triggers `after(30s)`, variables, entry/exit actions, guarded and internal transitions, notes                    |
| [`order.scxml`](example/order.scxml)                     | choice (recognised without markup), junction, submachine state, do activity, stereotype, terminate                    |
| [`media-player.scxml`](example/media-player.scxml)       | compound state, deep history, entry/exit points, deferred events, invariant, local and internal transitions           |
| [`washing-machine.scxml`](example/washing-machine.scxml) | orthogonal regions, fork and join — and the warnings PlantUML's region limitation produces                            |
| [`thermostat.json`](example/thermostat.json)             | the same concepts written directly in the model's JSON shape                                                          |
| [`kiosk.json`](example/kiosk.json)                       | nesting, behaviours, deferred events, invariants, guards, time triggers, transition kinds, notes                      |
| [`shop.json`](example/shop.json)                         | every state kind (parallel, submachine, all pseudostates, history, named finals) and stereotypes                      |
| [`support.json`](example/support.json)                   | the machine `shop.json` runs in its submachine states: an entry point, an exit point of the machine, a final state    |

Run any of them with `fsm -i example/<name>` and compare with the `.puml` beside it; add `-F scxml` or `-F json`
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
  Name           string     // required
  Initial        string     // top-level initial state
  InitialActions []string   // effect of the initial transition, names
  Variables      []Variable // context attributes
  Note           string
  States         []*State
  Transitions    []*Transition
}

type Variable struct {
  Name  string
  Value string // initial value expression
}

type State struct {
  Name           string
  Parent         string // "" = top level
  Kind           StateKind
  Initial        string   // for compound states
  InitialActions []string // effect of the transition to Initial, names
  OnEntry        []string // entry behaviour, names
  OnExit         []string // exit behaviour, names
  Do             []string // do activity, names
  Defer          []string // deferred events
  Submachine     string   // referenced machine, for submachine states
  Invariant      string   // a condition, like a guard
  Variables      []Variable
  Stereotype     string
  Note           string

  // notes on what the state lists
  EntryNote, ExitNote, DoNote, InvariantNote string
  DeferNotes map[string]string // by deferred event
}

type Transition struct {
  Source  string
  Targets []string       // empty = targetless
  Event   string         // trigger; none and no After = completion transition
  After   string         // time trigger: a duration ("5s", "250ms") or a name
  Cond    string         // guard: names with not, and, or, parentheses; or "else"
  Actions []string       // effect, names
  Kind    TransitionKind // "" = external
  Note    string
}
```

`State` has the predicates `IsNormal`, `IsParallel`, `IsFinal`, `IsTerminate`, `IsHistory`, `IsDeepHistory`,
`IsConnector` (choice, junction, fork, join, entry/exit point) and `IsPseudo` (anything but normal and parallel).
`Transition` has `IsExternal`, `IsLocal`, `IsInternal` and `Trigger()`, which returns the event or `after(delay)`.
A submachine state's `Submachine` is the name of the machine it runs, never a file; `Validate` rejects one holding `.`, `/` or `\`.
`IsName` and `IsCondition` check an action and a guard; `Validate` rejects a model whose guards and actions fail them
(see [Guards and actions are names](#guards-and-actions-are-names)).
`Validate` also rejects a machine without a name, so an SCXML document needs the `name` attribute and a JSON document the `name` key.

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
| `IsReference name`           | whether the state is an entry or exit point of a submachine state        |
| `InitialOf name`             | initial child of a state, or of the machine for `""`                     |
| `InitialActionsOf name`      | effect of that initial transition                                        |
| `Ancestors name`             | parent, grandparent, … of a state, nearest first                         |
| `CommonAncestor name...`     | innermost state containing all the named states (`""` = top level)       |
| `ScopeOf transition`         | innermost state containing a transition's source and targets             |
| `OutgoingTransitions source` | transitions leaving a state                                              |
| `IncomingTransitions target` | transitions entering a state                                             |
| `IsRegion name`              | whether the state is a composite child of a parallel state               |
| `HasTable name`              | whether the state gets a transition table of its own                     |
| `TableOf name`               | the table a state's rows go in (`""` = the machine's)                    |
| `Initials table`             | the states a table starts in, each with its `.State` and `.Scope`        |
| `ForkTarget fork`            | the state a fork's transitions enter together                            |
| `JoinOwner join`             | the parallel state a join's sources lie in, or `""`                      |
| `DefaultEntry outer inner`   | whether entering `outer` without naming a substate reaches `inner`       |
| `AlwaysActive inner outer`   | whether `inner` is active whenever `outer` is                            |
| `Lift transition`            | where a transition is written in the tables, or nil (see below)          |
| `include "name" data`        | like `template`, but returns a string so it can be piped (`\| indent 4`) |
| `prefix p s`                 | `p + s`, or `""` if `s` is empty                                         |
| `surround p s q`             | `p + s + q`, or `""` if `s` is empty                                     |
| `joinNonEmpty sep s...`      | joins the strings, skipping empty ones                                   |
| `warn format args...`        | records a warning for the user (printf-style) and returns `""`           |
| `error format args...`       | records an error (printf-style) and returns `""`; the rendering fails    |
| `file name`                  | starts a file: the output up to the next `file` goes into `name`         |

Call `error` wherever the output would not behave as the model does, and `warn` wherever it gets the same
behaviour through a workaround because the format has no exact counterpart.
The rendering fails with every error the template recorded, after the warnings.

The table functions serve templates that write one transition table per composite state, as Boost.SML does.
A parallel state's regions are flattened into its own table.
`Lift` returns `.From` and `.To`, the states standing for the transition's ends.
A history, an entry point or a fork stands for the state it enters, and a join's source for its parallel state.
`.Source` and `.Target` are those states lifted out of the tables they are nested in, up to `.Table`, the innermost table holding both.
`Lift` returns nil when the transition leaves a history, an entry point, an exit point or a fork.
It also returns nil when the transition reaches an exit point, an entry point with no parent, or a join outside any parallel state.

A template that writes several files calls `file` before each one.
`-o` then names the folder they go in, which is created if needed.
Only white space may come before the first `file`, a name takes no folder, and no name may be used twice.

[`assets/puml.gotmpl`](assets/puml.gotmpl) is the reference template: it shows how to recurse through compound states,
draw parallel regions, and render pseudo-states.

### The built-in PlantUML template

It draws everything in the coverage table above. Things to know:

- Every state is declared under its own name, so notes and behaviours can attach to any of them: final states are
  `<<end>>`, history states `<<history>>` / `<<history*>>`, and `[*]` only marks the initial state. Names PlantUML
  would misread (a dot nests, a dash breaks an arrow) get an alias: `state "a.b" as a_b`.
- A user stereotype is kept as `<<stereotype>>`, which PlantUML uses for skinparams but does not print, and shown
  as `«stereotype»` above the name. Pseudo-states already carry their kind as stereotype, so theirs warns and goes
  into their note.
- Each scope (top level, compound state, region) declares its states before its transitions, and a transition is
  drawn in the innermost scope containing both its ends. PlantUML needs this: an arrow to a nested state that has
  not been declared yet would create a copy at the wrong level.
- PlantUML refuses arrows across the boundary of any parallel region but the first; such transitions are skipped
  with a warning, and the note on their source lists them. Regions are anonymous in PlantUML, so a compound
  region's own transitions, behaviours and note are listed in the note on its parallel state instead.
- A junction is drawn as a filled circle using `<<start>>`, since PlantUML has no junction stereotype. Terminate has
  no symbol either and is drawn as a final state, whose note says it is a terminate. PlantUML ignores description lines on a final state, so its entry
  and exit behaviours are written into its note.
- Local transitions are drawn as external ones marked `«local»`, and a note on an internal transition joins the
  note on its state. Both warn.
- The machine name becomes the diagram `title`; real line breaks in guards and variable values become `\n`.

### The built-in Boost.SML template

`-t sml` writes a state machine in C++20 on [Boost.SML](https://github.com/boost-ext/sml) ([`assets/sml.gotmpl`](assets/sml.gotmpl)).
It writes four files named after the machine, and `FsmTimers.h`, into the folder `-o` names:

```bash
fsm -i internal/render/testdata/sml/kiosk.json -t sml -o gen/
```

| File                | Holds                                                                        |
|---------------------|------------------------------------------------------------------------------|
| `FsmTimers.h`       | `class FsmTimers`, the interface that starts and cancels timers, the same    |
|                     | for every machine                                                            |
| `KioskFsmEvents.h`  | `class KioskFsmEvents`, the interface with one method per event              |
| `KioskFsmActions.h` | `class KioskFsmActions`, the interface with the guards, actions and          |
|                     | invariants to supply                                                         |
| `KioskFsm.h`        | `class KioskFsm`, the state machine, which implements `KioskFsmEvents`       |
|                     | and `class KioskFsmListener`, the interface it tells when it ends or         |
|                     | an invariant does not hold                                                   |
| `KioskFsm.cpp`      | its implementation, the only file that includes SML                          |

That document is `example/kiosk.json` without what Boost.SML cannot run, and
[`internal/render/testdata/sml/kiosk`](internal/render/testdata/sml/kiosk) holds what the template makes of it.
The application implements the actions and sends the events:

```cpp
#include "KioskFsm.h"

class Kiosk : public KioskFsmActions {
    bool inStock() const override;                  // a guard returns bool and is const
    void addLine() override;                        // an action returns nothing
    std::chrono::milliseconds authTimeout() const override;  // a named delay, after(authTimeout)
    // ...
};

Kiosk kiosk;
Timers timers;                                      // implements FsmTimers
KioskFsm fsm{kiosk, timers};
fsm.enterFsm();
fsm.touch();
```

The machine waits, ignoring every event, until `enterFsm()` starts it in its initial state.
`enterFsm(KioskFsm::Entry::p)` starts it at its entry point `p` instead.
Either one starts the machine again from the beginning if it is already running.
`stopFsm()` leaves the current state, running the exit behaviours of the active states, and makes it wait again.
`terminateFsm()` makes it wait again at once, running no exit behaviour, as a terminate state does.
These names keep them apart from the events' methods; an event named like one of them is an error, since the files
would not compile.

`setListener` takes a `KioskFsmListener`, which the machine tells when it ends, and when an invariant does not hold.
`onFinished()` is called when it reaches its final state.
`onTerminated()` is called when it reaches a terminate state, at any depth, once the current event is done.
`onExitP()` is called when it leaves by its exit point `p`.
`onInvariantViolated("paying")` is called when the invariant of the active state `paying` does not hold.
The machine checks it once `enterFsm` or an event is processed, while the state is active.
Each name of the invariant is a method of the actions interface, named after it: `[basket and not card]` calls
`invariantBasket()` and `invariantCard()`, each a `bool … const`.
The machine stops before telling it.

### Time triggers

A machine with time triggers takes an `FsmTimers` in its constructor.
`FsmTimers.h` is the same for every machine, so one implementation serves them all:

```cpp
class Timers : public FsmTimers {
public:
    Id startTimer(std::chrono::milliseconds delay, std::function<void()> fire) override;
    void cancelTimer(Id id) override;
};
```

- Each state has a timer per delay of its time triggers.
  Entering the state starts it, after the state's entry behaviour.
  Leaving the state cancels it, before the state's exit behaviour.
- An orthogonal state starts and cancels the timers of its regions.
- `fire` makes the machine process the timer's event, which only the timer can send.
  Call it on the thread that sends the machine its events, and never from within `startTimer` or `cancelTimer`.
- A timer that fires after its state was left, or after the machine was stopped or terminated, does nothing.
  So cancelling a timer need not be exact.
- `after(90s)` is written as `std::chrono::seconds{90}`, and `after(1.5s)` as `std::chrono::milliseconds{1500}`.
  A delay finer than a millisecond is rounded down, with a warning.
- A named delay, `after(authTimeout)`, is a method of the actions interface, read each time the timer starts.

### Submachine states

A submachine state runs another machine, generated from its own document with `-t sml`.
The submachine state's reference is the name of the machine it runs: `payment` gives `PaymentFsm`.
The constructor takes one instance of it per submachine state, and each state needs its own:

```cpp
Support support;                                    // implements SupportFsmActions
Shop shop;                                          // implements ShopFsmActions
SupportFsm helpdesk{support};
SupportFsm aftersales{support};
ShopFsm fsm{shop, helpdesk, aftersales};            // shop.json's helpdesk and aftersales
fsm.enterFsm();

fsm.help();            // enters helpdesk, which starts its SupportFsm
helpdesk.human();      // the submachine's events go to its instance
```

- Entering the submachine state runs its entry behaviour, then starts the machine with `enterFsm`, at the entry point
  the transition named, if any.
- The machine's listener is the submachine state.
  When the machine finishes, the submachine state takes its completion transition.
  When it leaves by an exit point, the submachine state takes that point's transition.
- Leaving the submachine state on an event of its own stops the machine with `stopFsm`, which runs the exit
  behaviours of its active states, then runs the submachine state's exit behaviour.
- A terminate state in the machine terminates the submachine state's machine too, without exit behaviours, and it
  tells its own listener `onTerminated()`.
- The machine may finish while it is being entered.
  The outer machine then queues what it reports and handles it once the current event is done.
- `ShopFsm.h` only declares `class SupportFsm;`.
  `ShopFsm.cpp` includes `SupportFsm.h`, so the build needs both machines' files.
- `shop.json` is generated without reading `support.json`.
  An entry or exit point that `support` does not declare shows up as a compile error.

Things to know:

- The names come from the machine's name in PascalCase, `my-shop` giving `MyShopFsm`.
- A guard `[inStock]` calls `inStock()` on the actions, and an action `/ addLine` calls `addLine()`.
  A name used both ways is one method that returns `bool` and is not `const`.
- A guard made of names, `not`, `and`, `or` and parentheses becomes SML's `!`, `&&` and `||`.
- Inside `KioskFsm.cpp`, the SML code sits in an anonymous namespace.
  Events are empty structs in `namespace event`.
  Simple states and pseudostates are string-literal states, `"idle"_s`.
  A compound or orthogonal state is a struct of its own, used as `sml::state<ordering>`.
  An orthogonal state's regions share its table, with one initial state each.
- A final state is `sml::X`, and a completion transition is SML's anonymous transition.
  SML takes an anonymous transition leaving a composite state once the composite state reaches `X`, which is UML's
  completion.
- SML has no transitions into or out of a composite state's inside.
  A completion transition that leaves from inside goes to `X` instead, and the composite state's own completion
  transition takes it on, with a warning.
  A transition that leaves a region is written for its orthogonal state, which is active exactly when the region is,
  with a warning.
  A transition into the inside enters the composite state, with a warning when that reaches the target through
  initial states, and an error otherwise.
  Any other transition leaving from inside is an error, since the composite state's would fire in all its states.
- A choice or a junction is a state that anonymous transitions leave at once, tried in order, `else` last.
- SML's history marks a region's initial state, `"browsing"_s(sml::H)`, so every entry into that region resumes it.
  A transition that enters the state other than through its history is therefore an error.
  So is a default transition that leads elsewhere than the initial state, or that has an effect.
  A deep history is written as a shallow one, with a warning, when the state holds no nested states, and is an
  error otherwise.
- The machine waits in `sml::state<internal::stopped>`, which `enterFsm` leaves with the effect of the initial
  transition, `/ boot = "idle"_s`.
  Inside a state, an effect on the initial transition goes through a transient state,
  `*"outer.initial"_s / greet = "inner"_s`.
- Deferred events use SML's `defer_queue` policy.
- A time trigger `after(90s)` in the state `ordering` fires the event `internal::ordering_after_90s`, through a
  timer (see [Time triggers](#time-triggers)).
- A join ends each region that reaches it in `X`.
  Its outgoing transition becomes the orthogonal state's completion transition, taken once every region has ended.
  It warns, and it is an error when its incoming transitions do not all leave the regions of one orthogonal state.
- A fork enters its orthogonal state, with a warning, when the regions start in the states it leads to and it has
  no effect; otherwise it is an error.
- Notes and stereotypes become `//` comments in `KioskFsm.cpp`, as does each invariant above the rows that check it,
  the machine's note goes on the class,
  and its variables are listed on the actions interface, whose implementation holds them.
  SML has no do activities, so one is an error, as are a multi-target transition, a local transition, and a
  region's own behaviours.
- The machine's own entry points are the values of `Entry`, and its exit points are states that end it.
  A transition to its entry point, or from its exit point, is an error.
  A state's entry point enters the state, with a warning, when that reaches the state it leads to through initial
  states and it has no guard or effect; otherwise it is an error, as is a state's exit point.
- A terminate state is `sml::state<internal::terminated>`, which its region stays in.
  Entering it flags the machine, which ends once the current event is done.
  In an orthogonal state, the other regions may still act on that event first, so it warns there.

## SCXML

### Native elements

`<state>`, `<parallel>`, `<final>`, `<history type="shallow|deep">`, `initial` attribute and `<initial>` element,
`<transition>` (`event`, `cond`, multiple `target`s, `type="internal"`), `<onentry>`, `<onexit>`, `<datamodel>` with
`<data>` (on the machine and on states, as variables) and `<invoke>`.
An action is a `<script>` holding a name, `<script>addLine</script>`.
Any other executable content (`<log>`, `<assign>`, `<raise>`, `<if>`, `<script src>`, …) is an error, as is a
`<send>` or `<cancel>` that is not a time trigger.
SCXML elements with no place in the model (`<donedata>`, a top-level `<script>`, …) are an error.
An element of another namespace, which SCXML engines ignore, is dropped with a warning.

`<invoke>` is a **submachine state** when it invokes another SCXML machine (`type` absent or `scxml`), whose name is
its `src`: `<invoke src="payment"/>`.
It is a **do activity** otherwise, named by its `src`: `<invoke type="tpuml:do" src="spin"/>`.
Another `type` is not kept, with a warning.
An `<invoke>` with children or without `src`, and a second SCXML `<invoke>` in one state, are errors.

### Idioms recognised as UML

- A `<state>` with no content of its own (no children, actions, invokes, or data) and only eventless transitions is
  a **choice** when it has several transitions and at least one guard, and a **fork** when it has one transition
  with several targets. Add `tpuml:kind="normal"` to keep such a state as a state.
- A `<send event="E" delay="D">` in `<onentry>` whose event `E` is consumed by a `<transition>` of the same state
  is a **time trigger**: the transition gets `After: D` and the `<send>` (and a matching `<cancel>` in `<onexit>`)
  are not treated as actions. `delayexpr` counts as well, and the emitter writes `After` back to whichever of the
  two fits: `delay` when it is a time value such as `5s` or `250ms`, `delayexpr` when it is anything the engine has
  to evaluate.
- A transition on `done.state.S` in the state `S` is a **completion transition**, since SCXML raises that event when
  a compound or parallel state completes. So is one on `done.invoke.I` in a submachine state whose `<invoke>` has
  the id `I`, and one on a bare `done.invoke` there when the submachine is the state's only invoke. The emitter
  writes a completion transition back on those events, giving the submachine's `<invoke>` the id `S.submachine`.

### The tpuml extension vocabulary

SCXML permits attributes and elements from other namespaces everywhere, and engines ignore them. `fsm` uses the
namespace `https://github.com/vincedupuis/transplantUML` (any prefix; `tpuml` below) for what SCXML cannot say:

| Extension                                                                    | On                                   | Meaning                                   |
|------------------------------------------------------------------------------|--------------------------------------|-------------------------------------------|
| `tpuml:kind="junction\|join\|entry-point\|exit-point\|choice\|fork\|normal"` | `<state>`                            | the pseudo-state kind (a transient state) |
| `tpuml:kind="terminate"`                                                     | `<final>`                            | a terminate pseudo-state                  |
| `tpuml:kind="initial"`                                                       | `<state>` the `<scxml>` starts in    | holds the machine's initial transition    |
| `tpuml:kind="local"`                                                         | `<transition>`                       | a local transition, as `type="internal"`  |
| `tpuml:defer="e1 e2"`                                                        | `<state>`                            | deferred events                           |
| `tpuml:invariant="expr"`                                                     | `<state>`                            | state invariant                           |
| `tpuml:stereotype="name"`                                                    | `<state>`                            | stereotype                                |
| `<tpuml:note>text</tpuml:note>`                                              | `<scxml>`, `<state>`, `<transition>` | a note                                    |
| `<tpuml:note about="entry\|exit\|do\|invariant">text</tpuml:note>`           | `<state>`                            | a note on that behaviour or the invariant |
| `<tpuml:note about="defer" event="e">text</tpuml:note>`                      | `<state>`                            | a note on deferring `e`                   |

See [`internal/scxml/testdata/uml.scxml`](internal/scxml/testdata/uml.scxml) for a document using all of them.

### Writing SCXML (`-F scxml`)

Everything the model holds is written, natively where SCXML has the element and through the extension vocabulary
for what does not change how an engine runs the machine (kinds, notes, stereotypes, invariants); the namespace is declared only when it is used. The document validates against the W3C schema. It is
equivalent, not byte-identical, to the one it came from:

- the initial child is written as an `initial` attribute, or as an `<initial>` element when its transition has an
  effect; `<scxml>` takes no `<initial>` element, so the machine's own initial effect goes in a transient state
  `tpuml:kind="initial"` that the machine starts in and leaves at once, with a warning;
- each action becomes a `<script>` holding its name, and each do activity `<invoke type="tpuml:do" src="name"/>`;
- a time trigger becomes `<send event="after.D" delay="D" id="…">` in `<onentry>`, the matching `<cancel>` in
  `<onexit>`, and a transition on that event;
- a completion transition waits for `done.state.S` on a compound or parallel state and `done.invoke.S.submachine`
  on a submachine state; elsewhere it stays eventless, which SCXML takes at once, so a simple state with a do
  activity is an error;
- a local transition is written as `type="internal"`, which SCXML runs as UML's local transition when the source is
  compound and every target lies inside it; any other local transition is an error;
- comments and anything the parser dropped are not preserved.

A terminate at the top level is written as a `<final>` state, which ends the machine the same way, with a warning.
So is the machine's initial effect, written in a transient state.
What SCXML would run differently is an error: join (the first region to reach it would leave the parallel state),
a terminate inside a state or with exit actions, deferred events, a completion transition that cannot wait for a do
activity, and the entry and exit points of a submachine state.
