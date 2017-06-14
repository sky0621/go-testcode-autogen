package {{.PkgName}}

import (
	"testing"
{{range .Imports}}
    {{.ImportPkg}}
{{end}}
)
{{range .Functions}}
func Test{{.FuncName}}(t *testing.T) {
    // FIXME write the necessary test code 
}
{{end}}
