// github shows the issue result.
package main

import (
	"fmt"
	"log"
	"os"

	"book/ch04"
)

func main() {
	result, err := ch04.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
