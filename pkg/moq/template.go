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
	{{.SmallName}}Calls int
{{- end }}
}
{{ range .Methods }}

func (mock *Mock{{$obj.InterfaceName}}) SetMock{{.Name}}(mockFunc func({{ .Arglist }}) {{.ReturnArglist}}) (*Mock{{$obj.InterfaceName}}) {
	mock.Mock{{.Name}} = mockFunc
	return mock
}

func (mock *Mock{{$obj.InterfaceName}}) {{.Name}}({{.Arglist}}) {{.ReturnArglist}} {
	mock.{{.SmallName}}Calls++
	return mock.Mock{{.Name}}({{ .ArgCallList }})
}

func (mock *Mock{{$obj.InterfaceName}}) {{.Name}}Calls() int {
	return mock.{{.SmallName}}Calls
}
{{- end}}

func (mock *Mock{{$obj.InterfaceName}}) SetOpts (opts {{.InterfaceName}}Opts) {
	{{- range .Methods }}
	mock.Mock{{.Name}} = func({{ .Arglist }}) {{.ReturnArglist}} {
		if opts.Is{{.Name}}Error {
			return {{.ReturnValuelist false}}
		}
		return {{.ReturnValuelist true}}
	}
	{{- end }}
}

func NewMock{{.InterfaceName}}(opts {{.InterfaceName}}Opts) *Mock{{.InterfaceName}} {
	mock := new(Mock{{.InterfaceName}})
	mock.SetOpts(opts)
	return mock
}
{{ end -}}`
