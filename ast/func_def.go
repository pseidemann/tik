package ast

import "fmt"

// FuncDef is the declaration of a function.
type FuncDef struct {
	Name   string
	Params []*Param
	Body   *Block
}

func (f *FuncDef) String() string {
	return fmt.Sprintf("(func=%v %v)", f.Name, f.Params)
}

// Children returns the node's children.
func (f *FuncDef) Children() []Node {
	return []Node{f.Body}
}
