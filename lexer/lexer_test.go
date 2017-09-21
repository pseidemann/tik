package lexer

import (
	"os"
	"reflect"
	"testing"
)

func TestFunctions(t *testing.T) {
	f, err := os.Open("../testdata/functions.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := New(f)

	var out []*Token

	for {
		tok, err := lex.NextToken()
		if err != nil {
			if err != ErrEOF {
				t.Error("expected EOF error")
			}
			break
		}
		out = append(out, tok)
	}

	expected := []*Token{
		&Token{TokenType: "identifier", Value: "outer"},
		&Token{TokenType: "assignment", Precedence: 100},
		&Token{TokenType: "number", Value: "1"},
		&Token{TokenType: "newline", Precedence: 1000},
		&Token{TokenType: "newline", Precedence: 1000},
		&Token{TokenType: "keyword", Precedence: 10, Value: "func"},
		&Token{TokenType: "identifier", Value: "greet"},
		&Token{TokenType: "paren-left", Precedence: 100},
		&Token{TokenType: "identifier", Value: "a"},
		&Token{TokenType: "comma"},
		&Token{TokenType: "identifier", Value: "b"},
		&Token{TokenType: "paren-right", Precedence: 100},
		&Token{TokenType: "brace-left"},
		&Token{TokenType: "newline", Precedence: 1000},
		&Token{TokenType: "identifier", Value: "inner"},
		&Token{TokenType: "assignment", Precedence: 100},
		&Token{TokenType: "number", Value: "4"},
		&Token{TokenType: "newline", Precedence: 1000},
		&Token{TokenType: "keyword", Precedence: 10, Value: "print"},
		&Token{TokenType: "paren-left", Precedence: 100},
		&Token{TokenType: "identifier", Value: "outer"},
		&Token{TokenType: "comma"},
		&Token{TokenType: "identifier", Value: "a"},
		&Token{TokenType: "comma"},
		&Token{TokenType: "identifier", Value: "b"},
		&Token{TokenType: "comma"},
		&Token{TokenType: "identifier", Value: "inner"},
		&Token{TokenType: "paren-right", Precedence: 100},
		&Token{TokenType: "newline", Precedence: 1000},
		&Token{TokenType: "brace-right"},
		&Token{TokenType: "newline", Precedence: 1000},
		&Token{TokenType: "newline", Precedence: 1000},
		&Token{TokenType: "identifier", Value: "greet"},
		&Token{TokenType: "paren-left", Precedence: 100},
		&Token{TokenType: "number", Value: "2"},
		&Token{TokenType: "comma"},
		&Token{TokenType: "number", Value: "3"},
		&Token{TokenType: "paren-right", Precedence: 100},
		&Token{TokenType: "newline", Precedence: 1000},
	}

	if !reflect.DeepEqual(out, expected) {
		t.Error("unexpected token output")
	}
}

func TestMath(t *testing.T) {
	f, err := os.Open("../testdata/math.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := New(f)

	var out []*Token

	for {
		tok, err := lex.NextToken()
		if err != nil {
			if err != ErrEOF {
				t.Error("expected EOF error")
			}
			break
		}
		out = append(out, tok)
	}

	expected := []*Token{
		&Token{TokenType: "keyword", Precedence: 10, Value: "print"},
		&Token{TokenType: "paren-left", Precedence: 100},
		&Token{TokenType: "number", Value: "10"},
		&Token{TokenType: "operator", Precedence: 1, Value: "+"},
		&Token{TokenType: "number", Value: "2"},
		&Token{TokenType: "operator", Precedence: 2, Value: "*"},
		&Token{TokenType: "paren-left", Precedence: 100},
		&Token{TokenType: "number", Value: "31"},
		&Token{TokenType: "operator", Precedence: 1, Value: "-"},
		&Token{TokenType: "number", Value: "2"},
		&Token{TokenType: "operator", Precedence: 2, Value: "/"},
		&Token{TokenType: "number", Value: "2"},
		&Token{TokenType: "paren-right", Precedence: 100},
		&Token{TokenType: "paren-right", Precedence: 100},
		&Token{TokenType: "newline", Precedence: 1000},
	}

	if !reflect.DeepEqual(out, expected) {
		t.Error("unexpected token output")
	}
}

func TestPrint(t *testing.T) {
	f, err := os.Open("../testdata/print.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := New(f)

	var out []*Token

	for {
		tok, err := lex.NextToken()
		if err != nil {
			if err != ErrEOF {
				t.Error("expected EOF error")
			}
			break
		}
		out = append(out, tok)
	}

	expected := []*Token{
		&Token{TokenType: "keyword", Precedence: 10, Value: "print"},
		&Token{TokenType: "paren-left", Precedence: 100},
		&Token{TokenType: "string", Value: "hello"},
		&Token{TokenType: "paren-right", Precedence: 100},
		&Token{TokenType: "newline", Precedence: 1000},
		&Token{TokenType: "keyword", Precedence: 10, Value: "print"},
		&Token{TokenType: "paren-left", Precedence: 100},
		&Token{TokenType: "string", Value: "world1"},
		&Token{TokenType: "comma"},
		&Token{TokenType: "string", Value: "world2"},
		&Token{TokenType: "paren-right", Precedence: 100},
		&Token{TokenType: "newline", Precedence: 1000},
		&Token{TokenType: "keyword", Precedence: 10, Value: "print"},
		&Token{TokenType: "paren-left", Precedence: 100},
		&Token{TokenType: "string", Value: "world3"},
		&Token{TokenType: "comma"},
		&Token{TokenType: "number", Value: "1"},
		&Token{TokenType: "operator", Precedence: 1, Value: "+"},
		&Token{TokenType: "number", Value: "2"},
		&Token{TokenType: "operator", Precedence: 2, Value: "*"},
		&Token{TokenType: "number", Value: "3"},
		&Token{TokenType: "paren-right", Precedence: 100},
		&Token{TokenType: "newline", Precedence: 1000},
	}

	if !reflect.DeepEqual(out, expected) {
		t.Error("unexpected token output")
	}
}

func TestVariables(t *testing.T) {
	f, err := os.Open("../testdata/variables.tik")
	if err != nil {
		t.Fatal(err)
	}
	lex := New(f)

	var out []*Token

	for {
		tok, err := lex.NextToken()
		if err != nil {
			if err != ErrEOF {
				t.Error("expected EOF error")
			}
			break
		}
		out = append(out, tok)
	}

	expected := []*Token{
		&Token{TokenType: "identifier", Value: "vara"},
		&Token{TokenType: "assignment", Precedence: 100},
		&Token{TokenType: "number", Value: "1"},
		&Token{TokenType: "operator", Precedence: 1, Value: "+"},
		&Token{TokenType: "number", Value: "2"},
		&Token{TokenType: "newline", Precedence: 1000},
		&Token{TokenType: "identifier", Value: "varb"},
		&Token{TokenType: "assignment", Precedence: 100},
		&Token{TokenType: "number", Value: "2"},
		&Token{TokenType: "operator", Precedence: 2, Value: "*"},
		&Token{TokenType: "number", Value: "4"},
		&Token{TokenType: "newline", Precedence: 1000},
		&Token{TokenType: "identifier", Value: "varc"},
		&Token{TokenType: "assignment", Precedence: 100},
		&Token{TokenType: "identifier", Value: "vara"},
		&Token{TokenType: "operator", Precedence: 1, Value: "+"},
		&Token{TokenType: "identifier", Value: "varb"},
		&Token{TokenType: "newline", Precedence: 1000},
		&Token{TokenType: "keyword", Precedence: 10, Value: "print"},
		&Token{TokenType: "paren-left", Precedence: 100},
		&Token{TokenType: "identifier", Value: "vara"},
		&Token{TokenType: "comma"},
		&Token{TokenType: "identifier", Value: "varb"},
		&Token{TokenType: "comma"},
		&Token{TokenType: "identifier", Value: "varc"},
		&Token{TokenType: "paren-right", Precedence: 100},
		&Token{TokenType: "newline", Precedence: 1000},
	}

	if !reflect.DeepEqual(out, expected) {
		t.Error("unexpected token output")
	}
}
