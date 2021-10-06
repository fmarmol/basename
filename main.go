package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %v filepath", os.Args[0])
		os.Exit(1)
	}
	fp := os.Args[1]
	base := filepath.Base(fp)
	ext := filepath.Ext(fp)
	fmt.Println(strings.ReplaceAll(base, ext, ""), strings.TrimLeft(ext, "."))
}
