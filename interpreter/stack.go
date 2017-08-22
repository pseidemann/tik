package interpreter

type contextStack struct {
	s []*context
}

func (s *contextStack) push(t *context) {
	s.s = append(s.s, t)
}

func (s *contextStack) pop() *context {
	l := len(s.s)
	if l == 0 {
		return nil
	}
	var t *context
	t, s.s = s.s[l-1], s.s[:l-1]
	return t
}

func (s *contextStack) peek() *context {
	l := len(s.s)
	if l == 0 {
		return nil
	}
	return s.s[l-1]
}

func (s *contextStack) size() int {
	return len(s.s)
}
