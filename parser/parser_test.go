package parser

import (
	"os"
	"reflect"
	"testing"

	"github.com/pseidemann/tik/ast"
	"github.com/pseidemann/tik/lexer"
)

func TestFunctions(t *testing.T) {
	f, err := os.Open("../testdata/functions.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := lexer.New(f)
	par := New(lex)
	a := par.CreateAST()

	expected := &ast.Block{
		Name: "main",
		Stmts: []ast.Node{
			&ast.Assign{
				Left:  &ast.Ident{Name: "outer"},
				Right: &ast.Number{Num: "1"},
			},
			&ast.FuncDef{
				Name: "greet",
				Params: []*ast.Param{
					&ast.Param{Name: "a"},
					&ast.Param{Name: "b"},
				},
				Body: &ast.Block{
					Name: "func",
					Stmts: []ast.Node{
						&ast.Assign{
							Left:  &ast.Ident{Name: "inner"},
							Right: &ast.Number{Num: "4"},
						},
						&ast.FuncCall{
							Name: "print",
							Args: []ast.Node{
								&ast.Ident{Name: "outer"},
								&ast.Ident{Name: "a"},
								&ast.Ident{Name: "b"},
								&ast.Ident{Name: "inner"},
							},
						},
					},
				},
			},
			&ast.FuncCall{
				Name: "greet",
				Args: []ast.Node{
					&ast.Number{Num: "2"},
					&ast.Number{Num: "3"},
				},
			},
		},
	}

	if !reflect.DeepEqual(a, expected) {
		t.Error("unexpected AST")
	}
}

func TestMath(t *testing.T) {
	f, err := os.Open("../testdata/math.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := lexer.New(f)
	par := New(lex)
	a := par.CreateAST()

	expected := &ast.Block{
		Name: "main",
		Stmts: []ast.Node{
			&ast.FuncCall{
				Name: "print",
				Args: []ast.Node{
					&ast.Operation{
						OpType: ast.OpAdd,
						Left:   &ast.Number{Num: "10"},
						Right: &ast.Operation{
							OpType: ast.OpMul,
							Left:   &ast.Number{Num: "2"},
							Right: &ast.Operation{
								OpType: ast.OpSub,
								Left:   &ast.Number{Num: "31"},
								Right: &ast.Operation{
									OpType: ast.OpDiv,
									Left:   &ast.Number{Num: "2"},
									Right:  &ast.Number{Num: "2"},
								},
							},
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(a, expected) {
		t.Error("unexpected AST")
	}
}

func TestPrint(t *testing.T) {
	f, err := os.Open("../testdata/print.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := lexer.New(f)
	par := New(lex)
	a := par.CreateAST()

	expected := &ast.Block{
		Name: "main",
		Stmts: []ast.Node{
			&ast.FuncCall{
				Name: "print",
				Args: []ast.Node{
					&ast.String{Str: "hello"},
				},
			},
			&ast.FuncCall{
				Name: "print",
				Args: []ast.Node{
					&ast.String{Str: "world1"},
					&ast.String{Str: "world2"},
				},
			},
			&ast.FuncCall{
				Name: "print",
				Args: []ast.Node{
					&ast.String{Str: "world3"},
					&ast.Operation{
						OpType: ast.OpAdd,
						Left:   &ast.Number{Num: "1"},
						Right: &ast.Operation{
							OpType: ast.OpMul,
							Left:   &ast.Number{Num: "2"},
							Right:  &ast.Number{Num: "3"},
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(a, expected) {
		t.Error("unexpected AST")
	}
}

func TestVariables(t *testing.T) {
	f, err := os.Open("../testdata/variables.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := lexer.New(f)
	par := New(lex)
	a := par.CreateAST()

	expected := &ast.Block{
		Name: "main",
		Stmts: []ast.Node{
			&ast.Assign{
				Left: &ast.Ident{Name: "vara"},
				Right: &ast.Operation{
					OpType: ast.OpAdd,
					Left:   &ast.Number{Num: "1"},
					Right:  &ast.Number{Num: "2"},
				},
			},
			&ast.Assign{
				Left: &ast.Ident{Name: "varb"},
				Right: &ast.Operation{
					OpType: ast.OpMul,
					Left:   &ast.Number{Num: "2"},
					Right:  &ast.Number{Num: "4"},
				},
			},
			&ast.Assign{
				Left: &ast.Ident{Name: "varc"},
				Right: &ast.Operation{
					OpType: ast.OpAdd,
					Left:   &ast.Ident{Name: "vara"},
					Right:  &ast.Ident{Name: "varb"},
				},
			},
			&ast.FuncCall{
				Name: "print",
				Args: []ast.Node{
					&ast.Ident{Name: "vara"},
					&ast.Ident{Name: "varb"},
					&ast.Ident{Name: "varc"},
				},
			},
		},
	}

	if !reflect.DeepEqual(a, expected) {
		t.Error("unexpected AST")
	}
}
