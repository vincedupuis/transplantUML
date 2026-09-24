# The fsm DSL syntax

This is a syntax synopsis of the fsm language, in the style of a man page or a language reference.
The synopsis covers syntax only.

## Conventions

| Notation      | Meaning                                        |
|---------------|------------------------------------------------|
| `[ x ]`       | x is optional.                                 |
| `x ...`       | x is repeated zero or more times.              |
| `a \| b`      | a or b.                                        |
| `<x>`         | A building block defined below.                |
| Anything else | Literal, including `{ }`, `( )` and `"[" "]"`. |

## Building blocks

| Building block | Syntax                                                                        |
|----------------|-------------------------------------------------------------------------------|
| `<name>`       | letters, digits and `_`                                                       |
| `<note>`       | `\|free text\|`                                                               |
| `<stereotype>` | `<<<name>>>`                                                                  |
| `<duration>`   | `250ms` \| `1.5s`                                                             |
| `<expr>`       | `<name> \| not <expr> \| <expr> and <expr> \| <expr> or <expr> \| ( <expr> )` |
| `<guard>`      | `"[" <expr> "]"`                                                              |
| `<actions>`    | `/ <name> [, <name> ...]`                                                     |
| `<target>`     | `<name> \| final \| terminate \| [<name>.]H \| [<name>.]H*`                   |
| `<goto>`       | `goto [local] <target>`                                                       |
| `<trigger>`    | `on <name> \| after ( <duration> \| <name> )`                                 |
| comment        | `# text`                                                                      |

## Declarations

### Machine

```
[<note>]
fsm <name> {
    <member> ...
}
```

```
| A self-service kiosk:
  browse, pay, take the receipt. |
fsm kiosk {
  initial / boot goto idle
  state idle { on touch / wake goto browsing }
}
```

### Initial transition

```
initial [<actions>] goto <name>
```

```
initial / boot goto idle
```

This line is the only way to give the initial transition an effect.
Without an effect, you can put `initial` in front of the state's declaration instead.

### State

```
[<note>]
[initial] state <name> [<stereotype>] {
    <member> ...
}
```

```
| The basket is locked from here on. |
state paying {
  entry / lockBasket
  on approved / receipt goto done

  initial state authorizing {
    on cancel goto browsing
  }
}
```

A leaf state has an empty body: `state quick {}`.

### Parallel state and region

```
[<note>]
[initial] parallel state <name> [<stereotype>] {
    <region> | <point> | <clause> ...
}

[<note>]
region <name> [<stereotype>] {
    <member> ...
}
```

```
parallel state shipping {
  region warehouse {
    initial state packing { on packed goto packed }
    state packed { goto merge }
  }
  region accounting {
    initial state invoicing { on sent goto sent }
    state sent { [paidInFull] goto merge }
  }
}
```

### Submachine

```
[<note>]
[initial] submachine <name> [<stereotype>] {
    <point> | <reference> | <clause> ...
}

<reference> = [<note>] entry point <name> [<stereotype>]
```

```
submachine support {
  entry / openChat
  entry point urgent
  exit point escalated / page goto checkout
  goto browsing
}
```

The name is the name of the machine it refers to.
An `entry point` without a `goto` is a connection point reference.

### Final and terminate

```
[<note>]
final state [<name>] [<stereotype>]

[<note>]
terminate state [<name>] [<stereotype>]
```

```
| Money returned. |
final state refunded

| Fraud closes the shop. |
terminate state
```

With a name, the line declares a named final or terminate state.
Without a name, it annotates the one that `goto final` or `goto terminate` reaches in that scope.

### History

```
[<note>]
H | H* [<stereotype>] [ [<actions>] <goto> ]
```

```
| Back to the step the customer left. |
H

H / greet goto browsing
```

The `goto` is the history state's default transition.

### Choice and junction

```
[<note>]
[initial] choice <name> [<stereotype>] {
    <branch> ...
}

[<note>]
[initial] choice <name> [<stereotype>] <branch>

<branch> = [<note>] [ "[" else "]" | <guard> ] [<actions>] <goto>
```

`junction` has the same two forms.

```
choice route {
  [large] goto review
  [else] goto paying
}

junction paid / receipt goto split
```

### Fork and join

```
[<note>]
fork <name> [<stereotype>] {
    [<note>] [<actions>] <goto> ...
}

[<note>]
join <name> [<stereotype>] [<actions>] <goto>
```

```
fork split {
  goto packing
  / notify goto invoicing
}

join merge / close goto done
```

### Entry and exit points

```
[<note>]
entry point | exit point <name> [<stereotype>] [<actions>] <goto>
```

```
entry point reorder / loadBasket goto checkout
exit point cancelled goto browsing
```

A point can be declared on the machine, on a state, or on a parallel state.

## Clauses

### Behaviours

```
[<note>]
entry | exit | do <actions>
```

```
entry / lockBasket
exit / clearBasket, unlock
| Until the bank answers. |
do / spin
```

### Deferred event

```
[<note>]
on <name> / defer
```

```
on touch / defer
```

### Invariant

```
[<note>]
invariant <guard>
```

```
invariant [basket and not (card and cash)]
```

### Transition with a trigger

```
[<note>]
<trigger> [<guard>] [<actions>] [<goto>]
```

It needs `<actions>`, `<goto>`, or both.

```
on touch / wake goto browsing                    # external
on back goto local browsing                      # local: the target is inside the source
on add [inStock] / addLine                       # internal: no goto
| Nobody touched the screen. |
after(90s) goto idle                             # time trigger with a fixed delay
after(authTimeout) [retries] / warn goto again   # time trigger with a named delay
```

### Completion transition

```
[<note>]
[<guard>] [<actions>] <goto>
```

```
goto browsing
[paidInFull] goto merge
```

This transition has no trigger.
It fires when the state finishes.

### Targets

```
goto <name> | final | terminate | H | H* | <name>.H | <name>.H*
```

```
on ack goto final
on fraud goto terminate
on resume goto H
on reopen goto checkout.H*
```

## Where each member is allowed

| Scope      | Allowed members                                                                                      |
|------------|------------------------------------------------------------------------------------------------------|
| `fsm`      | initial, state, parallel, submachine, final, terminate, choice, junction, fork, join, point, clause  |
| `state`    | the same as `fsm`, plus `H` / `H*`                                                                   |
| `region`   | initial, state, parallel, submachine, final, terminate, choice, junction, fork, join, `H` / `H*`     |
| `parallel` | region, point, clause                                                                                |

