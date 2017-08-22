package ast

import "fmt"

// FuncCall is the calling of a function.
type FuncCall struct {
	Name string
	Args []Node
}

func (f *FuncCall) String() string {
	return fmt.Sprintf("(funccall=%v %v)", f.Name, f.Args)
}

// Children returns the node's children.
func (f *FuncCall) Children() []Node {
	return nil
}
