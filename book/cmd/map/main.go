package main

import "fmt"

func main() {
	var maps []map[string]int
	fmt.Printf("maps: &maps=%p,maps=%p,maps=%[2]v\n", &maps, maps)

	var m map[string]int
	fmt.Printf("m:  &m=%p,m=%p,m=%[2]v\n", &m, m)
	maps = append(maps, m)
	m = map[string]int{}
	maps = append(maps, m)
	m = make(map[string]int)
	maps = append(maps, m)
	for i, m := range maps {
		fmt.Printf("m%d: &m=%p,m=%p,m=%[3]v\n", i, &m, m)
	}
	fmt.Printf("m:  &m=%p,m=%p,m=%[2]v\n", &m, m)
}
