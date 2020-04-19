// echo1 command
package ch1

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo() {
	start := time.Now()
	method1()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("method1 total time : %v\n", elapsed)

	start = time.Now()
	method2()
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("method2 total time : %v\n", elapsed)
}

func method1() {
	for index, arg := range os.Args[0:] {
		fmt.Printf("%d \t %s\n", index, arg)
	}
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func method2() {
	var sep, res string
	for _, arg := range os.Args[0:] {
		res += sep + arg
		sep = " "
	}
	fmt.Println(res)
}
