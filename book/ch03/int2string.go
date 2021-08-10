package ch03

import (
	"bytes"
	"fmt"
	"log"
)

func IntsToString(ints []int) string {
	var buf bytes.Buffer
	if err := buf.WriteByte('['); err != nil {
		log.Println(err)
	}
	for i, v := range ints {
		if i != 0 {
			if _, err := buf.WriteString(", "); err != nil {
				log.Println(err)
			}
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	if err := buf.WriteByte(']'); err != nil {
		log.Println(err)
	}
	return buf.String()
}
