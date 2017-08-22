package ast

import "fmt"

// OpType declares the operation type.
type OpType int

// All available operation types.
const (
	_ OpType = iota
	OpAdd
	OpSub
	OpMul
	OpDiv
)

// Operation is a arithmetic operation.
type Operation struct {
	OpType OpType
	Name   string
	Left   Node
	Right  Node
}

func (o *Operation) String() string {
	return fmt.Sprintf("(op `%v` left:%v right:%v)", o.Name, o.Left, o.Right)
}

// Children returns the node's children.
func (o *Operation) Children() []Node {
	return []Node{o.Left, o.Right}
}
