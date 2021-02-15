// kv shows the map example with the slice of the string key
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"book/ch04"
)

func main() {
	kv := ch04.NewKV()
	input := bufio.NewScanner(os.Stdin)
	lines := map[string][]string{}
	for input.Scan() {
		words := strings.Split(input.Text(), " ")
		key := kv.Add(words)
		lines[key] = words
	}
	for _, words := range lines {
		fmt.Printf("%v: %d\n", words, kv.Count(words))
	}
}
