// Package inspect implements printing of AST nodes.
package inspect

import (
	"fmt"
	"strings"

	"github.com/pseidemann/tik/parser/ast"
)

// PrintAST prints AST nodes.
func PrintAST(root ast.Node) {
	print(root, 0)
}

func print(n ast.Node, depth int) {
	indent := strings.Repeat("    ", depth)
	fmt.Printf("%s|__ %s\n", indent, n)
	depth++
	for _, child := range n.Children() {
		print(child, depth)
	}
}
