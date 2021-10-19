package main

import (
	"fmt"
	"os"

	"github.com/fmarmol/basename/pkg/basename"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %v filepath", os.Args[0])
		os.Exit(1)
	}
	fp := os.Args[1]
	fileInfo := basename.ParseFile(fp)
	fmt.Println(fileInfo.Basename, fileInfo.Ext)
}
