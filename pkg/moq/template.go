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
	{{.SmallName}}Chan chan int
{{- end }}
}
{{ range .Methods }}

func (mock *Mock{{$obj.InterfaceName}}) SetMock{{.Name}}(mockFunc func({{ .Arglist }}) {{.ReturnArglist}}) (*Mock{{$obj.InterfaceName}}) {
	mock.Mock{{.Name}} = mockFunc
	return mock
}

func (mock *Mock{{$obj.InterfaceName}}) {{.Name}}({{.Arglist}}) {{.ReturnArglist}} {
	mock.{{.SmallName}}Calls++{{if .HasReturnArgs}}
	select {
    case mock.{{.SmallName}}Chan <- 1:
    default:
    }
	return mock.Mock{{.Name}}({{ .ArgCallList }}){{end}}
}

func (mock *Mock{{$obj.InterfaceName}}) {{.Name}}Calls() int {
	return mock.{{.SmallName}}Calls
}

func (mock *Mock{{$obj.InterfaceName}}) {{.Name}}Chan() chan int {
	return mock.{{.SmallName}}Chan
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
	mock.{{.SmallName}}Chan = make(chan int)
	{{- end }}
}

func NewMock{{.InterfaceName}}(opts {{.InterfaceName}}Opts) *Mock{{.InterfaceName}} {
	mock := new(Mock{{.InterfaceName}})
	mock.SetOpts(opts)
	return mock
}
{{ end -}}`
