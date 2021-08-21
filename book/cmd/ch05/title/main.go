package main

import (
	"fmt"
	"os"

	"book/ch05"
)

func main() {
	title, err := ch05.Title(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "title: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(title)
}
