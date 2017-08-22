package interpreter

import "testing"

func TestLIFO(t *testing.T) {
	var s contextStack

	t1 := &context{}
	t2 := &context{}
	t3 := &context{}

	s.push(t1)
	s.push(t2)
	s.push(t3)

	if s.size() != 3 {
		t.Error("wrong size")
	}

	if s.peek() != t3 {
		t.Error("peek() must return last pushed element")
	}

	if s.pop() != t3 {
		t.Error("expected t3")
	}
	if s.pop() != t2 {
		t.Error("expected t2")
	}
	if s.pop() != t1 {
		t.Error("expected t1")
	}

	if s.peek() != nil {
		t.Error("expected emtpy stack")
	}
	if s.pop() != nil {
		t.Error("expected emtpy stack")
	}
}
