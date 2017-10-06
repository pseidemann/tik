package interpreter

import (
	"bytes"
	"os"
	"testing"

	"github.com/pseidemann/tik/lexer"
	"github.com/pseidemann/tik/parser"
)

func TestFuncArgs(t *testing.T) {
	f, err := os.Open("../testdata/func_args.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := lexer.New(f)
	par := parser.New(lex)
	a := par.CreateAST()
	var out bytes.Buffer
	in := New(&out)
	in.Execute(a)

	expected := "1 2\n"

	if out.String() != expected {
		t.Error("unexpected output")
	}
}

func TestFuncReturn(t *testing.T) {
	f, err := os.Open("../testdata/func_return.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := lexer.New(f)
	par := parser.New(lex)
	a := par.CreateAST()
	var out bytes.Buffer
	in := New(&out)
	in.Execute(a)

	expected := "4\n"

	if out.String() != expected {
		t.Error("unexpected output")
	}
}

func TestFuncScope(t *testing.T) {
	f, err := os.Open("../testdata/func_scope.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := lexer.New(f)
	par := parser.New(lex)
	a := par.CreateAST()
	var out bytes.Buffer
	in := New(&out)
	in.Execute(a)

	expected := "1 2 3\n"

	if out.String() != expected {
		t.Error("unexpected output")
	}
}

func TestFuncSimple(t *testing.T) {
	f, err := os.Open("../testdata/func_simple.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := lexer.New(f)
	par := parser.New(lex)
	a := par.CreateAST()
	var out bytes.Buffer
	in := New(&out)
	in.Execute(a)

	expected := "Hello, world!\n"

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
