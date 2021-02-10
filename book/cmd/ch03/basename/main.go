// basename, go version of unix basename.
package main

import (
	"fmt"

	"book/ch03/basename/v2"
)

func main() {
	fmt.Println(basename.Basename("a/b/c.go"))
}
