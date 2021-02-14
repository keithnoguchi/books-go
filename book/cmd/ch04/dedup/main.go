// dedup: shows the first occurence of the line.
package main

import (
	"os"

	"book/ch04"
)

func main() {
	ch04.Dedup(os.Stdin, os.Stdout)
}
