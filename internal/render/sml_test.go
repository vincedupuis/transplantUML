package render

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/vincedupuis/transplantUML/assets"
	"github.com/vincedupuis/transplantUML/internal/model"
)

// smlGoldens maps the documents whose Boost.SML rendering is checked in to the
// folder holding its files. kiosk, shop and uml are the examples and uml.scxml
// without what Boost.SML cannot run (TestSMLErrors). Regenerate one with:
// go run ./cmd/fsm -i <input> -t sml -o internal/render/testdata/sml/<name>
var smlGoldens = map[string]string{
	"testdata/sml/kiosk.json":    "testdata/sml/kiosk",
	"testdata/sml/shop.json":     "testdata/sml/shop",
	"../../example/support.json": "testdata/sml/support",
	"testdata/sml/uml.scxml":     "testdata/sml/uml",
}

// smlFiles renders sm with the SML template and splits the output into its
// files.
func smlFiles(t *testing.T, sm *model.StateMachine) ([]File, model.Warnings) {
	t.Helper()
	out, warnings, err := Render(sm, assets.SML)
	if err != nil {
		t.Fatal(err)
	}
	files, err := Files(out)
	if err != nil {
		t.Fatal(err)
	}
	return files, warnings
}

// smlFilesOf parses path with the parser its extension names and renders it
// with the SML template.
func smlFilesOf(t *testing.T, path string) ([]File, model.Warnings) {
	t.Helper()
	out, warnings := renderWith(t, path, assets.SML)
	files, err := Files(out)
	if err != nil {
		t.Fatal(err)
	}
	return files, warnings
}

func fileNames(files []File) []string {
	var names []string
	for _, f := range files {
		names = append(names, f.Name)
	}
	return names
}

func TestSMLGolden(t *testing.T) {
	for input, golden := range smlGoldens {
		t.Run(filepath.Base(input), func(t *testing.T) {
			files, _ := smlFilesOf(t, input)
			entries, err := os.ReadDir(golden)
			if err != nil {
				t.Fatal(err)
			}
			var want []string
			for _, e := range entries {
				want = append(want, e.Name())
			}
			if got := fileNames(files); !reflect.DeepEqual(slices.Sorted(slices.Values(got)), want) {
				t.Errorf("files = %v, want those in %s: %v", got, golden, want)
			}
			for _, f := range files {
				want, err := os.ReadFile(filepath.Join(golden, f.Name))
				if err == nil && f.Content != string(want) {
					t.Errorf("%s differs from %s\n--- got ---\n%s--- want ---\n%s", f.Name, golden, f.Content, want)
				}
			}
		})
	}
}

// The files come in the order an include needs them, named after the machine.
func TestSMLFiles(t *testing.T) {
	files, _ := smlFiles(t, &model.StateMachine{Name: "my-coffee machine"})
	want := []string{"FsmTimers.h", "MyCoffeeMachineFsmEvents.h", "MyCoffeeMachineFsmActions.h", "MyCoffeeMachineFsm.h", "MyCoffeeMachineFsm.cpp"}
	if got := fileNames(files); !reflect.DeepEqual(got, want) {
		t.Errorf("files = %v, want %v", got, want)
	}
}

// The SML template says where it writes the model through a workaround that
// behaves the same.
func TestSMLWarnings(t *testing.T) {
	_, warnings := smlFilesOf(t, "testdata/sml/uml.scxml")
	want := []string{
		`state "in": Boost.SML has no entry points; transitions to it enter "outer", which starts in the state the entry point leads to`,
		`state "split": Boost.SML has no fork; transitions to it enter "both", whose regions start in the states it forks to`,
		`state "sync": Boost.SML has no join; transitions to it end their region (X), and its outgoing transition is taken once every region of "both" has ended`,
	}
	if !reflect.DeepEqual([]string(warnings), want) {
		t.Errorf("warnings =\n%s\nwant\n%s", strings.Join(warnings, "\n"), strings.Join(want, "\n"))
	}
}

// What Boost.SML cannot run the way the model means it is an error, and
// rendering fails with all of them.
func TestSMLErrors(t *testing.T) {
	documents := map[string][]string{
		"../scxml/testdata/uml.scxml": {
			`state "out": Boost.SML has no exit points`,
			`state "work": Boost.SML has no do activities, so runJob cannot run`,
			`transition outer -> inner: Boost.SML has no local transitions`,
		},
		"../scxml/testdata/edge.scxml": {
			`transition a -> a b: a Boost.SML transition has one target`,
		},
		"../../example/kiosk.json": {
			`state "authorizing": Boost.SML has no do activities, so spin cannot run`,
			`transition paying -> again: Boost.SML has no transitions into a composite state, and entering "paying" would not reach "again"`,
			`transition authorizing -> browsing: Boost.SML has no transitions from inside a composite state; as a transition of "paying" it would fire from any of its states`,
			`state "ordering.H": Boost.SML resumes a history in its region's initial state the first time, not through the default transition to browsing`,
			`transition idle -> browsing: Boost.SML has no transitions into a composite state, and entering "ordering" would not reach "browsing"`,
			`transition ordering -> browsing: Boost.SML has no transitions into a composite state, and entering "ordering" would not reach "browsing"`,
			`transition ordering -> browsing: Boost.SML has no local transitions`,
			`transition paying -> done: Boost.SML has no transitions from inside a composite state; as a transition of "ordering" it would fire from any of its states`,
		},
		"../../example/shop.json": {
			`state "express": Boost.SML has no entry points; entering "checkout" would not reach the state the entry point leads to, or would drop its guard or actions`,
			`state "cancelled": Boost.SML has no exit points`,
			`state "checkout.H-deep": Boost.SML has only shallow history, and "checkout" holds nested states a deep history would resume`,
			`transition reorder -> checkout: "checkout" has a history, which Boost.SML would resume instead of entering it at its initial state`,
			`transition browsing -> checkout: "checkout" has a history, which Boost.SML would resume instead of entering it at its initial state`,
			`transition escalated -> checkout: "checkout" has a history, which Boost.SML would resume instead of entering it at its initial state`,
			`state "split": Boost.SML has no fork; entering "shipping" would not start its regions in the states the fork leads to, or would drop the fork's actions`,
		},
	}
	for path, want := range documents {
		t.Run(filepath.Base(path), func(t *testing.T) {
			smlErrors(t, parseFile(t, path), want)
		})
	}

	models := map[string]struct {
		sm   *model.StateMachine
		want []string
	}{
		// The machine's own entry point is reached only by enterFsm, and it
		// ends at its exit point. An event named like a method that starts or
		// stops the machine, a named delay like a guard, or an invariant's
		// method like a guard would not compile.
		"points and names": {&model.StateMachine{
			Name: "m", Initial: "a",
			States: []*model.State{
				{Name: "a", Kind: model.Normal, Invariant: "open"},
				{Name: "in", Kind: model.EntryPoint},
				{Name: "out", Kind: model.ExitPoint},
			},
			Transitions: []*model.Transition{
				{Source: "a", Targets: []string{"a"}, Event: "shut", Cond: "invariantOpen"},
				{Source: "in", Targets: []string{"a"}},
				{Source: "a", Targets: []string{"in"}, Event: "enterFsm"},
				{Source: "a", Targets: []string{"out"}, Event: "leave"},
				{Source: "out", Targets: []string{"a"}},
				{Source: "a", Targets: []string{"a"}, After: "ok", Cond: "ok"},
			},
		}, []string{
			`event "enterFsm": its method enterFsm clashes with the one that starts or stops the machine; the files would not compile`,
			`state "in": the machine is entered at its entry point only by enterFsm, so a transition cannot reach it`,
			`state "out": the machine ends at its exit point, so no transition can leave it`,
			`delay "ok": its method ok clashes with a guard or action of that name; the files would not compile`,
			`invariant method invariantOpen: it clashes with a guard, action or delay of that name; the files would not compile`,
		}},
		// A region has no behaviour of its own, X none either, and a
		// completion can be lifted out of one composite state only.
		"regions, finals and depth": {&model.StateMachine{
			Name: "m", Initial: "p",
			States: []*model.State{
				{Name: "p", Kind: model.Parallel},
				{Name: "r", Parent: "p", Kind: model.Normal, Initial: "c", OnEntry: []string{"wake"}},
				{Name: "c", Parent: "r", Kind: model.Normal, Initial: "d"},
				{Name: "d", Parent: "c", Kind: model.Normal, Initial: "e"},
				{Name: "e", Parent: "d", Kind: model.Normal},
				{Name: "end", Kind: model.Final, OnEntry: []string{"bye"}},
			},
			Transitions: []*model.Transition{
				{Source: "e", Targets: []string{"end"}},
			},
		}, []string{
			`region "r": Boost.SML regions are anonymous, so they have no behaviours, invariants or deferred events`,
			`transition e -> end: Boost.SML has no transitions from inside a composite state, and "e" lies too deep inside "p" to complete it`,
			`state "end": Boost.SML's final state X has no entry or exit behaviour`,
		}},
		// A join needs its incoming transitions to leave the regions of one
		// orthogonal state.
		"join without owner": {&model.StateMachine{
			Name: "m", Initial: "a",
			States: []*model.State{
				{Name: "a", Kind: model.Normal},
				{Name: "b", Kind: model.Normal},
				{Name: "j", Kind: model.Join},
			},
			Transitions: []*model.Transition{
				{Source: "a", Targets: []string{"j"}, Event: "x"},
				{Source: "b", Targets: []string{"j"}, Event: "y"},
				{Source: "j", Targets: []string{"a"}},
			},
		}, []string{
			`state "j": Boost.SML has no join, and its incoming transitions do not all come from one orthogonal state`,
		}},
	}
	for name, c := range models {
		t.Run(name, func(t *testing.T) {
			if err := c.sm.Validate(); err != nil {
				t.Fatal(err)
			}
			smlErrors(t, c.sm, c.want)
		})
	}
}

// smlErrors checks that rendering sm with the SML template fails with want.
func smlErrors(t *testing.T, sm *model.StateMachine, want []string) {
	t.Helper()
	out, _, err := Render(sm, assets.SML)
	if err == nil {
		t.Fatalf("rendered without an error:\n%s", out)
	}
	if got := strings.Split(err.Error(), "\n"); !reflect.DeepEqual(got, want) {
		t.Errorf("errors =\n%s\nwant\n%s", err, strings.Join(want, "\n"))
	}
}

// A terminate ends the machine once the current event is done, which in an
// orthogonal state lets the other regions act on it first.
func TestSMLTerminateInRegionWarns(t *testing.T) {
	sm := &model.StateMachine{
		Name: "m", Initial: "p",
		States: []*model.State{
			{Name: "p", Kind: model.Parallel},
			{Name: "r1", Parent: "p", Kind: model.Normal, Initial: "a"},
			{Name: "a", Parent: "r1", Kind: model.Normal},
			{Name: "t", Parent: "r1", Kind: model.Terminate},
			{Name: "r2", Parent: "p", Kind: model.Normal},
		},
		Transitions: []*model.Transition{{Source: "a", Targets: []string{"t"}, Event: "e"}},
	}
	if err := sm.Validate(); err != nil {
		t.Fatal(err)
	}
	_, warnings := smlFiles(t, sm)
	want := []string{`state "t": Boost.SML finishes the current event in the other regions before the machine terminates`}
	if !reflect.DeepEqual([]string(warnings), want) {
		t.Errorf("warnings = %q, want %q", warnings, want)
	}
}

// clashes is a machine whose names collide in C++: a guard called event, which
// would hide the events' namespace, an action called checkout like a compound
// state, a compound state called machine like the machine's own table, a name
// used as a guard in one table and as a guard and an action in another, a
// state name holding a quote, and a guarded row that starts its table.
var clashes = &model.StateMachine{
	Name: "my-shop", Initial: "pick",
	States: []*model.State{
		{Name: "pick", Kind: model.Choice},
		{Name: `say "hi"`, Kind: model.Normal},
		{Name: "checkout", Kind: model.Normal, Initial: "waiting"},
		{Name: "waiting", Parent: "checkout", Kind: model.Normal},
		{Name: "paid", Parent: "checkout", Kind: model.Normal},
		{Name: "machine", Kind: model.Normal, Initial: "m1"},
		{Name: "m1", Parent: "machine", Kind: model.Normal},
	},
	Transitions: []*model.Transition{
		{Source: "pick", Targets: []string{"checkout"}, Cond: "event"},
		{Source: "pick", Targets: []string{`say "hi"`}, Cond: "else"},
		{Source: `say "hi"`, Targets: []string{"checkout"}, Event: "checkout", Actions: []string{"checkout"}},
		{Source: "waiting", Targets: []string{"paid"}, Event: "card", Cond: "card and not blocked", Actions: []string{"charge", "log"}},
		{Source: "paid", Targets: []string{"waiting"}, Event: "retry", Actions: []string{"blocked"}},
		{Source: "checkout", Targets: []string{"machine"}, Event: "go", Cond: "blocked"},
		{Source: "m1", Event: "tick", Actions: []string{"log"}},
	},
}

func TestSMLNames(t *testing.T) {
	if err := clashes.Validate(); err != nil {
		t.Fatal(err)
	}
	files, warnings := smlFiles(t, clashes)
	if len(warnings) != 0 {
		t.Errorf("warnings = %q", warnings)
	}
	content := map[string]string{}
	for _, f := range files {
		content[f.Name] = f.Content
	}
	for name, wants := range map[string][]string{
		"MyShopFsmEvents.h": {
			"\nclass MyShopFsmEvents {\n",
			"\n    virtual void checkout() = 0;\n    virtual void card() = 0;\n    virtual void retry() = 0;\n    virtual void go() = 0;\n    virtual void tick() = 0;\n",
		},
		"MyShopFsmActions.h": {
			"\n    // Guards\n    virtual bool blocked() = 0;\n    virtual bool card() const = 0;\n    virtual bool event() const = 0;\n",
			"\n    // Actions\n    virtual void charge() = 0;\n    virtual void checkout() = 0;\n    virtual void log() = 0;\n",
		},
		"MyShopFsm.h": {
			"\nclass MyShopFsm : public MyShopFsmEvents {\n",
			"\n    explicit MyShopFsm(MyShopFsmActions& actions);\n",
			"\n    void checkout() override;\n",
		},
		"MyShopFsm.cpp": {
			"\nstruct machine_state {\n",
			"\n            *\"waiting\"_s + sml::event<event::card> [card && !blocked] / (charge, log) = \"paid\"_s,\n",
			"\n            *sml::state<::internal::stopped> + sml::event<::internal::enter> = \"pick\"_s,\n",
			"\n            \"pick\"_s [event] = sml::state<::checkout>,\n",
			"\n            \"pick\"_s = \"say \\\"hi\\\"\"_s,\n",
			"\n            \"say \\\"hi\\\"\"_s + sml::event<::event::checkout> / checkout = sml::state<::checkout>,\n",
			"\n            sml::state<::checkout> + sml::event<::event::go> [blocked] = sml::state<machine_state>,\n",
			"\n            sml::state<::checkout> + sml::event<::internal::stop> = sml::state<::internal::stopped>,\n",
			"\n        const auto blocked = [](MyShopFsmActions& actions) { return actions.blocked(); };\n",
			"\nvoid MyShopFsm::checkout() { machine_->sm.process_event(event::checkout{}); }\n",
		},
	} {
		for _, want := range wants {
			if !strings.Contains(content[name], want) {
				t.Errorf("%s lacks %q:\n%s", name, want, content[name])
			}
		}
	}
	if strings.Contains(content["MyShopFsm.cpp"], "const MyShopFsmActions& actions) { return actions.blocked()") {
		t.Errorf("blocked is also an action, so no lambda may take the actions as const:\n%s", content["MyShopFsm.cpp"])
	}
}

// smlToolchain returns the C++ compiler and the directory holding
// boost/sml.hpp, or skips the calling test. CXX names the compiler, or
// clang++ or g++ on PATH is used; SML_INCLUDE names the directory (see
// `make sml`).
func smlToolchain(t *testing.T) (cxx, include string) {
	t.Helper()
	include = os.Getenv("SML_INCLUDE")
	if include == "" {
		t.Skip("Boost.SML not found; run `make sml` to set up SML_INCLUDE")
	}
	for _, name := range []string{os.Getenv("CXX"), "clang++", "g++"} {
		if name == "" {
			continue
		}
		if path, err := exec.LookPath(name); err == nil {
			return path, include
		}
	}
	t.Skip("no C++ compiler found; set CXX or install clang++ or g++")
	return "", ""
}

var (
	pureVirtual = regexp.MustCompile(`(?m)^    virtual (bool|void|std::chrono::milliseconds) ([A-Za-z0-9_]+)\(\)( const)? = 0;$`)
	className   = regexp.MustCompile(`(?m)^class ([A-Za-z0-9_]+) \{$`)
	constructor = regexp.MustCompile(`(?m)^    explicit ([A-Za-z0-9_]+)\([A-Za-z0-9_]+& actions(, FsmTimers& timers)?((?:, [A-Za-z0-9_]+& [A-Za-z0-9_]+)*)\);$`)
	machineArg  = regexp.MustCompile(`, ([A-Za-z0-9_]+)& ([A-Za-z0-9_]+)`)
)

// smlStubs returns the includes of every state machine among files, timers
// that record which delays start and are cancelled and fire only when told
// to, and, for each actions interface, a struct implementing it: its guards
// return false, its actions append their name to trace, and its delays are a
// second.
func smlStubs(files []File) string {
	var b strings.Builder
	for _, f := range files {
		if strings.HasSuffix(f.Name, "Fsm.h") {
			b.WriteString("#include \"" + f.Name + "\"\n")
		}
	}
	b.WriteString(`#include "FsmTimers.h"

#include <string>
#include <vector>

std::vector<std::string> trace;

struct StubTimers : FsmTimers {
  Id startTimer(std::chrono::milliseconds delay, std::function<void()> fire) override {
    trace.push_back("start " + std::to_string(delay.count()));
    delays.push_back(delay.count());
    fires.push_back(std::move(fire));
    return fires.size() - 1;
  }
  void cancelTimer(Id id) override { trace.push_back("cancel " + std::to_string(delays[id])); }
  // Fires the timer id, as many times as it is called.
  void fire(Id id) { fires[id](); }
  std::vector<long long> delays;
  std::vector<std::function<void()>> fires;
};
`)
	for _, f := range files {
		if !strings.HasSuffix(f.Name, "Actions.h") {
			continue
		}
		class := className.FindStringSubmatch(f.Content)[1]
		b.WriteString("\nstruct Stub" + class + " : " + class + " {\n")
		for _, m := range pureVirtual.FindAllStringSubmatch(f.Content, -1) {
			switch m[1] {
			case "bool":
				b.WriteString("  bool " + m[2] + "()" + m[3] + " override { return false; }\n")
			case "std::chrono::milliseconds":
				b.WriteString("  std::chrono::milliseconds " + m[2] + "() const override { return std::chrono::seconds{1}; }\n")
			default:
				b.WriteString("  void " + m[2] + "() override { trace.push_back(\"" + m[2] + "\"); }\n")
			}
		}
		b.WriteString("};\n")
	}
	return b.String()
}

// smlConstruct returns the statements that build the state machine fsm.h
// declares as the variable fsm, from a stub of its actions, the stub timers
// timers when it has time triggers, and a machine of its own for each
// submachine state.
func smlConstruct(files []File, fsm string) string {
	var header string
	for _, f := range files {
		if f.Name == fsm+".h" {
			header = f.Content
		}
	}
	m := constructor.FindStringSubmatch(header)
	var b strings.Builder
	args := "stub" + fsm + "Actions"
	b.WriteString("  Stub" + fsm + "Actions " + args + ";\n")
	if m[2] != "" {
		b.WriteString("  StubTimers timers;\n")
		args += ", timers"
	}
	for _, arg := range machineArg.FindAllStringSubmatch(m[3], -1) {
		b.WriteString("  Stub" + arg[1] + "Actions " + arg[2] + "Actions;\n")
		b.WriteString("  " + arg[1] + " " + arg[2] + "{" + arg[2] + "Actions};\n")
		args += ", " + arg[2]
	}
	b.WriteString("  " + fsm + " fsm{" + args + "};\n")
	return b.String()
}

// smlMain returns a program that builds the state machine named fsm, enters
// it and sends it every event.
func smlMain(files []File, fsm string) string {
	var b strings.Builder
	b.WriteString(smlStubs(files))
	b.WriteString("\nint main() {\n" + smlConstruct(files, fsm) + "  fsm.enterFsm();\n")
	for _, f := range files {
		if f.Name == fsm+"Events.h" {
			for _, m := range pureVirtual.FindAllStringSubmatch(f.Content, -1) {
				b.WriteString("  fsm." + m[2] + "();\n")
			}
		}
	}
	b.WriteString("}\n")
	return b.String()
}

// smlStandIns renders, for the machines the submachine states of sm run, a
// machine with the entry and exit points they reference, so that sm's files
// compile.
func smlStandIns(t *testing.T, sm *model.StateMachine) []File {
	t.Helper()
	var names []string
	points := map[string][]*model.State{}
	for _, s := range sm.States {
		if s.Submachine == "" {
			continue
		}
		name := s.Submachine
		if _, ok := points[name]; !ok {
			names = append(names, name)
		}
		points[name] = append(points[name], sm.Children(s.Name)...)
	}
	var files []File
	for _, name := range names {
		m := &model.StateMachine{Name: name, Initial: "s", States: []*model.State{{Name: "s", Kind: model.Normal}}}
		for _, p := range points[name] {
			if m.State(p.Name) != nil {
				continue
			}
			m.States = append(m.States, &model.State{Name: p.Name, Kind: p.Kind})
			if p.Kind == model.EntryPoint {
				m.Transitions = append(m.Transitions, &model.Transition{Source: p.Name, Targets: []string{"s"}})
			}
		}
		if err := m.Validate(); err != nil {
			t.Fatal(err)
		}
		f, _ := smlFiles(t, m)
		files = append(files, f...)
	}
	return files
}

// smlFsm returns the name of the state machine class among files.
func smlFsm(files []File) string {
	for _, f := range files {
		if strings.HasSuffix(f.Name, ".cpp") {
			return strings.TrimSuffix(f.Name, ".cpp")
		}
	}
	return ""
}

// smlWrite writes files and the program main into dir and returns the paths
// of the C++ sources among them.
func smlWrite(t *testing.T, dir string, files []File, main string) []string {
	t.Helper()
	sources := []string{filepath.Join(dir, "main.cpp")}
	if err := os.WriteFile(sources[0], []byte(main), 0o644); err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if err := os.WriteFile(filepath.Join(dir, f.Name), []byte(f.Content), 0o644); err != nil {
			t.Fatal(err)
		}
		if strings.HasSuffix(f.Name, ".cpp") {
			sources = append(sources, filepath.Join(dir, f.Name))
		}
	}
	return sources
}

// The goldens only prove the output has not changed. This compiles every
// document's rendering together with a program that implements its actions
// and sends it its events, so that SML itself checks the tables, with a
// stand-in for each machine a submachine state runs.
func TestSMLCompiles(t *testing.T) {
	cxx, include := smlToolchain(t)
	rendered := map[string][]File{}
	for input := range smlGoldens {
		files, _ := smlFilesOf(t, input)
		rendered[filepath.Base(input)] = append(files, smlStandIns(t, parseFile(t, input))...)
	}
	rendered["clashes"], _ = smlFiles(t, clashes)
	for name, files := range rendered {
		t.Run(name, func(t *testing.T) {
			sources := smlWrite(t, t.TempDir(), files, smlMain(files, smlFsm(files)))
			args := append([]string{"-std=c++20", "-fsyntax-only", "-Wall", "-Wextra", "-Werror", "-I", include}, sources...)
			if out, err := exec.Command(cxx, args...).CombinedOutput(); err != nil {
				t.Errorf("the rendered files do not compile: %v\n%s", err, out)
			}
		})
	}
}

// testdata/sml/desk.json runs the machine agent.json in a submachine state, which it enters at the
// machine's initial state or at one of its entry points, and leaves when the
// machine completes, when it leaves by its exit point, or on an event of its
// own. The entry point quick completes agent at once, while desk enters it,
// and broken terminates it at once. agent also terminates from inside the
// composite state person, which ends desk too.
//
// A submachine state runs a machine of its own, generated from another
// document: this links desk and agent and runs them, checking the actions
// they take in turn. Leaving the submachine state early, or stopping desk,
// runs the exit behaviour of agent's active state first. A terminate in agent
// ends both machines without running any exit behaviour, even while desk is
// entering agent, and desk then ignores its events.
func TestSMLSubmachineRuns(t *testing.T) {
	cxx, include := smlToolchain(t)
	var files []File
	for _, name := range []string{"desk", "agent"} {
		sm := parseFile(t, "testdata/sml/"+name+".json")
		f, _ := smlFiles(t, sm)
		files = append(files, f...)
	}
	main := smlStubs(files) + `
#include <cstdio>

struct Listener : DeskFsmListener {
  void onFinished() override { trace.push_back("finished"); }
  void onTerminated() override { trace.push_back("terminated"); }
};

void step(const char* what) {
  std::printf("%s:", what);
  for (const auto& action : trace) {
    std::printf(" %s", action.c_str());
  }
  std::printf("\n");
  trace.clear();
}

int main() {
` + smlConstruct(files, "DeskFsm") + `  Listener listener;
  fsm.setListener(&listener);
  fsm.enterFsm();
  step("enter");
  fsm.ask();
  step("ask");
  help.human();
  step("human");
  help.escalate();
  step("escalate");
  fsm.done();
  step("done");
  fsm.enterFsm();
  fsm.hurry();
  step("hurry");
  help.solved();
  step("solved");
  fsm.rush();
  step("rush");
  fsm.ask();
  fsm.quit();
  step("ask, quit");
  help.human();
  step("human");
  fsm.hurry();
  fsm.quit();
  step("hurry, quit");
  fsm.hurry();
  fsm.stopFsm();
  step("hurry, stopFsm");
  help.solved();
  step("solved");
  fsm.enterFsm();
  fsm.hurry();
  help.hangup();
  step("hurry, hangup");
  fsm.ask();
  help.human();
  step("ask, human");
  fsm.enterFsm();
  fsm.crash();
  step("crash");
  fsm.ask();
  step("ask");
}
`
	dir := t.TempDir()
	sources := smlWrite(t, dir, files, main)
	program := filepath.Join(dir, "program")
	args := append([]string{"-std=c++20", "-Wall", "-Wextra", "-Werror", "-I", include, "-o", program}, sources...)
	if out, err := exec.Command(cxx, args...).CombinedOutput(); err != nil {
		t.Fatalf("the rendered files do not compile: %v\n%s", err, out)
	}
	out, err := exec.Command(program).CombinedOutput()
	if err != nil {
		t.Fatalf("the program failed: %v\n%s", err, out)
	}
	want := `enter:
ask: open greet
human: assign
escalate: release close page
done: finished
hurry: open assign
solved: release close
rush: open close
ask, quit: open greet close
human:
hurry, quit: open assign release close
hurry, stopFsm: open assign release close
solved:
hurry, hangup: open assign terminated
ask, human:
crash: open terminated
ask:
`
	if string(out) != want {
		t.Errorf("the program printed\n%s\nwant\n%s", out, want)
	}
}

// testdata/sml/door.json has an invariant on opened, made of two names.
// An invariant is checked once its state is entered and after every event
// while it is active, each of its names calling invariant<Name> on the
// actions, and the listener hears when it does not hold.
func TestSMLInvariantsRun(t *testing.T) {
	cxx, include := smlToolchain(t)
	sm := parseFile(t, "testdata/sml/door.json")
	files, _ := smlFiles(t, sm)
	main := `#include "DoorFsm.h"

#include <cstdio>

bool clear = false;
bool locked = false;

struct Actions : DoorFsmActions {
  void listen() override { std::printf(" listen"); }
  bool invariantClear() const override { return clear; }
  bool invariantLocked() const override { return locked; }
};

struct Listener : DoorFsmListener {
  void onInvariantViolated(const char* state) override { std::printf(" violated %s", state); }
};

int main() {
  Actions actions;
  Listener listener;
  DoorFsm fsm{actions};
  fsm.setListener(&listener);
  std::printf("enter:");
  fsm.enterFsm();
  std::printf("\nknock while closed:");
  fsm.knock();
  std::printf("\nopen, not clear:");
  fsm.open();
  clear = true;
  std::printf("\nknock, clear:");
  fsm.knock();
  locked = true;
  std::printf("\nknock, locked:");
  fsm.knock();
  std::printf("\nclose:");
  fsm.close();
  std::printf("\nknock while closed:");
  fsm.knock();
  std::printf("\n");
}
`
	dir := t.TempDir()
	sources := smlWrite(t, dir, files, main)
	program := filepath.Join(dir, "program")
	args := append([]string{"-std=c++20", "-Wall", "-Wextra", "-Werror", "-I", include, "-o", program}, sources...)
	if out, err := exec.Command(cxx, args...).CombinedOutput(); err != nil {
		t.Fatalf("the rendered files do not compile: %v\n%s", err, out)
	}
	out, err := exec.Command(program).CombinedOutput()
	if err != nil {
		t.Fatalf("the program failed: %v\n%s", err, out)
	}
	want := `enter:
knock while closed:
open, not clear: violated opened
knock, clear: listen
knock, locked: listen violated opened
close:
knock while closed:
`
	if string(out) != want {
		t.Errorf("the program printed\n%s\nwant\n%s", out, want)
	}
}

// testdata/sml/oven.json has timers on a composite state, one with a named
// delay whose transition is internal, one on a nested state, and a terminate.
// A state starts its timers when it is entered and cancels them when it is
// left, a timer that fired is not cancelled, and one that fires after its
// state was left, or after the machine was terminated or stopped, does
// nothing.
func TestSMLTimersRun(t *testing.T) {
	cxx, include := smlToolchain(t)
	sm := parseFile(t, "testdata/sml/oven.json")
	files, _ := smlFiles(t, sm)
	main := smlStubs(files) + `
#include <cstdio>

void step(const char* what) {
  std::printf("%s:", what);
  for (const auto& action : trace) {
    std::printf(" %s", action.c_str());
  }
  std::printf("\n");
  trace.clear();
}

int main() {
` + smlConstruct(files, "OvenFsm") + `  fsm.enterFsm();
  step("enter");
  fsm.start();
  step("start");
  fsm.hot();
  step("hot");
  timers.fire(2);
  step("fire 250ms");
  timers.fire(2);
  step("fire 250ms again");
  timers.fire(1);
  step("fire preheat");
  fsm.hot();
  fsm.open();
  step("hot, open");
  timers.fire(0);
  step("fire the 30s of the first visit");
  fsm.start();
  fsm.burn();
  step("start, burn");
  timers.fire(4);
  step("fire the 30s before burn");
  fsm.enterFsm();
  fsm.start();
  fsm.stopFsm();
  step("enter, start, stopFsm");
  fsm.enterFsm();
  fsm.start();
  timers.fire(8);
  step("enter, start, fire 30s");
}
`
	dir := t.TempDir()
	sources := smlWrite(t, dir, files, main)
	program := filepath.Join(dir, "program")
	args := append([]string{"-std=c++20", "-Wall", "-Wextra", "-Werror", "-I", include, "-o", program}, sources...)
	if out, err := exec.Command(cxx, args...).CombinedOutput(); err != nil {
		t.Fatalf("the rendered files do not compile: %v\n%s", err, out)
	}
	out, err := exec.Command(program).CombinedOutput()
	if err != nil {
		t.Fatalf("the program failed: %v\n%s", err, out)
	}
	want := `enter:
start: start 30000 start 1000
hot: start 250
fire 250ms: beep
fire 250ms again:
fire preheat: ready
hot, open: start 250 cancel 250 cancel 30000
fire the 30s of the first visit:
start, burn: start 30000 start 1000 cancel 30000 cancel 1000
fire the 30s before burn:
enter, start, stopFsm: start 30000 start 1000 cancel 30000 cancel 1000
enter, start, fire 30s: start 30000 start 1000 cancel 1000 ding
`
	if string(out) != want {
		t.Errorf("the program printed\n%s\nwant\n%s", out, want)
	}
}

// A delay becomes a std::chrono duration, in seconds when it is whole ones. A
// named delay is read from the actions. A timer on a region that holds states
// starts with its orthogonal state, and its transitions become the orthogonal
// state's. What has to be rounded is said.
func TestSMLDelays(t *testing.T) {
	sm := &model.StateMachine{
		Name: "d", Initial: "a",
		States: []*model.State{
			{Name: "a", Kind: model.Normal},
			{Name: "p", Kind: model.Parallel},
			{Name: "r1", Parent: "p", Kind: model.Normal, Initial: "b"},
			{Name: "b", Parent: "r1", Kind: model.Normal},
			{Name: "r2", Parent: "p", Kind: model.Normal},
		},
		Transitions: []*model.Transition{
			{Source: "a", Targets: []string{"p"}, After: "1.5s"},
			{Source: "a", Targets: []string{"p"}, After: "2.0s"},
			{Source: "a", Targets: []string{"p"}, After: "250ms"},
			{Source: "a", Targets: []string{"p"}, After: "1.2345s"},
			{Source: "a", Targets: []string{"p"}, After: "ok", Cond: "ready"},
			{Source: "r1", Targets: []string{"a"}, After: "3s"},
		},
	}
	if err := sm.Validate(); err != nil {
		t.Fatal(err)
	}
	files, warnings := smlFiles(t, sm)
	wantWarnings := []string{
		`transition a -> p: the delay 1.2345s is rounded down to whole milliseconds`,
		`transition r1 -> a: Boost.SML has no transitions from a region; written as a transition of "p", which is active exactly when "r1" is`,
	}
	if !reflect.DeepEqual([]string(warnings), wantWarnings) {
		t.Errorf("warnings =\n%s\nwant\n%s", strings.Join(warnings, "\n"), strings.Join(wantWarnings, "\n"))
	}
	content := map[string]string{}
	for _, f := range files {
		content[f.Name] = f.Content
	}
	for name, wants := range map[string][]string{
		"DFsmActions.h": {"\n    virtual std::chrono::milliseconds ok() const = 0;\n"},
		"DFsm.cpp": {
			"t.start(::timers::a_after_1_5s, std::chrono::milliseconds{1500}); }",
			"t.start(::timers::a_after_2_0s, std::chrono::seconds{2}); }",
			"t.start(::timers::a_after_250ms, std::chrono::milliseconds{250}); }",
			"t.start(::timers::a_after_1_2345s, std::chrono::milliseconds{1234}); }",
			"[](::timers& t, const DFsmActions& actions) { t.start(::timers::a_after_ok, actions.ok()); }",
			"\n            sml::state<p> + sml::on_entry<sml::_> / [](::timers& t) { t.start(::timers::r1_after_3s, std::chrono::seconds{3}); },\n",
		},
	} {
		for _, want := range wants {
			if !strings.Contains(content[name], want) {
				t.Errorf("%s lacks %q:\n%s", name, want, content[name])
			}
		}
	}
}
