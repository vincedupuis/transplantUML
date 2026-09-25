package render

import (
	"reflect"
	"testing"

	"github.com/vincedupuis/transplantUML/internal/model"
)

// idle forks into the regions of p, which join back to done. r1 is a region
// with states of its own, r2 a leaf that stands for its region. c has a history.
func tablesSample() *model.StateMachine {
	return &model.StateMachine{
		Initial: "idle",
		States: []*model.State{
			{Name: "idle", Kind: model.Normal},
			{Name: "fk", Kind: model.Fork},
			{Name: "p", Kind: model.Parallel},
			{Name: "r1", Parent: "p", Kind: model.Normal, Initial: "a1"},
			{Name: "a1", Parent: "r1", Kind: model.Normal},
			{Name: "a2", Parent: "r1", Kind: model.Normal},
			{Name: "r2", Parent: "p", Kind: model.Normal},
			{Name: "j", Kind: model.Join},
			{Name: "done", Kind: model.Final},
			{Name: "c", Kind: model.Normal, Initial: "c1"},
			{Name: "c1", Parent: "c", Kind: model.Normal},
			{Name: "h", Parent: "c", Kind: model.HistoryShallow},
		},
		Transitions: []*model.Transition{
			{Source: "idle", Targets: []string{"fk"}, Event: "go"},
			{Source: "fk", Targets: []string{"a1"}},
			{Source: "fk", Targets: []string{"r2"}},
			{Source: "a2", Targets: []string{"j"}},
			{Source: "r2", Targets: []string{"j"}},
			{Source: "j", Targets: []string{"done"}},
			{Source: "c1", Targets: []string{"idle"}, Event: "e"},
			{Source: "idle", Targets: []string{"h"}, Event: "back"},
			{Source: "a1", Event: "tick", Actions: []string{"count"}},
		},
	}
}

func TestTables(t *testing.T) {
	sm := tablesSample()
	if err := sm.Validate(); err != nil {
		t.Fatal(err)
	}
	x := tables{sm}
	for name, want := range map[string]bool{"r1": true, "r2": false, "p": false, "a1": false, "c": false} {
		if got := x.IsRegion(name); got != want {
			t.Errorf("IsRegion(%q) = %v", name, got)
		}
	}
	for name, want := range map[string]bool{"p": true, "c": true, "r1": false, "r2": false, "idle": false} {
		if got := x.HasTable(name); got != want {
			t.Errorf("HasTable(%q) = %v", name, got)
		}
	}
	for name, want := range map[string]string{"a1": "p", "r2": "p", "c1": "c", "p": ""} {
		if got := x.TableOf(name); got != want {
			t.Errorf("TableOf(%q) = %q, want %q", name, got, want)
		}
	}
	if got, want := x.Initials("p"), []Initial{{"a1", "r1"}, {"r2", "p"}}; !reflect.DeepEqual(got, want) {
		t.Errorf("Initials(p) = %v, want %v", got, want)
	}
	if got, want := x.Initials(""), []Initial{{"idle", ""}}; !reflect.DeepEqual(got, want) {
		t.Errorf("Initials() = %v, want %v", got, want)
	}
	if got := x.ForkTarget("fk"); got != "p" {
		t.Errorf("ForkTarget(fk) = %q", got)
	}
	if got := x.JoinOwner("j"); got != "p" {
		t.Errorf("JoinOwner(j) = %q", got)
	}
	want := []*Lift{
		{Table: "", Source: "idle", Target: "p", From: "idle", To: "p"},
		nil,
		nil,
		{Table: "p", Source: "a2", Target: "j", From: "a2", To: "j"},
		{Table: "p", Source: "r2", Target: "j", From: "r2", To: "j"},
		{Table: "", Source: "p", Target: "done", From: "p", To: "done"},
		{Table: "", Source: "c", Target: "idle", From: "c1", To: "idle"},
		{Table: "", Source: "idle", Target: "c", From: "idle", To: "c"},
		{Table: "p", Source: "a1", From: "a1"},
	}
	for i, tr := range sm.Transitions {
		if got := x.Lift(tr); !reflect.DeepEqual(got, want[i]) {
			t.Errorf("Lift(%s -> %v) = %+v, want %+v", tr.Source, tr.Targets, got, want[i])
		}
	}
}
