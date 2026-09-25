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
// folder holding its files. Regenerate one with:
// go run ./cmd/fsm -i <input> -t sml -o internal/render/testdata/sml/<name>
var smlGoldens = map[string]string{
	"../../example/kiosk.fsm":     "testdata/sml/kiosk",
	"../../example/shop.fsm":      "testdata/sml/shop",
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
		`state "sub": Boost.SML cannot run the machine "child.scxml", whose tables are private to its own files; written as a simple state`,
		`the exit behaviour of "outer": the action "log('bye')" is not a name; it is not written`,
		`state "stop": Boost.SML has no terminate; written as X, which ends only its own region`,
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
			"\n            (*\"pick\"_s) [event] = sml::state<::checkout>,\n",
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
)

// smlMain returns a program that implements the actions interface with stubs,
// builds the state machine from it and sends it every event.
func smlMain(files []File) string {
	var fsm, actions, events string
	for _, f := range files {
		switch {
		case strings.HasSuffix(f.Name, "Actions.h"):
			actions = f.Content
		case strings.HasSuffix(f.Name, "Events.h"):
			events = f.Content
		case strings.HasSuffix(f.Name, ".h"):
			fsm = strings.TrimSuffix(f.Name, ".h")
		}
	}
	var b strings.Builder
	b.WriteString("#include \"" + fsm + ".h\"\n\n")
	b.WriteString("struct Stub : " + className.FindStringSubmatch(actions)[1] + " {\n")
	for _, m := range pureVirtual.FindAllStringSubmatch(actions, -1) {
		if m[1] == "bool" {
			b.WriteString("  bool " + m[2] + "()" + m[3] + " override { return false; }\n")
		} else {
			b.WriteString("  void " + m[2] + "() override {}\n")
		}
	}
	b.WriteString("};\n\nint main() {\n  Stub stub;\n  " + fsm + " fsm{stub};\n")
	for _, m := range pureVirtual.FindAllStringSubmatch(events, -1) {
		b.WriteString("  fsm." + m[2] + "();\n")
	}
	b.WriteString("}\n")
	return b.String()
}

// The goldens only prove the output has not changed. This compiles every
// document's rendering together with a program that implements its actions
// and sends it its events, so that SML itself checks the tables. Warnings count
// as failures.
func TestSMLCompiles(t *testing.T) {
	cxx, include := smlToolchain(t)
	rendered := map[string][]File{}
	for input := range goldens(t) {
		files, _ := smlFilesOf(t, input)
		rendered[filepath.Base(input)] = files
	}
	rendered["clashes"], _ = smlFiles(t, clashes)
	for name, files := range rendered {
		t.Run(name, func(t *testing.T) {
			dir := t.TempDir()
			sources := []string{filepath.Join(dir, "main.cpp")}
			if err := os.WriteFile(sources[0], []byte(smlMain(files)), 0o644); err != nil {
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
			args := append([]string{"-std=c++20", "-fsyntax-only", "-Wall", "-Wextra", "-Werror", "-I", include}, sources...)
			if out, err := exec.Command(cxx, args...).CombinedOutput(); err != nil {
				t.Errorf("the rendered files do not compile: %v\n%s", err, out)
			}
		})
	}
}
