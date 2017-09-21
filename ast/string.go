package ast

import "fmt"

// String is a sequence of characters.
type String struct {
	Str string
}

func (s *String) String() string {
	return fmt.Sprintf("(str=%q)", s.Str)
}

// Children returns the node's children.
func (s *String) Children() []Node {
	return nil
}
