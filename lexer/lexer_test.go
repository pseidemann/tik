package lexer

import (
	"os"
	"reflect"
	"testing"
)

func TestFuncSimple(t *testing.T) {
	f, err := os.Open("../testdata/func_simple.tik")
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
		{TokenType: TypeKeyword, Precedence: 10, Value: "func"},
		{TokenType: TypeIdent, Value: "greet"},
		{TokenType: TypeParenL, Precedence: 100},
		{TokenType: TypeParenR, Precedence: 100},
		{TokenType: TypeBraceL},
		{TokenType: TypeNewline, Precedence: 1000},
		{TokenType: TypeKeyword, Precedence: 10, Value: "print"},
		{TokenType: TypeParenL, Precedence: 100},
		{TokenType: TypeString, Value: "Hello, world!"},
		{TokenType: TypeParenR, Precedence: 100},
		{TokenType: TypeNewline, Precedence: 1000},
		{TokenType: TypeBraceR},
		{TokenType: TypeNewline, Precedence: 1000},
		{TokenType: TypeNewline, Precedence: 1000},
		{TokenType: TypeIdent, Value: "greet"},
		{TokenType: TypeParenL, Precedence: 100},
		{TokenType: TypeParenR, Precedence: 100},
		{TokenType: TypeNewline, Precedence: 1000},
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
		{TokenType: TypeKeyword, Precedence: 10, Value: "print"},
		{TokenType: TypeParenL, Precedence: 100},
		{TokenType: TypeNum, Value: "10"},
		{TokenType: TypeOp, Precedence: 1, Value: "+"},
		{TokenType: TypeNum, Value: "2"},
		{TokenType: TypeOp, Precedence: 2, Value: "*"},
		{TokenType: TypeParenL, Precedence: 100},
		{TokenType: TypeNum, Value: "31"},
		{TokenType: TypeOp, Precedence: 1, Value: "-"},
		{TokenType: TypeNum, Value: "2"},
		{TokenType: TypeOp, Precedence: 2, Value: "/"},
		{TokenType: TypeNum, Value: "2"},
		{TokenType: TypeParenR, Precedence: 100},
		{TokenType: TypeParenR, Precedence: 100},
		{TokenType: TypeNewline, Precedence: 1000},
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
		{TokenType: TypeKeyword, Precedence: 10, Value: "print"},
		{TokenType: TypeParenL, Precedence: 100},
		{TokenType: TypeString, Value: "hello"},
		{TokenType: TypeParenR, Precedence: 100},
		{TokenType: TypeNewline, Precedence: 1000},
		{TokenType: TypeKeyword, Precedence: 10, Value: "print"},
		{TokenType: TypeParenL, Precedence: 100},
		{TokenType: TypeString, Value: "world1"},
		{TokenType: TypeComma},
		{TokenType: TypeString, Value: "world2"},
		{TokenType: TypeParenR, Precedence: 100},
		{TokenType: TypeNewline, Precedence: 1000},
		{TokenType: TypeKeyword, Precedence: 10, Value: "print"},
		{TokenType: TypeParenL, Precedence: 100},
		{TokenType: TypeString, Value: "world3"},
		{TokenType: TypeComma},
		{TokenType: TypeNum, Value: "1"},
		{TokenType: TypeOp, Precedence: 1, Value: "+"},
		{TokenType: TypeNum, Value: "2"},
		{TokenType: TypeOp, Precedence: 2, Value: "*"},
		{TokenType: TypeNum, Value: "3"},
		{TokenType: TypeParenR, Precedence: 100},
		{TokenType: TypeNewline, Precedence: 1000},
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
		{TokenType: TypeIdent, Value: "vara"},
		{TokenType: TypeAssign, Precedence: 100},
		{TokenType: TypeNum, Value: "1"},
		{TokenType: TypeOp, Precedence: 1, Value: "+"},
		{TokenType: TypeNum, Value: "2"},
		{TokenType: TypeNewline, Precedence: 1000},
		{TokenType: TypeIdent, Value: "varb"},
		{TokenType: TypeAssign, Precedence: 100},
		{TokenType: TypeNum, Value: "2"},
		{TokenType: TypeOp, Precedence: 2, Value: "*"},
		{TokenType: TypeNum, Value: "4"},
		{TokenType: TypeNewline, Precedence: 1000},
		{TokenType: TypeIdent, Value: "varc"},
		{TokenType: TypeAssign, Precedence: 100},
		{TokenType: TypeIdent, Value: "vara"},
		{TokenType: TypeOp, Precedence: 1, Value: "+"},
		{TokenType: TypeIdent, Value: "varb"},
		{TokenType: TypeNewline, Precedence: 1000},
		{TokenType: TypeKeyword, Precedence: 10, Value: "print"},
		{TokenType: TypeParenL, Precedence: 100},
		{TokenType: TypeIdent, Value: "vara"},
		{TokenType: TypeComma},
		{TokenType: TypeIdent, Value: "varb"},
		{TokenType: TypeComma},
		{TokenType: TypeIdent, Value: "varc"},
		{TokenType: TypeParenR, Precedence: 100},
		{TokenType: TypeNewline, Precedence: 1000},
	}

	if !reflect.DeepEqual(out, expected) {
		t.Error("unexpected token output")
	}
}
