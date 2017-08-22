package ast

import "fmt"

// Param is an argument in a function declaration.
type Param struct {
	Name string
}

func (p *Param) String() string {
	return fmt.Sprintf("(param=%v)", p.Name)
}

// Children returns the node's children.
func (p *Param) Children() []Node {
	return nil
}
