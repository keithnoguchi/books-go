// lissajous: it generates GIF animations of random Lissajous figures.
package main

import (
	"os"

	"book/ch01"
)

func main() {
	ch01.Lissajous(os.Stdout)
}
