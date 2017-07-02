package result

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
