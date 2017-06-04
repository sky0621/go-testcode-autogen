package main

import (
	"bufio"
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

	// TODO go/astパッケージのWalkを使うようにする(ただし、パッケージ名が取れない問題とセットで考慮)
	filepath.Walk(*target, Apply)
}

func Apply(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("【01】")
		panic(err)
	}

	if !IsTarget(path, info) {
		return nil
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("【02】")
		panic(err)
	}

	out := &TestCodeInfo{PkgName: GetPackageName(fset.Position(f.Package).Filename)}

	fncs := []*Function{}
	// TODO astだけでなくtypesも駆使すれば解析の幅も広がるはず
	// TODO importする3rdパーティモジュールに応じたプラグイン処理を入れる（aws-sdk-goがあったら特定のテストコードロジックを自動生成するなど）
	for _, decl := range f.Decls {
		//ast.Print(fset, decl)
		switch declType := decl.(type) {
		case *ast.GenDecl:
			switch declType.Tok {
			// MEMO パッケージ名が取れない...
			//case token.PACKAGE:
			//	fmt.Printf("PACKAGE: %#v\n", declType)
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

	tmpl := template.Must(template.ParseFiles("./testcode.md"))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, out)
	if err != nil {
		fmt.Printf("【03】path: %v\n", path)
		panic(err)
	}

	fp, err := GetOutputFile(path)
	if err != nil {
		fmt.Printf("【04】path: %v\n", path)
		panic(err)
	}
	defer fp.Close()

	w := bufio.NewWriter(fp)
	w.WriteString(buf.String())
	w.Flush()

	return nil
}

// TODO フィルタリング機能としてパッケージ分け
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

// TODO ast経由での取得の仕方を模索
func GetPackageName(path string) string {
	fl, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fl.Close()
	sc := bufio.NewScanner(fl)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			break
		}
		// TODO 適当すぎなので、もう少し洗練されたやり方に直す
		if strings.Contains(sc.Text(), "package ") {
			return strings.Split(sc.Text(), "package ")[1]
		}
	}
	return path
}

// TODO とりあえず適当実装なので見直す
func GetOutputFile(path string) (*os.File, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	absDir := filepath.Dir(absPath)
	absFile := filepath.Base(absPath)
	absFileNm := strings.Split(absFile, ".")[0]
	absTestFile := fmt.Sprintf("%s_test.go", absFileNm)
	absOutFile := filepath.Join(absDir, absTestFile)
	if Exists(absOutFile) {
		return nil, err
	}
	fl, err := os.Create(absOutFile)
	if err != nil {
		return nil, err
	}
	return fl, nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
