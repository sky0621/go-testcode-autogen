package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func main() {
	target := flag.String("target", "../example/sampleproject", "Parse Target")
	flag.Parse()

	filepath.Walk(*target, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if info.IsDir() {
			return nil
		}

		// FIXME vendor配下を見ないようフィルタリング
		// FIXME テストコードを見ないようフィルタリング

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Name: %#v\n", f.Name)
		decls := f.Decls
		for _, decl := range decls {
			fmt.Printf("decl: %#v\n", decl)
		}
		fmt.Println("======================================")
		return nil
	})

}
