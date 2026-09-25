// Package assets embeds the templates shipped with the tool.
package assets

import (
	_ "embed"
	"sort"
)

// PlantUML is the built-in template that renders the model as a PlantUML
// state diagram. It is used when no -t template is given.
//
//go:embed puml.gotmpl
var PlantUML string

// SML is the built-in template that renders the model as a C++ header for
// Boost.SML (https://github.com/boost-ext/sml).
//
//go:embed sml.gotmpl
var SML string

// Templates maps the name -t accepts in place of a file to each built-in
// template.
var Templates = map[string]string{
	"puml": PlantUML,
	"sml":  SML,
}

// TemplateNames lists the built-in template names, sorted.
func TemplateNames() []string {
	names := make([]string, 0, len(Templates))
	for name := range Templates {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
