package testinfo

type TestInfo struct {
	PackageName string
	ImportNames []string
	Functions   []*Function
}

type Function struct {
	FunctionName string
}
