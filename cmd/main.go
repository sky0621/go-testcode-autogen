package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	target := flag.String("target", "../example/sampleproject", "Parse Target")
	flag.Parse()

	// TODO config読み込んでロジックに反映

	// TODO go/astパッケージのWalkを使うようにする
	filepath.Walk(*target, Apply)
}

func Apply(path string, info os.FileInfo, err error) error {
	if err != nil {
		panic(err)
	}

	if !IsTarget(path, info) {
		return nil
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	fmt.Println("======================================")
	out := &TestCodeInfo{}

	fncs := []*Function{}
	for _, decl := range f.Decls {
		//ast.Print(fset, decl)
		switch declType := decl.(type) {
		case *ast.GenDecl:
			//for _, spec := range declType.Specs {
			//	fmt.Printf("[GenDecl]: %#v\n", spec)
			//}
			switch declType.Tok {
			//case token.PACKAGE:
			//fmt.Printf("PACKAGE: %#v\n", declType)
			//	case token.IMPORT:
			//	case token.TYPE:
			//	case token.CONST:
			//	case token.VAR:
			default:
			}
		case *ast.FuncDecl:
			if IsFirstUpper(declType.Name.String()) {
				fncs = append(fncs, &Function{FuncName: declType.Name.String()})
			}
		default:
		}
	}

	if len(fncs) < 1 {
		return nil
	}

	out.Functions = fncs
	//os.OpenFile()

	tmpl := template.Must(template.ParseFiles("./testcode.md"))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, out)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())

	fmt.Println("======================================")
	return nil
}

func IsTarget(path string, info os.FileInfo) bool {
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
	PkgName   string
	Functions []*Function
}

type Function struct {
	FuncName string
}

func IsFirstUpper(v string) bool {
	if v == "" {
		return false
	}
	r := rune(v[0])
	return unicode.IsUpper(r)
}
