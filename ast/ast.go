// Package ast contains data structures for the AST.
package ast

// Node is an AST node.
type Node interface {
	String() string
	Children() []Node
}
