package main

import (
	"fmt"

	"go.com/practice/pkg/ch1"
)

func main() {
	fmt.Printf("Hello\n")
	fmt.Printf("%+v\n", *ch1.NewUsage())
}
