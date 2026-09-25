package fsm

import (
	"os"
	"path/filepath"
	"testing"
)

// The generated parser accepts the grammar's own idioms: nested states
// interleaved with the events that concern them, entry/exit actions, guards,
// action lists, and every kind of goto.
func TestGrammarAccepts(t *testing.T) {
	src, err := os.ReadFile("testdata/coffee.fsm")
	if err != nil {
		t.Fatal(err)
	}
	tree, err := parse(string(src))
	if err != nil {
		t.Fatalf("syntax errors: %v", err)
	}
	if got := tree.Identifier().GetText(); got != "coffee" {
		t.Errorf("machine name = %q, want coffee", got)
	}
	if got := len(tree.AllState()); got != 2 {
		t.Errorf("top-level states = %d, want 2", got)
	}
}

func TestGrammarRejects(t *testing.T) {
	for _, src := range []string{
		"fsm {}",                                 // missing name
		"fsm m { state s { entry } }",            // a behaviour needs its actions
		"fsm m { state s { exit } }",             //
		"fsm m { state s { do } }",               //
		"fsm m { state s { entry goto s } }",     // entry cannot leave the state
		"fsm m { state s { on entry / a } }",     // a behaviour is not a trigger
		"fsm m { state s { do goto s } }",        // a do activity is not a transition
		"fsm m { state s { on e [g] / defer } }", // no guard on a deferred event
		"fsm m { state s { defer e } }",          // a deferred event names its event first
		"fsm m { state s { on e [] } }",          // empty guard
		"fsm m { state s { on e goto } }",        // goto without a target
		"fsm m { state s { on e } }",             // neither effect nor target
		"fsm m { state s { on e [g] } }",         // a guard alone is not a transition
		"fsm m { state s { on e goto a/b } }",    // a target is a name, not a path
		"fsm m { state s { after(5s) } }",        // neither effect nor target
		"fsm m { state s { after 5s goto s } }",  // a delay sits in parentheses
		"fsm m { state s { after(5) goto s } }",  // a delay carries its unit
		"fsm m { state s { after(5m) goto s } }", // and the unit is ms or s
		"fsm m { state s { after() goto s } }",   // a delay is not optional
		"fsm m { state s { on after(5s) goto s } }",
		"fsm m { initial }", // initial modifies a state
		"fsm m { state s { initial on e / a } }",
		"fsm m { state initial {} }", // initial is a keyword, not a name
		"fsm m { state s { / a } }",  // a completion transition needs its goto
		"fsm m { state s { [g] } }",
		"fsm m { state s { on e [else] goto s } }",                         // else guards only a choice or junction branch
		"fsm m { state s { [else] goto s } }",                              //
		"fsm m { parallel state p { state s {} } }",                        // a parallel state holds regions
		"fsm m { parallel p { region r {} } }",                             // and is written "parallel state"
		"fsm m { region r {} }",                                            // a region sits only in a parallel state
		"fsm m { state s { region r {} } }",                                //
		"fsm m { parallel state p { region r { on e goto r } } }",          // and has no events
		"fsm m { parallel state p { region r { entry point e goto r } } }", // nor points
		"fsm m { initial region r {} }",
		"fsm m { initial fork f { goto f } }",                       // a fork cannot be the starting child
		"fsm m { initial [g] goto a state a {} }",                   // the initial transition has no guard
		"fsm m { initial on e goto a state a {} }",                  // nor a trigger
		"fsm m { initial goto final }",                              // and names a state
		"fsm m { parallel state p { initial goto r region r {} } }", // a parallel state starts in all its regions
		"fsm m { choice c { on e goto c } }",                        // a branch has no trigger
		"fsm m { choice c on e goto c }",                            // on its own line either
		"fsm m { junction j { after(1s) goto j } }",
		"fsm m { fork f { [g] goto f } }", // no guard leaving a fork
		"fsm m { fork f { on e goto f } }",
		"fsm m { join j { goto j } }",        // a join is one line
		"fsm m { join j [g] goto j }",        // with no guard
		"fsm m { join j }",                   // and a goto
		"fsm m { entry point e { goto e } }", // a point is one line too
		"fsm m { entry point e [g] goto e }",
		"fsm m { exit point e }",
		"fsm m { point e goto e }",
		"fsm m { state s { on e goto s.x } }", // after a dot comes H or H*
		"fsm m { state s { on e goto .H } }",
		"fsm m { state s { on e goto s.final } }",
		"fsm m { state s { on e goto H.s } }",
		"fsm m { state s { on e goto . } }",           // a state names itself
		"fsm m { state s { on e goto local final } }", // final and terminate sit outside the source
		"fsm m { state s { on e goto local terminate } }",
		"fsm m { state s { on e local goto s } }", // local qualifies the target
		"fsm m { state local {} }",
		"fsm m { state parallel {} }", // the new keywords are not names
		"fsm m { state region {} }",
		"fsm m { state choice {} }",
		"fsm m { state junction {} }",
		"fsm m { state fork {} }",
		"fsm m { state join {} }",
		"fsm m { state point {} }",
		"fsm m { state terminate {} }",
		"fsm m { state else {} }",
		"fsm m { state s { on e [else or g] goto s } }",
		"fsm m { state submachine {} }",
		"fsm m { state invariant {} }",
		"fsm m { state s | n | {} }",                                // a note comes before its directive
		"fsm m { state s { on e goto s | n | } }",                   //
		"fsm m { | a | | b | state s {} }",                          // once
		"fsm m { state s { | n | } }",                               // and before something
		"fsm m { | n  state s {} }",                                 // between two bars
		"| a | | b | fsm m {}",                                      //
		"fsm m { state s { on e | n | goto s } }",                   // not inside a clause
		"fsm m { state s <<>> {} }",                                 // a stereotype names something
		"fsm m { state s <<a b>> {} }",                              // one thing
		"fsm m { state s <<a, b>> {} }",                             //
		"fsm m { state s <<a> {} }",                                 // between << and >>
		"fsm m { state s <<\"a\">> {} }",                            // as a name
		"fsm m { <<a>> state s {} }",                                // after the declared name
		"fsm m { state <<a>> s {} }",                                //
		"fsm m { initial state s {} <<a>> }",                        // not after the body
		"fsm m { state s <<a>> <<b>> {} }",                          // at most once
		"fsm m { state s { on e goto t <<a>> } }",                   // on a declaration only
		"fsm m { state s { invariant } }",                           // an invariant is a condition
		"fsm m { state s { invariant [] } }",                        //
		"fsm m { state s { invariant g } }",                         // in brackets
		"fsm m { state s { invariant [g] / a } }",                   // and never a transition
		"fsm m { state s { invariant [else] } }",                    // else guards only a branch
		"fsm m { state s { on e invariant [g] } }",                  // an invariant is no trigger
		"fsm m { parallel state p { region r { invariant [g] } } }", // a region holds only states
		"fsm m { submachine s { state t {} } }",                     // a submachine state's states are the machine's
		"fsm m { submachine s { choice c goto s } }",                // and so are its pseudostates
		"fsm m { submachine s { exit point e } }",                   // its exit point leaves by a goto
		"fsm m { state s { entry point e } }",                       // only a submachine's entry point has none
		"fsm m { submachine state s {} }",                           // submachine replaces state
		"fsm m { submachine s }",                                    // and keeps its braces
		"fsm m { parallel state p { submachine s {} } }",            // a parallel state holds regions only
		"fsm m { H goto a state a {} }",                             // a history belongs to a state
		"fsm m { parallel state p { H goto p region r {} } }",       // or a region, never a parallel state
		"fsm m { submachine s { H goto s } }",                       // nor a submachine state
		"fsm m { state s { H / a } }",                               // a history default ends in a goto
		"fsm m { state s { H goto s <<a>> } }",                      // its stereotype follows H
		"fsm m { final state f {} }",                                // a final state has no behaviours
		"fsm m { final state f { entry / a } }",                     //
		"fsm m { final f }",                                         // and is written "final state"
		"fsm m { initial final state f }",                           // a scope never starts in it
		"fsm m { final state f.g }",                                 // its name is a name
		"fsm m { parallel state p { final state f region r {} } }",  // a parallel state holds regions only
		"fsm m { submachine s { final state f } }",                  // a submachine state's states are the machine's
		"fsm m { terminate t }",                                     // a terminate is written "terminate state"
		"fsm m { terminate state t {} }",                            // on one line
		"fsm m { initial terminate state t }",                       //
		"fsm m { parallel state p { terminate state region r {} } }",
		"fsm m { state s { s.H goto s } }", // it is declared inside its state
	} {
		if _, err := parse(src); err == nil {
			t.Errorf("%q: accepted, want a syntax error", src)
		}
	}
}

// Two alternatives cover every combination of guard, actions and goto that
// carries an effect or a target, and three more the behaviours of a state,
// the events it defers and its invariant; TestGrammarRejects covers the rest.
func TestEventForms(t *testing.T) {
	for _, body := range []string{
		"on e / a",
		"on e / a, b",
		"on e goto t",
		"on e [g] / a",
		"on e [g] goto t",
		"on e / a goto t",
		"on e [g] / a goto t",
		"entry / a",
		"exit / a, b",
		"do / poll",
		"do / poll, refresh",
		"on pause / defer",
		"invariant [g]",
		"| n | on e goto t",
		"| multi\n line | on e / a",
		"| bar \\| | goto t",
		"invariant [not g or (h and i)]",
		"after(5s) / a",
		"after(250ms) goto t",
		"after(1.5s) [g] / a goto t",
		"after(retryDelay) goto t",
		"goto t",
		"[g] goto t",
		"/ a goto t",
		"[g] / a goto t",
	} {
		if _, err := parse("fsm m { state s { " + body + " } state t {} }"); err != nil {
			t.Errorf("%q: %v", body, err)
		}
	}
}

// Every goto target form lexes: a state name, one of the keywords standing
// for a state the document never declares, or the history of a named state.
func TestGotoTargets(t *testing.T) {
	for _, target := range []string{
		"final",
		"terminate",
		"H",
		"H*",
		"t",
		"t.H",
		"t.H*",
		"local t",
		"local H",
		"local t.H*",
	} {
		if _, err := parse("fsm m { state s { on e goto " + target + " } }"); err != nil {
			t.Errorf("goto %s: %v", target, err)
		}
	}
}

// The documents in example/ are what the README points a reader at, so they
// have to parse.
func TestExamplesParse(t *testing.T) {
	paths, err := filepath.Glob("../../example/*.fsm")
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) == 0 {
		t.Fatal("no examples found")
	}
	for _, path := range paths {
		src, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := parse(string(src)); err != nil {
			t.Errorf("%s: %v", path, err)
		}
	}
}

// "initial" marks the child a compound state starts in, and the machine's own
// starting state at the top level. It is a named lexer rule declared ahead of
// Identifier, so a state cannot be called "initial" and the visitor gets an
// Initial() accessor.
func TestInitialState(t *testing.T) {
	tree, err := parse("fsm m { initial state a { initial state b {} state c {} } state d {} }")
	if err != nil {
		t.Fatalf("syntax errors: %v", err)
	}
	top := tree.AllState()
	if top[0].INITIAL() == nil {
		t.Error("a: not marked initial")
	}
	if top[1].INITIAL() != nil {
		t.Error("d: marked initial")
	}
	inner := top[0].AllState()
	if inner[0].INITIAL() == nil {
		t.Error("b: not marked initial")
	}
	if inner[1].INITIAL() != nil {
		t.Error("c: marked initial")
	}
}

// A final state, a terminate and a history are one-line declarations, named
// or not, and a history's default is optional. Line breaks carry no meaning,
// so a clause after a bare H is the declaring state's own.
func TestEndingsAccepted(t *testing.T) {
	for _, src := range []string{
		"fsm m { final state }",
		"fsm m { final state done }",
		"| n | fsm m { | n | final state done <<ok>> | n | terminate state <<alarm>> }",
		"fsm m { state s { final state terminate state abort final state f <<x>> } }",
		"fsm m { parallel state p { region r { final state terminate state t } } }",
		"fsm m { state s { H } }",
		"fsm m { state s { | n | H* <<kept>> } }",
		"fsm m { state s { H <<kept>> / a goto s } }",
		"fsm m { state s { H on e goto s } }",
		"fsm m { state s { H [g] goto s } }",
	} {
		if _, err := parse(src); err != nil {
			t.Errorf("%q: %v", src, err)
		}
	}
}

// A note may stand before every directive: the machine, each declaration,
// each clause, branch and fork line.
func TestNotesAccepted(t *testing.T) {
	src := `| m | fsm m {
		| s | initial state s { | e | entry / a | t | on e goto c }
		| p | parallel state p { | r | region r { | i | initial state i {} } | x | exit point x goto s }
		| u | submachine u {}
		| c | choice c { | b | [g] goto s | o | [else] goto s }
		| k | junction k | b | goto s
		| f | fork f { | l | goto i }
		| j | join j goto s
	}`
	if _, err := parse(src); err != nil {
		t.Fatal(err)
	}
}

// A name is a letter or an underscore followed by letters, digits and
// underscores. The punctuation other formats allow in an id, dots and hyphens,
// is not part of a name here, so "goto ." stays a keyword and no document can
// declare the synthesized "<scope>.final" or "<state>.H".
func TestIdentifierCharset(t *testing.T) {
	src := "fsm vending_machine { state _idle1 { on coin_in [hasChange] / open_1 goto payNow } state payNow {} }"
	if _, err := parse(src); err != nil {
		t.Errorf("charset: %v", err)
	}
	for _, src := range []string{
		"fsm m { state .a {} }",
		"fsm m { state -a {} }",
		"fsm m { state 1a {} }",
		"fsm vending.machine { state a {} }",
		"fsm m { state pay.now {} }",
		"fsm m { state pay-now {} }",
		"fsm m { state a { on coin.in / g goto a } }",
		"fsm m { state a { on e [has-change] / g goto a } }",
		"fsm m { state a { on e / g goto pay.now } state pay.now {} }",
		"fsm m { state a { after(retry.delay) goto a } }",
	} {
		if _, err := parse(src); err == nil {
			t.Errorf("%q: accepted, want a syntax error", src)
		}
	}
}
