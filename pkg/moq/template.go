package moq

// moqImports are the imports all moq files get.
var moqImports = []string{"sync"}

// moqTemplate is the template for mocked code.
var moqTemplate = `package {{.PackageName}}

{{ range $i, $obj := .Objects -}}
const {{.InterfaceName}}ErrorMessage = "{{.InterfaceName}}-error"

type {{.InterfaceName}}Opts struct {
	{{- range .Methods }}
	Is{{.Name}}Error bool
{{- end }}
}

type Mock{{.InterfaceName}} struct {
{{- range .Methods }}
	Mock{{.Name}} func({{ .Arglist }}) {{.ReturnArglist}}
{{- end }}
}
{{ range .Methods }}
func (mock *Mock{{$obj.InterfaceName}}) {{.Name}}({{.Arglist}}) {{.ReturnArglist}} {
	return mock.Mock{{.Name}}({{ .ArgCallList }})
}
{{ end -}}

func NewMock{{.InterfaceName}}(opts {{.InterfaceName}}Opts) *Mock{{.InterfaceName}} {
	mock := new(Mock{{.InterfaceName}})

	{{- range .Methods }}
		mock.Mock{{.Name}} = func({{ .Arglist }}) {{.ReturnArglist}} {
			if opts.Is{{.Name}}Error {
				return
			}
			return
		}
	{{- end }}

	return mock
}
{{ end -}}`
