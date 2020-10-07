package stringenum

import (
	"bytes"
	"fmt"
	gofmt "go/format"
)

var generatedSrcTmpl = defineTemplate(`
// Code generated by stringenum - DO NOT EDIT.
package {{.PackageName}}

import "fmt"

{{range .Enums}}

// {{.Type}}Values contains all possible values of {{.Type}}.
var {{.Type}}Values = []{{.Type}}{
	{{range $name, $value := .Values -}}
	{{ $name }},
	{{end}}
}

// {{.Type}}FromValue returns a {{.Type}} for given value.
func {{.Type}}FromValue(s string) (v {{.Type}}, err error) {
	v = ({{.Type}})(s)
	if !v.IsValid() {
		err = fmt.Errorf("%s is not a valid {{.Type}}", s)
		return
	}
	return v, nil
}

// Is{{.Type}} returns true if given value is a valid {{.Type}}.
func (v {{.Type}}) IsValid() bool {
	for _, val := range {{.Type}}Values {
		if val == v {
			return true
		}
	}
	return false
}

// {{.Type}} returns an error if the value is not valid.
func (v {{.Type}}) Validate() error {
	if _, err := {{.Type}}FromValue(string(v)); err != nil {
		{{ $possibleValues := StringJoin .EnumValues ", " -}}
		return fmt.Errorf("%w. possible values are: {{ js $possibleValues }}")
	}
	return nil
}

// String returns a string value of {{.Type}}.
func (v {{.Type}}) String() string {
	return string(v)
}

// MarshalText implements encoding.TextMarshaler interface which is compatible with JSON, YAML.
func (v {{.Type}}) MarshalText() ([]byte, error) {
	return []byte(v), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface which is compatible with JSON, YAML.
func (v *{{.Type}}) UnmarshalText(d []byte) error {
	vv, err := {{.Type}}FromValue(string(d))
	if err != nil {
		return err
	}
	*v = vv
	return nil
}

{{end}}
`)

func GenerateCode(p ParsedFile) (string, error) {
	buf := bytes.NewBufferString("")
	if err := generatedSrcTmpl.Execute(buf, p); err != nil {
		return "", err
	}
	code, err := gofmt.Source(buf.Bytes())
	if err != nil {
		return "", fmt.Errorf("go fmt: %v. source code was:\n%s", err, DumpWithLine(buf.String()))
	}
	return string(code), nil
}
