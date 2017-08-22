// Tik is an interpreted programming language.
package main

import (
	"fmt"
	"os"

	"github.com/pseidemann/tik/inspect"
	"github.com/pseidemann/tik/interpreter"
	"github.com/pseidemann/tik/parser"
	"github.com/pseidemann/tik/parser/lexer"
)

func main() {
	testFile("examples/functions.tik")
	testFile("examples/hello.tik")
	testFile("examples/math.tik")
	testFile("examples/vars.tik")
}

func testFile(filename string) {
	fmt.Println("### test file", filename)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR: failed to open file:", err)
	}

	lex := lexer.New(f)
	par := parser.New(lex)
	ast := par.CreateAST()

	fmt.Println("--- print ast")
	inspect.PrintAST(ast)

	fmt.Println("--- execute ast")
	in := interpreter.New()
	in.Execute(ast)

	fmt.Println("--- done")
}
