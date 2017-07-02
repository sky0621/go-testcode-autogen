package result

import "github.com/sky0621/go-testcode-autogen/template"

type Aggregater struct {
	PackageName string
	ImportNames []string
	Functions   []*Function
}

type Function struct {
	FunctionName    string
	FunctionReturns map[interface{}]string // key:type, value:name
	FunctionParams  map[interface{}]string // key:type, value:name
}

func (a *Aggregater) Convert2TestClassInfo() *template.TestClass {
	// FIXME 適宜変換ロジック実装
	tc := &template.TestClass{
		PackageName: a.PackageName,
		ImportNames: a.ImportNames,
		Functions:   convert2TemplateFunctions(a.Functions),
	}
	return tc
}

func convert2TemplateFunctions(functions []*Function) []*template.Function {
	rf := []*template.Function{}
	for _, f := range functions {
		rf = append(rf, &template.Function{
			FunctionName:    f.FunctionName,
			FunctionReturns: f.FunctionReturns,
			FunctionParams:  f.FunctionParams,
		})
	}
	return rf
}
