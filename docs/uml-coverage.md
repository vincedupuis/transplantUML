# The fsm DSL against UML state machines

This compares tpuml's fsm language (`internal/fsm/fsm.g4`) with UML 2.5.1, chapter 14 (State Machines).
Every supported example parses with `tpuml`.

Legend: ✅ supported · ⚠️ partial · ❌ not in the DSL

## State machine

| UML concept                                                | DSL                                                                             | Example                                             |
|------------------------------------------------------------|---------------------------------------------------------------------------------|-----------------------------------------------------|
| State machine (named)                                      | ✅                                                                              | `fsm kiosk { … }`                                   |
| Comment attached to an element                             | ✅ The note goes before the element.                                            | `\| Self-service kiosk. \| fsm kiosk { … }`         |
| Connection points on the machine (entry/exit point)        | ✅                                                                              | `entry point reorder / load goto checkout`          |
| Several top-level regions                                  | ⚠️ There is no top-level `region`, so you wrap the regions in a parallel state. | `parallel state main { region a {…} region b {…} }` |
| Several machines in one document (to define a submachine)  | ❌ One `fsm` per document. A submachine is only referenced by name.             | —                                                   |
| Machine parameters, context classifier                     | ❌                                                                              | —                                                   |
| Redefinition / extension (`{extended}`, `isLeaf`)          | ❌                                                                              | —                                                   |
| Protocol state machine (`{protocol}`, pre/post conditions) | ❌                                                                              | —                                                   |

## States

| UML concept                                            | DSL                                                                                 | Example                                                                                         |
|--------------------------------------------------------|-------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| Simple state                                           | ✅                                                                                  | `state idle {}`                                                                                 |
| Composite state                                        | ✅                                                                                  | `state ordering { initial state browsing {} state paying {} }`                                  |
| Orthogonal state + regions                             | ✅ Only inside `parallel state`.                                                    | `parallel state ship { region w { initial state pack {} } region a { initial state bill {} } }` |
| Submachine state                                       | ✅ It is named after the referenced machine.                                        | `submachine support { on done goto browsing }`                                                  |
| Final state (unnamed / named)                          | ✅                                                                                  | `goto final` · `final state refunded`                                                           |
| Entry behavior                                         | ✅                                                                                  | `entry / dim, beep`                                                                             |
| Exit behavior                                          | ✅                                                                                  | `exit / clearBasket`                                                                            |
| Do activity                                            | ✅                                                                                  | `do / spin`                                                                                     |
| State invariant                                        | ✅ Only one per state.                                                              | `invariant [basket and not cash]`                                                               |
| Deferrable trigger                                     | ⚠️ Named events only. You cannot defer a time event.                                | `on touch / defer`                                                                              |
| Stereotype                                             | ⚠️ One stereotype per state, no tagged values.                                      | `state checkout <<secure>> { … }`                                                               |
| Same name in different scopes (qualified names `A::B`) | ❌ All names share one namespace, so names must be unique across the whole machine. | —                                                                                               |
| State redefinition `{extended}`                        | ❌                                                                                  | —                                                                                               |

## Pseudostates

| UML concept                                        | DSL                                    | Example                                                                         |
|----------------------------------------------------|----------------------------------------|---------------------------------------------------------------------------------|
| Initial (marker)                                   | ✅                                     | `initial state browsing { … }`                                                  |
| Initial transition with an effect                  | ✅                                     | `initial / boot goto idle`                                                      |
| Initial → choice/junction                          | ✅                                     | `initial choice depth { [regular] goto quick  [else] goto thorough }`           |
| Shallow history                                    | ✅                                     | `on resume goto H` · `goto checkout.H`                                          |
| Deep history                                       | ✅                                     | `goto checkout.H*`                                                              |
| History default transition                         | ✅                                     | `H / greet goto browsing`                                                       |
| History note/stereotype                            | ✅                                     | `\| back where we left \| H`                                                    |
| Choice (guards evaluated on arrival)               | ✅                                     | `choice route { [large] goto review  [else] goto paying }`                      |
| Junction (static branch/merge)                     | ✅                                     | `junction paid / receipt goto split`                                            |
| Fork                                               | ✅                                     | `fork split { goto packing  / notify goto invoicing }`                          |
| Join                                               | ✅                                     | `join merge / close goto done` (and in each region: `goto merge`)               |
| Entry point on a composite state                   | ✅                                     | `state checkout { entry point express / useCard goto paying … }`                |
| Exit point on a composite state                    | ✅                                     | `exit point cancelled goto browsing`                                            |
| Connection point reference (on a submachine state) | ✅ An entry point there has no `goto`. | `submachine support { entry point urgent  exit point escalated goto checkout }` |
| Terminate                                          | ✅                                     | `on fraud goto terminate` · `terminate state`                                   |

## Transitions

| UML concept                                   | DSL                                                                                         | Example                                                    |
|-----------------------------------------------|---------------------------------------------------------------------------------------------|------------------------------------------------------------|
| External transition                           | ✅                                                                                          | `on checkout goto paying`                                  |
| Local transition                              | ✅ The target must be inside the source.                                                    | `on back goto local browsing`                              |
| Internal transition                           | ✅ Leave out `goto`.                                                                        | `on add / addLine`                                         |
| Self transition                               | ✅                                                                                          | `on remove / dropLine goto browsing` (inside `browsing`)   |
| Completion transition                         | ✅ Leave out the trigger.                                                                   | `goto final` · `[paidInFull] goto merge`                   |
| Guard                                         | ⚠️ Boolean `and`/`or`/`not`/parentheses over bare names only. No comparisons or arithmetic. | `on add [inStock and (card or cash)] / addLine`            |
| `else` guard                                  | ✅ Only on a choice or a junction.                                                          | `[else] goto paying`                                       |
| Effect                                        | ⚠️ A list of bare names. Actions take no arguments and there are no assignments.            | `on approved / receipt, beep goto done`                    |
| Signal / call event                           | ✅                                                                                          | `on touch goto browsing`                                   |
| Relative time event (`after`)                 | ⚠️ Units are `ms` and `s` only, or a named delay.                                           | `after(90s) goto idle` · `after(authTimeout) goto again`   |
| Absolute time event (`at`)                    | ❌                                                                                          | —                                                          |
| Change event (`when`)                         | ❌                                                                                          | —                                                          |
| Any-receive event (`all`)                     | ❌                                                                                          | —                                                          |
| Several triggers on one transition (`e1, e2`) | ❌ Write one clause per event.                                                              | `on cash goto paid` + `on card goto paid`                  |
| Event parameters (`e(x)`)                     | ❌                                                                                          | —                                                          |
| Transition into a nested state / out of one   | ✅ Names are global.                                                                        | `on resume goto authorizing`                               |
| Transition to an entry/exit point             | ✅                                                                                          | `on help goto urgent`                                      |
| Compound transition (via a junction)          | ✅                                                                                          | `on cash goto paid` → `junction paid / receipt goto split` |
| Several targets on one transition             | ⚠️ Only through an explicit `fork`.                                                         | `fork f { goto a  goto b }`                                |
| Note on a transition                          | ✅                                                                                          | `\| Nobody touched the screen. \| after(90s) goto idle`    |
| Note on entry/exit/do/defer/invariant         | ✅                                                                                          | `\| Until the bank answers. \| do / spin`                  |
| Stereotype on a transition                    | ❌                                                                                          | —                                                          |

## Summary

- All of UML's vertex kinds are covered: every state kind, every pseudostate, history defaults and connection point references.
- All three transition kinds are covered: external, local and internal.
- The gaps are on the trigger side: several triggers per transition, `when`, `at`, `all`, and event parameters.
- The next gap is the expression side: guards cannot compare values, and actions take no arguments.
- The machine level is also missing parameters, redefinition, protocol machines, and more than one machine per document.
- UML scopes names by region, so the same state name can appear in different places. The DSL rejects this.
- The model has a `Variables` field, but the DSL has no syntax to fill it. UML has no direct equivalent either, since a machine's data belongs to its context classifier.
