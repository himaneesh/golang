// Dup1 prints the text of each line that appears more than
package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func dup() {
	countmap := make(map[string]int16)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		countmap[input.Text()]++
	}

	for line, count := range countmap {
		if count > 1 {
			fmt.Printf("%d \t %s\n", count, line)
		}
	}
}
