package ast

import "fmt"

// OpType declares the operation type.
type OpType int

// All available operation types.
const (
	OpAdd OpType = iota
	OpSub
	OpMul
	OpDiv
)

var types = [...]string{
	"+",
	"-",
	"*",
	"/",
}

// Operation is a arithmetic operation.
type Operation struct {
	OpType OpType
	Left   Node
	Right  Node
}

func (o *Operation) String() string {
	return fmt.Sprintf("(op `%v` left:%v right:%v)", types[o.OpType], o.Left, o.Right)
}

// Children returns the node's children.
func (o *Operation) Children() []Node {
	return []Node{o.Left, o.Right}
}
