package parser

import (
	"testing"

	"github.com/pseidemann/tik/parser/lexer"
)

func TestLIFO(t *testing.T) {
	var s tokenStack

	t1 := &lexer.Token{Value: "1"}
	t2 := &lexer.Token{Value: "2"}
	t3 := &lexer.Token{Value: "3"}

	s.push(t1)
	s.push(t2)
	s.push(t3)

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
