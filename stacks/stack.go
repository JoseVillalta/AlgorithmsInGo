package stacks

import "errors"

type stack struct {
	data []int
}

func NewStack() * stack {
	return &stack {make([]int,0)}
}

func (s * stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s * stack) Pop() (int, error) {
	l := len(s.data)
	if l == 0 {
		return 0, errors.New("empty stack")
	}

	res := s.data[l-1]
	s.data = s.data[:l-1]
	return res, nil
}

