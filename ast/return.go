package ast

import "fmt"

// Return exits a function with an optional value.
type Return struct {
	Value Node
}

func (r *Return) String() string {
	return fmt.Sprintf("(return %v)", r.Value)
}

// Children returns the node's children.
func (r *Return) Children() []Node {
	return nil
}
