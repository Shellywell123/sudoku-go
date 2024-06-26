package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	file = flag.String("file", "", "file-name")
)

// entrypoint
func main() {

	flag.Parse()

	fmt.Println("\nfile name: ", *file)

	if *file == "" {
		fmt.Println("please specify a file")
		os.Exit(0)
	}

	s := ImportSudokuFromFile(*file)
	PrintSudoku(s)
	s.Solve()
	PrintSudoku(s)
}
