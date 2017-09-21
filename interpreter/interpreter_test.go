package interpreter

import (
	"bytes"
	"os"
	"testing"

	"github.com/pseidemann/tik/parser"
	"github.com/pseidemann/tik/parser/lexer"
)

func TestFunctions(t *testing.T) {
	f, err := os.Open("../testdata/functions.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := lexer.New(f)
	par := parser.New(lex)
	a := par.CreateAST()
	var out bytes.Buffer
	in := New(&out)
	in.Execute(a)

	expected := "1 2 3 4\n"

	if out.String() != expected {
		t.Error("unexpected output")
	}
}

func TestMath(t *testing.T) {
	f, err := os.Open("../testdata/math.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := lexer.New(f)
	par := parser.New(lex)
	a := par.CreateAST()
	var out bytes.Buffer
	in := New(&out)
	in.Execute(a)

	expected := "70\n"

	if out.String() != expected {
		t.Error("unexpected output")
	}
}

func TestPrint(t *testing.T) {
	f, err := os.Open("../testdata/print.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := lexer.New(f)
	par := parser.New(lex)
	a := par.CreateAST()
	var out bytes.Buffer
	in := New(&out)
	in.Execute(a)

	expected := "hello\nworld1 world2\nworld3 7\n"

	if out.String() != expected {
		t.Error("unexpected output")
	}
}

func TestVariables(t *testing.T) {
	f, err := os.Open("../testdata/variables.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := lexer.New(f)
	par := parser.New(lex)
	a := par.CreateAST()
	var out bytes.Buffer
	in := New(&out)
	in.Execute(a)

	expected := "3 8 11\n"

	if out.String() != expected {
		t.Error("unexpected output")
	}
}
