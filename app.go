package go_testcode_autogen

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/sky0621/go-testcode-autogen/config"
	"github.com/sky0621/go-testcode-autogen/filter"
	"github.com/sky0621/go-testcode-autogen/inspect"
	"github.com/sky0621/go-testcode-autogen/testinfo"
)

func Apply(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fmngr := &filter.FilterManager{Filter: config.NewConfig().Filter}
	if !fmngr.IsTarget(path, info) {
		return nil
	}
	fmt.Println("\n######################################################################################################################################\n")
	fmt.Println(path)
	fmt.Println("\n######################################################################################################################################\n")

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		return err
	}

	testInfo := &testinfo.TestInfo{}
	ast.Inspect(f, func(node ast.Node) bool {
		inspector := inspect.GetInspector(node)
		if inspector == nil {
			// TODO WARNログ
			return true
		}
		// TODO astだけでなくtypesも駆使すれば解析の幅も広がるはず
		// TODO importする3rdパーティモジュールに応じたプラグイン処理を入れる（aws-sdk-goがあったら特定のテストコードロジックを自動生成するなど）
		err := inspector.Inspect(node, testInfo)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	})

	tmpl := template.Must(template.ParseFiles("../template/testcode.md"))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, testInfo)
	if err != nil {
		return err
	}

	fp, err := GetOutputFile(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	w := bufio.NewWriter(fp)
	w.WriteString(buf.String())
	w.Flush()

	return nil
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
