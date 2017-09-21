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
		&Token{TokenType: TypeIdent, Value: "outer"},
		&Token{TokenType: TypeAssign, Precedence: 100},
		&Token{TokenType: TypeNum, Value: "1"},
		&Token{TokenType: TypeNewline, Precedence: 1000},
		&Token{TokenType: TypeNewline, Precedence: 1000},
		&Token{TokenType: TypeKeyword, Precedence: 10, Value: "func"},
		&Token{TokenType: TypeIdent, Value: "greet"},
		&Token{TokenType: TypeParenL, Precedence: 100},
		&Token{TokenType: TypeIdent, Value: "a"},
		&Token{TokenType: TypeComma},
		&Token{TokenType: TypeIdent, Value: "b"},
		&Token{TokenType: TypeParenR, Precedence: 100},
		&Token{TokenType: TypeBraceL},
		&Token{TokenType: TypeNewline, Precedence: 1000},
		&Token{TokenType: TypeIdent, Value: "inner"},
		&Token{TokenType: TypeAssign, Precedence: 100},
		&Token{TokenType: TypeNum, Value: "4"},
		&Token{TokenType: TypeNewline, Precedence: 1000},
		&Token{TokenType: TypeKeyword, Precedence: 10, Value: "print"},
		&Token{TokenType: TypeParenL, Precedence: 100},
		&Token{TokenType: TypeIdent, Value: "outer"},
		&Token{TokenType: TypeComma},
		&Token{TokenType: TypeIdent, Value: "a"},
		&Token{TokenType: TypeComma},
		&Token{TokenType: TypeIdent, Value: "b"},
		&Token{TokenType: TypeComma},
		&Token{TokenType: TypeIdent, Value: "inner"},
		&Token{TokenType: TypeParenR, Precedence: 100},
		&Token{TokenType: TypeNewline, Precedence: 1000},
		&Token{TokenType: TypeBraceR},
		&Token{TokenType: TypeNewline, Precedence: 1000},
		&Token{TokenType: TypeNewline, Precedence: 1000},
		&Token{TokenType: TypeIdent, Value: "greet"},
		&Token{TokenType: TypeParenL, Precedence: 100},
		&Token{TokenType: TypeNum, Value: "2"},
		&Token{TokenType: TypeComma},
		&Token{TokenType: TypeNum, Value: "3"},
		&Token{TokenType: TypeParenR, Precedence: 100},
		&Token{TokenType: TypeNewline, Precedence: 1000},
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
		&Token{TokenType: TypeKeyword, Precedence: 10, Value: "print"},
		&Token{TokenType: TypeParenL, Precedence: 100},
		&Token{TokenType: TypeNum, Value: "10"},
		&Token{TokenType: TypeOp, Precedence: 1, Value: "+"},
		&Token{TokenType: TypeNum, Value: "2"},
		&Token{TokenType: TypeOp, Precedence: 2, Value: "*"},
		&Token{TokenType: TypeParenL, Precedence: 100},
		&Token{TokenType: TypeNum, Value: "31"},
		&Token{TokenType: TypeOp, Precedence: 1, Value: "-"},
		&Token{TokenType: TypeNum, Value: "2"},
		&Token{TokenType: TypeOp, Precedence: 2, Value: "/"},
		&Token{TokenType: TypeNum, Value: "2"},
		&Token{TokenType: TypeParenR, Precedence: 100},
		&Token{TokenType: TypeParenR, Precedence: 100},
		&Token{TokenType: TypeNewline, Precedence: 1000},
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
		&Token{TokenType: TypeKeyword, Precedence: 10, Value: "print"},
		&Token{TokenType: TypeParenL, Precedence: 100},
		&Token{TokenType: TypeString, Value: "hello"},
		&Token{TokenType: TypeParenR, Precedence: 100},
		&Token{TokenType: TypeNewline, Precedence: 1000},
		&Token{TokenType: TypeKeyword, Precedence: 10, Value: "print"},
		&Token{TokenType: TypeParenL, Precedence: 100},
		&Token{TokenType: TypeString, Value: "world1"},
		&Token{TokenType: TypeComma},
		&Token{TokenType: TypeString, Value: "world2"},
		&Token{TokenType: TypeParenR, Precedence: 100},
		&Token{TokenType: TypeNewline, Precedence: 1000},
		&Token{TokenType: TypeKeyword, Precedence: 10, Value: "print"},
		&Token{TokenType: TypeParenL, Precedence: 100},
		&Token{TokenType: TypeString, Value: "world3"},
		&Token{TokenType: TypeComma},
		&Token{TokenType: TypeNum, Value: "1"},
		&Token{TokenType: TypeOp, Precedence: 1, Value: "+"},
		&Token{TokenType: TypeNum, Value: "2"},
		&Token{TokenType: TypeOp, Precedence: 2, Value: "*"},
		&Token{TokenType: TypeNum, Value: "3"},
		&Token{TokenType: TypeParenR, Precedence: 100},
		&Token{TokenType: TypeNewline, Precedence: 1000},
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
		&Token{TokenType: TypeIdent, Value: "vara"},
		&Token{TokenType: TypeAssign, Precedence: 100},
		&Token{TokenType: TypeNum, Value: "1"},
		&Token{TokenType: TypeOp, Precedence: 1, Value: "+"},
		&Token{TokenType: TypeNum, Value: "2"},
		&Token{TokenType: TypeNewline, Precedence: 1000},
		&Token{TokenType: TypeIdent, Value: "varb"},
		&Token{TokenType: TypeAssign, Precedence: 100},
		&Token{TokenType: TypeNum, Value: "2"},
		&Token{TokenType: TypeOp, Precedence: 2, Value: "*"},
		&Token{TokenType: TypeNum, Value: "4"},
		&Token{TokenType: TypeNewline, Precedence: 1000},
		&Token{TokenType: TypeIdent, Value: "varc"},
		&Token{TokenType: TypeAssign, Precedence: 100},
		&Token{TokenType: TypeIdent, Value: "vara"},
		&Token{TokenType: TypeOp, Precedence: 1, Value: "+"},
		&Token{TokenType: TypeIdent, Value: "varb"},
		&Token{TokenType: TypeNewline, Precedence: 1000},
		&Token{TokenType: TypeKeyword, Precedence: 10, Value: "print"},
		&Token{TokenType: TypeParenL, Precedence: 100},
		&Token{TokenType: TypeIdent, Value: "vara"},
		&Token{TokenType: TypeComma},
		&Token{TokenType: TypeIdent, Value: "varb"},
		&Token{TokenType: TypeComma},
		&Token{TokenType: TypeIdent, Value: "varc"},
		&Token{TokenType: TypeParenR, Precedence: 100},
		&Token{TokenType: TypeNewline, Precedence: 1000},
	}

	if !reflect.DeepEqual(out, expected) {
		t.Error("unexpected token output")
	}
}
