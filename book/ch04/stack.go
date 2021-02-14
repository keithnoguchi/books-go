package ch04

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("zero size stack")
	}
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x, nil
}
