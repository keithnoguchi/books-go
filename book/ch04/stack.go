package ch04

type Stack []interface{}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() interface{} {
	n := len(*s) - 1
	v := (*s)[n]
	*s = (*s)[:n]
	return v
}
