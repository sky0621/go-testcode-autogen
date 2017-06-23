package {{.PackageName}}

import (
	"testing"{{range .ImportNames}}
    "{{.}}"{{end}}
)
{{range .Functions}}
func Test{{.FunctionName}}(t *testing.T) {
    // FIXME write the necessary test code 
	t.Fail()
}
{{end}}