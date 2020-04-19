// Dup1 prints the text of each line that appears more than
// Also prints the name of source file
package ch1

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
)

var palette = []color.Color{color.White, color.Black}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]

	for _, file := range files {
		file, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		} else {
			// pass by value : here values is the copy of reference
			countLines(file, counts)
			file.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
