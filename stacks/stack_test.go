package stacks

import "testing"

func TestPop(t *testing.T) {

	s := NewStack()
	s.Push(5)
	var val int
	val, _ = s.Pop()
	if val != 5 {
		t.Error("Expected 5, got ",val)
	}
}
