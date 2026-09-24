# The fsm DSL against UML state machines

This compares the fsm language (`internal/fsm/fsm.g4`) with UML 2.5.1, chapter 14 (State Machines).
Every supported example parses with the `fsm` command.
The UML 2.5.1 column gives the sections of the [specification](https://www.omg.org/spec/UML/2.5.1/PDF) that define each concept.
Most are in chapter 14, but events are in chapter 13, comments, names and constraints in chapter 7, and stereotypes in chapter 12.

Legend: ✅ supported · ⚠️ partial · ❌ not in the DSL

## State machine

| UML concept                                                | UML 2.5.1          | DSL                                                                             | Example                                             |
|------------------------------------------------------------|--------------------|---------------------------------------------------------------------------------|-----------------------------------------------------|
| State machine (named)                                      | 14.2.3.1           | ✅                                                                              | `fsm kiosk { … }`                                   |
| Comment attached to an element                             | 7.2.3.2            | ✅ The note goes before the element.                                            | `\| Self-service kiosk. \| fsm kiosk { … }`         |
| Connection points on the machine (entry/exit point)        | 14.2.3.1, 14.2.3.7 | ✅                                                                              | `entry point reorder / load goto checkout`          |
| Several top-level regions                                  | 14.2.3.2           | ⚠️ There is no top-level `region`, so you wrap the regions in a parallel state. | `parallel state main { region a {…} region b {…} }` |
| Several machines in one document (to define a submachine)  | 14.2.3.4.7         | ❌ One `fsm` per document. A submachine is only referenced by name.             | —                                                   |
| Machine parameters, context classifier                     | 13.2.3.2, 13.2.3.4 | ❌                                                                              | —                                                   |
| Redefinition / extension (`{extended}`, `isLeaf`)          | 14.3               | ❌                                                                              | —                                                   |
| Protocol state machine (`{protocol}`, pre/post conditions) | 14.4               | ❌                                                                              | —                                                   |

## States

| UML concept                                            | UML 2.5.1              | DSL                                                                                 | Example                                                                                         |
|--------------------------------------------------------|------------------------|-------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| Simple state                                           | 14.2.3.4.1             | ✅                                                                                  | `state idle {}`                                                                                 |
| Composite state                                        | 14.2.3.4.1             | ✅                                                                                  | `state ordering { initial state browsing {} state paying {} }`                                  |
| Orthogonal state + regions                             | 14.2.3.4.1, 14.2.3.2   | ✅ Only inside `parallel state`.                                                    | `parallel state ship { region w { initial state pack {} } region a { initial state bill {} } }` |
| Submachine state                                       | 14.2.3.4.7             | ✅ It is named after the referenced machine.                                        | `submachine support { on done goto browsing }`                                                  |
| Final state (unnamed / named)                          | 14.2.3.6               | ✅                                                                                  | `goto final` · `final state refunded`                                                           |
| Entry behavior                                         | 14.2.3.4.3             | ✅                                                                                  | `entry / dim, beep`                                                                             |
| Exit behavior                                          | 14.2.3.4.3             | ✅                                                                                  | `exit / clearBasket`                                                                            |
| Do activity                                            | 14.2.3.4.3             | ✅                                                                                  | `do / spin`                                                                                     |
| State invariant                                        | 14.5.9                 | ✅ Only one per state.                                                              | `invariant [basket and not cash]`                                                               |
| Deferrable trigger                                     | 14.2.3.4.4, 14.2.4.8.6 | ⚠️ Named events only. You cannot defer a time event.                                | `on touch / defer`                                                                              |
| Stereotype                                             | 12.3.3.4               | ⚠️ One stereotype per state, no tagged values.                                      | `state checkout <<secure>> { … }`                                                               |
| Same name in different scopes (qualified names `A::B`) | 7.4.3.2, 7.4.4.2       | ❌ All names share one namespace, so names must be unique across the whole machine. | —                                                                                               |
| State redefinition `{extended}`                        | 14.3.3.3               | ❌                                                                                  | —                                                                                               |

## Pseudostates

| UML concept                                        | UML 2.5.1            | DSL                                    | Example                                                                         |
|----------------------------------------------------|----------------------|----------------------------------------|---------------------------------------------------------------------------------|
| Initial (marker)                                   | 14.2.3.7             | ✅                                     | `initial state browsing { … }`                                                  |
| Initial transition with an effect                  | 14.2.3.7             | ✅                                     | `initial / boot goto idle`                                                      |
| Initial → choice/junction                          | 14.2.3.7             | ✅                                     | `initial choice depth { [regular] goto quick  [else] goto thorough }`           |
| Shallow history                                    | 14.2.3.4.4, 14.2.3.7 | ✅                                     | `on resume goto H` · `goto checkout.H`                                          |
| Deep history                                       | 14.2.3.4.4, 14.2.3.7 | ✅                                     | `goto checkout.H*`                                                              |
| History default transition                         | 14.2.3.7             | ✅                                     | `H / greet goto browsing`                                                       |
| History note/stereotype                            | 7.2.3.2, 12.3.3.4    | ✅                                     | `\| back where we left \| H`                                                    |
| Choice (guards evaluated on arrival)               | 14.2.3.7             | ✅                                     | `choice route { [large] goto review  [else] goto paying }`                      |
| Junction (static branch/merge)                     | 14.2.3.7             | ✅                                     | `junction paid / receipt goto split`                                            |
| Fork                                               | 14.2.3.7             | ✅                                     | `fork split { goto packing  / notify goto invoicing }`                          |
| Join                                               | 14.2.3.7             | ✅                                     | `join merge / close goto done` (and in each region: `goto merge`)               |
| Entry point on a composite state                   | 14.2.3.7             | ✅                                     | `state checkout { entry point express / useCard goto paying … }`                |
| Exit point on a composite state                    | 14.2.3.7             | ✅                                     | `exit point cancelled goto browsing`                                            |
| Connection point reference (on a submachine state) | 14.2.3.5             | ✅ An entry point there has no `goto`. | `submachine support { entry point urgent  exit point escalated goto checkout }` |
| Terminate                                          | 14.2.3.7             | ✅                                     | `on fraud goto terminate` · `terminate state`                                   |

## Transitions

| UML concept                                   | UML 2.5.1              | DSL                                                                                         | Example                                                    |
|-----------------------------------------------|------------------------|---------------------------------------------------------------------------------------------|------------------------------------------------------------|
| External transition                           | 14.2.3.8.1             | ✅                                                                                          | `on checkout goto paying`                                  |
| Local transition                              | 14.2.3.8.1             | ✅ The target must be inside the source.                                                    | `on back goto local browsing`                              |
| Internal transition                           | 14.2.3.8.1             | ✅ Leave out `goto`.                                                                        | `on add / addLine`                                         |
| Self transition                               | 14.2.3.8.1             | ✅                                                                                          | `on remove / dropLine goto browsing` (inside `browsing`)   |
| Completion transition                         | 14.2.3.8.3             | ✅ Leave out the trigger.                                                                   | `goto final` · `[paidInFull] goto merge`                   |
| Guard                                         | 14.2.3.8, 7.6          | ⚠️ Boolean `and`/`or`/`not`/parentheses over bare names only. No comparisons or arithmetic. | `on add [inStock and (card or cash)] / addLine`            |
| `else` guard                                  | 14.2.3.7               | ✅ Only on a choice or a junction.                                                          | `[else] goto paying`                                       |
| Effect                                        | 14.2.3.8               | ⚠️ A list of bare names. Actions take no arguments and there are no assignments.            | `on approved / receipt, beep goto done`                    |
| Signal / call event                           | 13.3.3.2               | ✅                                                                                          | `on touch goto browsing`                                   |
| Relative time event (`after`)                 | 13.3.3.4               | ⚠️ Units are `ms` and `s` only, or a named delay.                                           | `after(90s) goto idle` · `after(authTimeout) goto again`   |
| Absolute time event (`at`)                    | 13.3.3.4               | ❌                                                                                          | —                                                          |
| Change event (`when`)                         | 13.3.3.3               | ❌                                                                                          | —                                                          |
| Any-receive event (`all`)                     | 13.3.3.2               | ❌                                                                                          | —                                                          |
| Several triggers on one transition (`e1, e2`) | 14.2.4.8               | ❌ Write one clause per event.                                                              | `on cash goto paid` + `on card goto paid`                  |
| Event parameters (`e(x)`)                     | 13.3.3.2, 14.2.4.8     | ❌                                                                                          | —                                                          |
| Transition into a nested state / out of one   | 14.2.3.4.5, 14.2.3.4.6 | ✅ Names are global.                                                                        | `on resume goto authorizing`                               |
| Transition to an entry/exit point             | 14.2.3.7               | ✅                                                                                          | `on help goto urgent`                                      |
| Compound transition (via a junction)          | 14.2.3.8.4             | ✅                                                                                          | `on cash goto paid` → `junction paid / receipt goto split` |
| Several targets on one transition             | 14.2.3.8.4             | ⚠️ Only through an explicit `fork`.                                                         | `fork f { goto a  goto b }`                                |
| Note on a transition                          | 7.2.3.2                | ✅                                                                                          | `\| Nobody touched the screen. \| after(90s) goto idle`    |
| Note on entry/exit/do/defer/invariant         | 7.2.3.2                | ✅                                                                                          | `\| Until the bank answers. \| do / spin`                  |
| Stereotype on a transition                    | 12.3.3.4               | ❌                                                                                          | —                                                          |

## Summary

- All of UML's vertex kinds are covered: every state kind, every pseudostate, history defaults and connection point references.
- All three transition kinds are covered: external, local and internal.
- The gaps are on the trigger side: several triggers per transition, `when`, `at`, `all`, and event parameters.
- The next gap is the expression side: guards cannot compare values, and actions take no arguments.
- The machine level is also missing parameters, redefinition, protocol machines, and more than one machine per document.
- UML scopes names by region, so the same state name can appear in different places. The DSL rejects this.
- The model has a `Variables` field, but the DSL has no syntax to fill it. UML has no direct equivalent either, since a machine's data belongs to its context classifier.
