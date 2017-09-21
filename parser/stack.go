package parser

import "github.com/pseidemann/tik/lexer"

type tokenStack struct {
	s []*lexer.Token
}

func (s *tokenStack) push(t *lexer.Token) {
	s.s = append(s.s, t)
}

func (s *tokenStack) pop() *lexer.Token {
	l := len(s.s)
	if l == 0 {
		return nil
	}
	var t *lexer.Token
	t, s.s = s.s[l-1], s.s[:l-1]
	return t
}

func (s *tokenStack) peek() *lexer.Token {
	l := len(s.s)
	if l == 0 {
		return nil
	}
	return s.s[l-1]
}
