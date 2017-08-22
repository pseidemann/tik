package ast

// Number is a mathematical object.
type Number struct {
	Num string
}

func (n *Number) String() string {
	return n.Num
}

// Children returns the node's children.
func (n *Number) Children() []Node {
	return nil
}
