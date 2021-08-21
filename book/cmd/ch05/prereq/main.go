package main

import (
	"fmt"

	"book/ch05"
)

var prerequisits = map[string][]string{
	"algorithm": {"data structures"},
	"calculus":  {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"compure organization",
	},
	"data structures":      {"discrete math"},
	"databases":            {"data structures"},
	"discrete math":        {"intro to programming"},
	"formal languages":     {"discrete math"},
	"networks":             {"operating systems"},
	"operating systems":    {"data structures", "computer organization"},
	"programing languages": {"data structures", "computer organization"},
}

func main() {
	order := ch05.TopoSort(prerequisits)
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
