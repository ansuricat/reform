package main

import (
	"text/template"

	"github.com/ansuricat/reform/parse"
)

type StructData struct {
	Imports map[string]struct{}
	parse.StructInfo
	FieldComments []string
}

var (
	prologTemplate = template.Must(template.New("prolog").Parse(`
import (
	{{- range $i, $_ := .Imports }}
	"{{ $i }}"
	{{- end }}
)
`))

	structTemplate = template.Must(template.New("struct").Parse(`
//go:generate reform

// {{ .Type }} represents a row in {{ .SQLName }} table.
//reform:{{ .SQLName }}
type {{ .Type }} struct {
	{{- range $i, $f := .Fields }}
    {{ $f.Name }} {{ $f.Type }} ` + "`" + `reform:"{{ $f.Column }}{{ if eq $i $.PKFieldIndex }},pk{{ end }}" json:"{{ $f.Column }}"` + "`" + ` {{ index $.FieldComments $i }}
	{{- end }}
}
`))
)
