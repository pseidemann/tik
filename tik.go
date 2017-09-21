// Tik is an interpreted programming language.
package main

import (
	"fmt"
	"os"

	"github.com/pseidemann/tik/inspect"
	"github.com/pseidemann/tik/interpreter"
	"github.com/pseidemann/tik/lexer"
	"github.com/pseidemann/tik/parser"
)

func main() {
	testFile("testdata/functions.tik")
	testFile("testdata/math.tik")
	testFile("testdata/print.tik")
	testFile("testdata/variables.tik")
}

func testFile(filename string) {
	fmt.Println("### test file", filename)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR: failed to open file:", err)
	}

	lex := lexer.New(f)
	par := parser.New(lex)
	a := par.CreateAST()

	fmt.Println("--- print ast")
	inspect.PrintAST(a)

	fmt.Println("--- execute ast")
	in := interpreter.New(os.Stdout)
	in.Execute(a)

	fmt.Println("--- done")
}
