package ast

import "fmt"

// Assign is an assignment statement.
type Assign struct {
	Left  Node
	Right Node
}

func (a *Assign) String() string {
	return fmt.Sprintf("(assign %v = %v)", a.Left, a.Right)
}

// Children returns the node's children.
func (a *Assign) Children() []Node {
	return nil
}
