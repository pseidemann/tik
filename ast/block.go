package ast

import "fmt"

// Block is a chunk of statements.
type Block struct {
	Name  string
	Stmts []Node
}

func (b *Block) String() string {
	return fmt.Sprintf("(block=%v)", b.Name)
}

// Children returns the node's children.
func (b *Block) Children() []Node {
	return b.Stmts
}
