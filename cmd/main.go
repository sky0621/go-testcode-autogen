package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	target := flag.String("target", "../example/sampleproject", "Parse Target")
	flag.Parse()

	filepath.Walk(*target, apply)
}

func apply(path string, info os.FileInfo, err error) error {
	if err != nil {
		panic(err)
	}

	if !isTarget(path, info) {
		return nil
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	fmt.Println("======================================")
	//for _, imp := range f.Imports {
	//	fmt.Println("##########################")
	//	fmt.Printf("Imports.Name: %#v\n", imp.Name)
	//	fmt.Printf("Imports.Path.Value: %#v\n", imp.Path.Value)
	//	fmt.Println("##########################")
	//}
	//fmt.Println("-----------------------")
	for _, decl := range f.Decls {
		fmt.Printf("decl: %#v\n", decl)
	}

	//os.OpenFile()

	tmpl := template.Must(template.ParseFiles("./testcode.md"))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, &TestCodeInfo{})
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())

	fmt.Println("======================================")
	return nil
}

func isTarget(path string, info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}
	if filepath.Ext(path) != ".go" {
		return false
	}
	if strings.Contains(path, "_test") {
		return false
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return false
	}
	if strings.Contains(absPath, "vendor") {
		return false
	}

	return true
}

type TestCodeInfo struct {
	Pkg       string
	Imports   []string
	Functions []*Function
}

type Function struct {
	FuncName string
}
