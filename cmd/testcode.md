package {{.Pkg}}

import (
	"testing"{{range .Imports}}
	{{.}}{{end}}
)

{{range .Functions}}func Test{{.FuncName}}(t *testing.T) {
    // FIXME write the necessary test code 
}{{end}}
