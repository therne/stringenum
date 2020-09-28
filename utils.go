package stringenum

import (
	"fmt"
	"strings"
	"text/template"
)

func defineTemplate(code string) *template.Template {
	code = strings.TrimLeft(code, "\n ")
	tmpl, err := template.New("").Parse(code)
	if err != nil {
		panic(err)
	}
	return tmpl
}

func DumpWithLine(code string) string {
	srcWithLine := ""
	for i, line := range strings.Split(code, "\n") {
		line = strings.ReplaceAll(line, "\t", "    ")
		srcWithLine += fmt.Sprintf(" %3d | %s\n", i+1, line)
	}
	return srcWithLine
}
