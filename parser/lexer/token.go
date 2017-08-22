package lexer

import "fmt"

// TokenType declares the token type.
type TokenType string

// All available token types.
const (
	TypeAssign  TokenType = "assignment"
	TypeIdent   TokenType = "identifier"
	TypeKeyword TokenType = "keyword"
	TypeComma   TokenType = "comma"
	TypeNewline TokenType = "newline"
	TypeNum     TokenType = "number"
	TypeOp      TokenType = "operator"
	TypeParenL  TokenType = "paren-left"  // (
	TypeParenR  TokenType = "paren-right" // )
	TypeBraceL  TokenType = "brace-left"  // {
	TypeBraceR  TokenType = "brace-right" // }
	TypeString  TokenType = "string"
)

// Token is a categorized lexeme.
type Token struct {
	TokenType  TokenType
	Precedence int
	Value      string
}

func (t *Token) String() string {
	return fmt.Sprintf("(%s<%d> %#v)", t.TokenType, t.Precedence, t.Value)
}
