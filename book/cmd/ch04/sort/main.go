// sort sorts slice of integer with the binary tree inserssion
package main

import (
	"fmt"
	"os"
	"strconv"

	"book/ch04"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: sort 3 4 2 5 2\n")
		os.Exit(0)
	}
	data := make([]int, 0, len(os.Args)-1)
	for _, num := range os.Args[1:] {
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Fprintf(os.Stderr, "sort: invalid value: %q\n", num)
			os.Exit(1)
		}
		data = append(data, n)
	}
	ch04.Sort(data)
	fmt.Println(data)
}
