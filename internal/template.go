package internal

import (
	"bytes"
	"github.com/Masterminds/sprig/v3"
	"strings"
	"text/template"
)

func RunTemplate(sm *StateMachine, tmplFile string) (string, error) {
	tmpl := template.New("WriteStateMachine")

	funcMap := template.FuncMap{
		"include": func(name string, data interface{}) (string, error) {
			buf := bytes.NewBuffer(nil)
			if err := tmpl.ExecuteTemplate(buf, name, data); err != nil {
				return "", err
			}
			return buf.String(), nil
		},
		"prefix": func(prefix, value string) string {
			if value == "" {
				return ""
			}
			return prefix + value
		},
		"surround": func(prefix, value, suffix string) string {
			if value == "" {
				return ""
			}
			return prefix + value + suffix
		},
		"toSlice": func(args ...string) []string {
			return args
		},
		"join": func(values []string, sep string) string {
			nonEmptyValues := make([]string, 0)
			for _, value := range values {
				if value != "" {
					nonEmptyValues = append(nonEmptyValues, value)
				}
			}
			return strings.Join(nonEmptyValues, sep)
		},
		"GetInnerStates":         sm.GetInnerStates,
		"GetIncomingTransitions": sm.GetIncomingTransitions,
		"GetOutgoingTransitions": sm.GetOutgoingTransitions,
	}

	tmpl.Funcs(sprig.FuncMap())
	tmpl.Funcs(funcMap)

	_, err := tmpl.Parse(tmplFile)
	if err != nil {
		return "", err
	}

	output := bytes.NewBuffer(nil)
	err = tmpl.Execute(output, sm)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
