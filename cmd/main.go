package main

import (
	"fmt"
	"log"

	"flag"

	"go/token"

	"golang.org/x/tools/go/loader"
)

func main() {
	target := flag.String("target", "../example/sampleproject", "Parse Target")
	flag.Parse()

	//var conf loader.Config
	conf := loader.Config{
		Fset: token.NewFileSet(),
	}
	conf.CreateFromFilenames(*target)
	prog, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(prog.AllPackages)
	fmt.Println(prog.Fset)
	//fmt.Println(prog)

	//fset := token.NewFileSet()
	//filter := func(os.FileInfo) bool {
	//	return true
	//}
	//pkgs, err := parser.ParseDir(fset, *target, filter, )
	//if err != nil {
	//	panic(err)
	//}

	//f := &ast.File{
	//	Name:  ast.NewIdent(*target),
	//	Decls: []ast.Decl{},
	//}
	//format.Node(os.Stdout, token.NewFileSet(), f)
}
