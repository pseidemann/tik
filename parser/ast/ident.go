package ast

import "fmt"

// Ident is a variable.
type Ident struct {
	Name string
}

func (i *Ident) String() string {
	return fmt.Sprintf("(ident=%v)", i.Name)
}

// Children returns the node's children.
func (i *Ident) Children() []Node {
	return nil
}
