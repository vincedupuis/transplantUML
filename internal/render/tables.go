package render

import "github.com/vincedupuis/transplantUML/internal/model"

// tables holds the functions for templates that write a state machine as
// transition tables, one per composite state, the way C++ libraries such as
// Boost.SML do: a parallel state's regions are flattened into its own table,
// and a transition goes in the innermost table holding both its ends.
type tables struct{ sm *model.StateMachine }

// IsRegion reports whether the named state is a region with states of its
// own: a composite child of a parallel state. A table flattens it into the
// parallel state's.
func (x tables) IsRegion(name string) bool {
	s := x.sm.State(name)
	if s == nil || !s.IsNormal() || s.Submachine != "" || s.Parent == "" || len(x.sm.Children(name)) == 0 {
		return false
	}
	p := x.sm.State(s.Parent)
	return p != nil && p.IsParallel()
}

// HasTable reports whether the named state gets a table of its own: a
// parallel state, or a composite state that is not a region.
func (x tables) HasTable(name string) bool {
	s := x.sm.State(name)
	if s == nil || x.IsRegion(name) || s.Submachine != "" {
		return false
	}
	return s.IsParallel() || (s.IsNormal() && len(x.sm.Children(name)) > 0)
}

// TableOf returns the table the named state's rows go in: its parent's, or
// the parallel state's when its parent is a region. "" is the machine's.
func (x tables) TableOf(name string) string {
	s := x.sm.State(name)
	if s == nil {
		return ""
	}
	if x.IsRegion(s.Parent) {
		return x.sm.State(s.Parent).Parent
	}
	return s.Parent
}

// Initial is a state a table starts in, and the scope whose initial state it
// is: the table's own, or that of one of its regions.
type Initial struct {
	State string
	Scope string
}

// Initials returns the states the table starts in: one per region of a
// parallel state, else the initial state of the table's scope.
func (x tables) Initials(table string) []Initial {
	var out []Initial
	if s := x.sm.State(table); s != nil && s.IsParallel() {
		for _, r := range x.sm.Children(table) {
			if !x.IsRegion(r.Name) {
				out = append(out, Initial{r.Name, table})
			} else if r.Initial != "" {
				out = append(out, Initial{r.Initial, r.Name})
			}
		}
	} else if i := x.sm.InitialOf(table); i != "" {
		out = append(out, Initial{i, table})
	}
	return out
}

// ForkTarget returns the state a fork's transitions enter together: the
// innermost state holding all their targets, a region counting as its
// parallel state.
func (x tables) ForkTarget(fork string) string {
	var targets []string
	for _, t := range x.sm.OutgoingTransitions(fork) {
		targets = append(targets, t.Targets...)
	}
	return x.outsideRegion(x.sm.CommonAncestor(targets...))
}

// JoinOwner returns the parallel state a join's incoming transitions all
// come from, or "" when there is none.
func (x tables) JoinOwner(join string) string {
	var sources []string
	for _, t := range x.sm.IncomingTransitions(join) {
		sources = append(sources, t.Source)
	}
	if s := x.sm.State(x.outsideRegion(x.sm.CommonAncestor(sources...))); s != nil && s.IsParallel() {
		return s.Name
	}
	return ""
}

func (x tables) outsideRegion(name string) string {
	if x.IsRegion(name) {
		return x.sm.State(name).Parent
	}
	return name
}

// Lift says where a transition is written. From and To are the states that
// stand for its ends: the source, or the parallel state a join synchronises;
// the first target, or the state a history, entry point or fork enters (a
// join stays itself, as the end of its region). Source and Target are From
// and To lifted out of the tables they are nested in, up to Table, the
// innermost table holding both. Target and To are "" for a transition
// without targets.
type Lift struct {
	Table, Source, Target, From, To string
}

// Lift returns where the transition is written, or nil when it is not: a
// transition leaving a history, entry or exit point or fork, or reaching an
// exit point, an entry point with no parent, or a join with no owner.
func (x tables) Lift(t *model.Transition) *Lift {
	s := x.sm.State(t.Source)
	from, to := t.Source, ""
	skip := s.IsHistory() || s.Kind == model.EntryPoint || s.Kind == model.ExitPoint || s.Kind == model.Fork
	if s.Kind == model.Join {
		from = x.JoinOwner(t.Source)
		skip = from == ""
	}
	var toChain []link
	if len(t.Targets) == 0 {
		toChain = x.chain(from)
	} else {
		tg := x.sm.State(t.Targets[0])
		to = tg.Name
		switch {
		case tg.IsHistory():
			to = tg.Parent
		case tg.Kind == model.EntryPoint:
			to = tg.Parent
			skip = skip || to == ""
		case tg.Kind == model.ExitPoint:
			skip = true
		case tg.Kind == model.Fork:
			to = x.ForkTarget(tg.Name)
		case tg.Kind == model.Join:
			if owner := x.JoinOwner(tg.Name); owner != "" {
				toChain = append([]link{{tg.Name, owner}}, x.chain(owner)...)
			} else {
				skip = true
			}
		}
		if toChain == nil && !skip {
			toChain = x.chain(to)
		}
	}
	if skip {
		return nil
	}
	for _, sl := range x.chain(from) {
		for _, tl := range toChain {
			if sl.table == tl.table {
				l := &Lift{Table: sl.table, Source: sl.state, From: from, To: to}
				if len(t.Targets) > 0 {
					l.Target = tl.state
				}
				return l
			}
		}
	}
	return nil
}

// link is a state and the table it is written in.
type link struct{ state, table string }

// chain returns the state and each ancestor that is not a region, nearest
// first, with the table each is written in.
func (x tables) chain(name string) []link {
	var out []link
	for _, n := range append([]string{name}, x.sm.Ancestors(name)...) {
		if !x.IsRegion(n) {
			out = append(out, link{n, x.TableOf(n)})
		}
	}
	return out
}
