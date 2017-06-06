package main

import (
	"flag"
	"path/filepath"

	ag "github.com/sky0621/go-testcode-autogen"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	target := flag.String("target", "../example/sampleproject", "Parse Target")
	flag.Parse()

	// TODO config読み込んでロジックに反映

	filepath.Walk(*target, ag.Apply)
}
