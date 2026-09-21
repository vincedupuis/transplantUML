// Package assets embeds the templates shipped with the tool.
package assets

import _ "embed"

// PlantUML is the built-in template that renders the model as a PlantUML
// state diagram. It is used when no -t template is given.
//
//go:embed puml.tmpl
var PlantUML string
