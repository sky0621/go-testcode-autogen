package "{{.PkgName}}"

import (
	"testing"
)
{{range .Functions}}
func Test{{.FuncName}}(t *testing.T) {
    // FIXME write the necessary test code 
}
{{end}}
