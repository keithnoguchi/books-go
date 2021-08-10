package main

import (
	"fmt"

	"book/ch04"
)

func main() {
	var a []interface{}
	fmt.Printf("&a=%p,a=%p\n", &a, a)

	/// single argument way.
	for _, b := range []interface{}{1, 2, 3, 4, 5, 6, 7, 8} {
		c := ch04.Append(a, b)
		fmt.Printf(
			"&a=%p,a=%p,len(a)=%02d,cap(a)=%02d,&c=%p,c=%p,len(c)=%02d,cap(c)=%02d\t%[6]v\n",
			&a, a, len(a), cap(a),
			&c, c, len(c), cap(c),
		)
		a = c
	}

	// variadic way.
	fmt.Println()
	a = nil
	fmt.Printf(
		"&a=%p,a=%p,len(a)=%02d,cap(a)=%02d\n",
		&a, a, len(a), cap(a),
	)
	c := ch04.Append(a, []interface{}{1, 2, 3, 4, 5}...)
	fmt.Printf(
		"&a=%p,a=%p,len(a)=%02d,cap(a)=%02d,"+
			"&c=%p,c=%p,len(c)=%02d,cap(c)=%02d\t%[6]v\n",
		&a, a, len(a), cap(a),
		&c, c, len(c), cap(c),
	)
}
