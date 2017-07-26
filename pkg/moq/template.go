package moq

// moqImports are the imports all moq files get.
var moqImports = []string{"sync"}

// moqTemplate is the template for mocked code.
var moqTemplate = `package {{.PackageName}}

// Mocks automatically generated with Moq :)
{{ range $i, $obj := .Objects -}}
const {{.InterfaceName}}ErrorMessage = "{{.InterfaceName}}-error"

type {{.InterfaceName}}Opts struct {
	{{- range .Methods }}
	Is{{.Name}}Error bool
{{- end }}
}

type {{.InterfaceName}}Mock struct {
{{- range .Methods }}
	Mock{{.Name}} func({{ .Arglist }}) {{.ReturnArglist}}
{{- end }}
}
{{ range .Methods }}
func (mock *{{$obj.InterfaceName}}Mock) {{.Name}}({{.Arglist}}) {{.ReturnArglist}} {
	return mock.Mock{{.Name}}({{ .ArgCallList }})
}
{{ end -}}

func New{{.InterfaceName}}Mock(opts {{.InterfaceName}}Opts) *{{.InterfaceName}} {
	mock := new({{.InterfaceName}}Mock)

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
