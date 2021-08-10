package ch04

import "log"

func Remove(a []interface{}, pos int) []interface{} {
	if pos >= len(a) {
		log.Printf("pos(%d) >= len(a)\n", pos)
		return a
	}
	copy(a[pos:], a[pos+1:])
	return a[:len(a)-1]
}

func Remove2(a []interface{}, pos int) []interface{} {
	if pos >= len(a) {
		log.Printf("pos(%d) >= len(a)", pos)
		return a
	}
	n := len(a) - 1
	a[pos] = a[n]
	return a[:n]
}
