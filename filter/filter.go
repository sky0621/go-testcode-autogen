package filter

import (
	"os"

	"path/filepath"

	"fmt"

	"regexp"

	"github.com/sky0621/go-testcode-autogen/config"
)

type FilterManager struct {
	Cfg *config.FilterConfig
}

func (m *FilterManager) IsTarget(path string, info os.FileInfo) bool {
	fmt.Println("JJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJ")
	fmt.Printf("Cfg: %#v\n", m.Cfg)
	fmt.Println("JJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJ")
	if info.IsDir() {
		return false
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return false
	}
	fmt.Println("YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY")
	fmt.Printf("path: %v, absPath: %v\n", path, absPath)
	fmt.Println("YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY")

	outDirs := m.Cfg.OutDir
	fmt.Printf("outDirs: %v\n", outDirs)
	for _, outDir := range outDirs {
		outDirExp, err := regexp.Compile(outDir)
		if err != nil {
			return false
		}
		if outDirExp.MatchString(absPath) {
			fmt.Printf("[OUT_DIR]absPath: %v, outDir: %v\n", absPath, outDir)
			return false
		}
	}

	outFiles := m.Cfg.OutFile
	for _, outFile := range outFiles {
		outFileExp, err := regexp.Compile(outFile)
		if err != nil {
			return false
		}
		if outFileExp.MatchString(path) {
			fmt.Printf("[OUT_FILE]path: %v, outFile: %v\n", path, outFile)
			return false
		}
	}

	inDirs := m.Cfg.InDir
	for _, inDir := range inDirs {
		inDirExp, err := regexp.Compile(inDir)
		if err != nil {
			return false
		}
		if !inDirExp.MatchString(absPath) {
			fmt.Printf("[IN_DIR]absPath: %v, inDir: %v\n", absPath, inDir)
			return false
		}
	}

	inFiles := m.Cfg.InFile
	for _, inFile := range inFiles {
		inFileExp, err := regexp.Compile(inFile)
		if err != nil {
			return false
		}
		if !inFileExp.MatchString(path) {
			fmt.Printf("[IN_FILE]path: %v, inFile: %v\n", path, inFile)
			return false
		}
	}

	return true
}
