package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"book/ch04"
	"github.com/google/uuid"
)

var (
	addr = flag.String("addr", "", "listening address")
	port = flag.Int("port", 8011, "listening port")
)

func main() {
	flag.Parse()
	if *addr == "" {
		cmdHandler()
	}
}

func cmdHandler() {
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
	n, err := strconv.ParseInt(os.Args[1], 10, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "uuid: %v\n", err)
		os.Exit(1)
	}
	counter := ch04.NewUUIDCounter()
	for i := 0; i < int(n); i++ {
		b, err := uuid.New().MarshalBinary()
		if err != nil {
			fmt.Fprintf(os.Stderr, "UUID.MarshalBinary: %v", err)
			os.Exit(1)
		}
		counter.Count(b)
	}
	for k, v := range *counter {
		fmt.Printf("%s(%d)->%d\n", k, len(k), v)
	}
}
