package lexer

import "fmt"

// TokenType declares the token type.
type TokenType int

// All available token types.
const (
	TypeAssign TokenType = iota
	TypeIdent
	TypeKeyword
	TypeComma
	TypeNewline
	TypeNum
	TypeOp
	TypeParenL // (
	TypeParenR // )
	TypeBraceL // {
	TypeBraceR // }
	TypeString
)

var types = [...]string{
	"assignment",
	"identifier",
	"keyword",
	"comma",
	"newline",
	"number",
	"operator",
	"paren-left",
	"paren-right",
	"brace-left",
	"brace-right",
	"string",
}

// Token is a categorized lexeme.
type Token struct {
	TokenType  TokenType
	Precedence int
	Value      string
}

func (t *Token) String() string {
	return fmt.Sprintf("(%s<%d> %#v)", types[t.TokenType], t.Precedence, t.Value)
}
