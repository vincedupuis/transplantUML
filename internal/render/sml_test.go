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
	"github.com/vincedupuis/transplantUML/internal/fsm"
	"github.com/vincedupuis/transplantUML/internal/model"
)

var fsmParser fsm.Parser

// smlGoldens maps the documents whose Boost.SML rendering is checked in to the
// folder holding its files. Regenerate one with:
// go run ./cmd/fsm -i <input> -t sml -o internal/render/testdata/sml/<name>
var smlGoldens = map[string]string{
	"../../example/kiosk.fsm":     "testdata/sml/kiosk",
	"../../example/shop.fsm":      "testdata/sml/shop",
	"../../example/support.fsm":   "testdata/sml/support",
	"../scxml/testdata/uml.scxml": "testdata/sml/uml",
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
	files, _ := smlFiles(t, &model.StateMachine{})
	want := []string{"MachineFsmEvents.h", "MachineFsmActions.h", "MachineFsm.h", "MachineFsm.cpp"}
	if got := fileNames(files); !reflect.DeepEqual(got, want) {
		t.Errorf("files = %v, want %v", got, want)
	}
	files, _ = smlFiles(t, &model.StateMachine{Name: "my-coffee machine"})
	if got := files[0].Name; got != "MyCoffeeMachineFsmEvents.h" {
		t.Errorf("first file = %q", got)
	}
}

// The SML template must say what it cannot write.
func TestSMLWarnings(t *testing.T) {
	_, warnings := smlFilesOf(t, "../scxml/testdata/uml.scxml")
	want := []string{
		`the initial transition of "outer": the action "log('hi')" is not a name; it is not written`,
		`state "in": Boost.SML has no entry points; transitions to it enter "outer" at its initial state, and its own transitions are not written`,
		`state "out": Boost.SML has no exit points; the transitions to and from it are not written`,
		`transition inner -> inner: Boost.SML has no local transitions; written as an external one`,
		`transition check -> failed: the guard [retries > 3] is not made of names; the transition is written as a comment`,
		`transition check -> work: the guard [retries <= 3] is not made of names; the transition is written as a comment`,
		`state "split": Boost.SML has no fork; transitions to it enter "both", whose regions start in their initial states, and its own transitions are not written`,
		`state "sync": Boost.SML has no join; transitions to it end their region (X), and its outgoing transition is taken once every region of "both" has ended`,
		`the entry behaviour of "work": the action "log('start')" is not a name; it is not written`,
		`state "work": Boost.SML has no do activities; invoke(job.py, http://example.com/worker) is written as a comment`,
		`transition work -> check: Boost.SML has no time triggers; written on the event after_5s, which the application raises 5s after entering work`,
		`transition work -> check: the action "retries = retries + 1" is not a name; it is not written`,
		`transition work -> check: Boost.SML has no time triggers; written on the event after_retryDelay, which the application raises retryDelay after entering work`,
		`the exit behaviour of "outer": the action "log('bye')" is not a name; it is not written`,
		`state "stop": Boost.SML has no terminate; written as X, which ends only its own region`,
	}
	if !reflect.DeepEqual([]string(warnings), want) {
		t.Errorf("warnings =\n%s\nwant\n%s", strings.Join(warnings, "\n"), strings.Join(want, "\n"))
	}
}

// The machine's own points are written, save what cannot reach or leave them:
// a transition to its entry point, which only enterFsm enters, and one from
// its exit point, where the machine ends. An event named like a method that
// starts or stops the machine makes the files fail to compile.
func TestSMLMachinePointWarnings(t *testing.T) {
	sm := &model.StateMachine{
		Name: "m", Initial: "a",
		States: []*model.State{
			{Name: "a", Kind: model.Normal},
			{Name: "in", Kind: model.EntryPoint},
			{Name: "out", Kind: model.ExitPoint},
		},
		Transitions: []*model.Transition{
			{Source: "in", Targets: []string{"a"}},
			{Source: "a", Targets: []string{"in"}, Event: "enterFsm"},
			{Source: "a", Targets: []string{"out"}, Event: "leave"},
			{Source: "out", Targets: []string{"a"}},
		},
	}
	if err := sm.Validate(); err != nil {
		t.Fatal(err)
	}
	_, warnings := smlFiles(t, sm)
	want := []string{
		`event "enterFsm": its method enterFsm clashes with the one that starts or stops the machine; the files do not compile`,
		`state "in": the machine is entered at its entry point by enterFsm; the transitions to it are not written`,
		`state "out": the machine ends at its exit point; the transitions from it are not written`,
	}
	if !reflect.DeepEqual([]string(warnings), want) {
		t.Errorf("warnings =\n%s\nwant\n%s", strings.Join(warnings, "\n"), strings.Join(want, "\n"))
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
			"\n            sml::state<::checkout> + sml::event<::event::go> [blocked] = sml::state<machine_state>\n",
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
	pureVirtual = regexp.MustCompile(`(?m)^    virtual (bool|void) ([A-Za-z0-9_]+)\(\)( const)? = 0;$`)
	className   = regexp.MustCompile(`(?m)^class ([A-Za-z0-9_]+) \{$`)
	constructor = regexp.MustCompile(`(?m)^    explicit ([A-Za-z0-9_]+)\([A-Za-z0-9_]+& actions((?:, [A-Za-z0-9_]+& [A-Za-z0-9_]+)*)\);$`)
	machineArg  = regexp.MustCompile(`, ([A-Za-z0-9_]+)& ([A-Za-z0-9_]+)`)
)

// smlStubs returns the includes of every state machine among files and, for
// each actions interface, a struct implementing it: its guards return false,
// and its actions append their name to trace.
func smlStubs(files []File) string {
	var b strings.Builder
	for _, f := range files {
		if strings.HasSuffix(f.Name, "Fsm.h") {
			b.WriteString("#include \"" + f.Name + "\"\n")
		}
	}
	b.WriteString("\n#include <string>\n#include <vector>\n\nstd::vector<std::string> trace;\n")
	for _, f := range files {
		if !strings.HasSuffix(f.Name, "Actions.h") {
			continue
		}
		class := className.FindStringSubmatch(f.Content)[1]
		b.WriteString("\nstruct Stub" + class + " : " + class + " {\n")
		for _, m := range pureVirtual.FindAllStringSubmatch(f.Content, -1) {
			if m[1] == "bool" {
				b.WriteString("  bool " + m[2] + "()" + m[3] + " override { return false; }\n")
			} else {
				b.WriteString("  void " + m[2] + "() override { trace.push_back(\"" + m[2] + "\"); }\n")
			}
		}
		b.WriteString("};\n")
	}
	return b.String()
}

// smlConstruct returns the statements that build the state machine fsm.h
// declares as the variable fsm, from a stub of its actions and a machine of
// its own for each submachine state.
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
	for _, arg := range machineArg.FindAllStringSubmatch(m[2], -1) {
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
		name := strings.TrimSuffix(filepath.Base(s.Submachine), filepath.Ext(s.Submachine))
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
// stand-in for each machine a submachine state runs. Warnings count as
// failures.
func TestSMLCompiles(t *testing.T) {
	cxx, include := smlToolchain(t)
	rendered := map[string][]File{}
	for input := range goldens(t) {
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

// desk runs the machine agent in a submachine state, which it enters at the
// machine's initial state or at one of its entry points, and leaves when the
// machine completes, when it leaves by its exit point, or on an event of its
// own. The entry point quick completes agent at once, while desk enters it.
const desk = `fsm desk {
  initial state idle {
    on ask goto help
    on hurry goto fast
    on rush goto quick
  }
  submachine help : agent {
    entry / open
    exit / close
    entry point fast
    entry point quick
    exit point up / page goto top
    on quit goto idle
    goto idle
  }
  state top { on done goto final }
}`

const agent = `fsm agent {
  entry point fast goto person
  entry point quick goto final
  exit point up
  initial state bot {
    entry / greet
    on human goto person
    on solved goto final
  }
  state person {
    entry / assign
    exit / release
    on solved goto final
    on escalate goto up
  }
}`

// A submachine state runs a machine of its own, generated from another
// document: this links desk and agent and runs them, checking the actions
// they take in turn.
func TestSMLSubmachineRuns(t *testing.T) {
	cxx, include := smlToolchain(t)
	var files []File
	for _, src := range []string{desk, agent} {
		sm, _, err := fsmParser.Parse([]byte(src))
		if err != nil {
			t.Fatal(err)
		}
		f, _ := smlFiles(t, sm)
		files = append(files, f...)
	}
	main := smlStubs(files) + `
#include <cstdio>

struct Listener : DeskFsmListener {
  void onFinished() override { trace.push_back("finished"); }
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
`
	if string(out) != want {
		t.Errorf("the program printed\n%s\nwant\n%s", out, want)
	}
}
