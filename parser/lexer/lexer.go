// Package lexer implements parsing source code into tokens.
package lexer

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
)

// ErrEOF is the error which is returned when all tokens are consumed.
var ErrEOF = errors.New("no more tokens")

// Lexer can parse source code into a sequence of tokens.
type Lexer struct {
	buf *bufio.Reader
}

// New creates a Lexer.
func New(rd io.Reader) *Lexer {
	return &Lexer{
		buf: bufio.NewReader(rd),
	}
}

// NextToken returns the next token available.
func (l *Lexer) NextToken() (*Token, error) {
	r, _, err := l.buf.ReadRune()
	if err != nil {
		if err == io.EOF {
			// wrap with our own error to encapsulate implementation details
			return nil, ErrEOF
		}
		return nil, err
	}

	if isDigit(r) {
		l.unreadRune()
		num, err := l.readWhile(isDigit)
		if err != nil {
			return nil, err
		}
		return &Token{TokenType: TypeNum, Value: num}, nil
	} else if isOp(r) {
		return &Token{TokenType: TypeOp, Value: string(r), Precedence: opPrecedence[r]}, nil
	} else if isIdent(r) {
		l.unreadRune()
		ident, err := l.readWhile(isIdent)
		if err != nil {
			return nil, err
		}
		if isKeyword(ident) {
			return &Token{TokenType: TypeKeyword, Precedence: 10, Value: ident}, nil
		}
		return &Token{TokenType: TypeIdent, Value: ident}, nil
	} else if r == '(' {
		if err != nil {
			return nil, err
		}
		return &Token{TokenType: TypeParenL, Precedence: 100}, nil
	} else if r == ')' {
		if err != nil {
			return nil, err
		}
		return &Token{TokenType: TypeParenR, Precedence: 100}, nil
	} else if r == '{' {
		if err != nil {
			return nil, err
		}
		return &Token{TokenType: TypeBraceL}, nil
	} else if r == '}' {
		if err != nil {
			return nil, err
		}
		return &Token{TokenType: TypeBraceR}, nil
	} else if r == '\n' {
		if err != nil {
			return nil, err
		}
		return &Token{TokenType: TypeNewline, Precedence: 1000}, nil
	} else if r == ',' {
		if err != nil {
			return nil, err
		}
		return &Token{TokenType: TypeComma}, nil
	} else if r == '"' {
		if err != nil {
			return nil, err
		}
		str, err := l.readWhile(insideString)
		if err != nil {
			return nil, err
		}
		_, _, err = l.buf.ReadRune() // discard closing "
		if err != nil {
			return nil, err
		}
		return &Token{TokenType: TypeString, Value: str}, nil
	} else if r == '=' {
		return &Token{TokenType: TypeAssign, Precedence: 100}, nil
	} else if isWhitespace(r) {
		_, err := l.readWhile(isWhitespace)
		if err != nil {
			return nil, err
		}
		return l.NextToken()
	}

	return nil, fmt.Errorf("invalid rune found %#v", string(r))
}

func (l *Lexer) unreadRune() {
	err := l.buf.UnreadRune()
	if err != nil {
		// should never happen
		panic(err)
	}
}

func (l *Lexer) readWhile(predicate func(rune) bool) (string, error) {
	var b bytes.Buffer

	for {
		r, _, err := l.buf.ReadRune()
		if err != nil {
			return "", err
		}
		if !predicate(r) {
			l.unreadRune()
			break
		}
		// fmt.Println("write rune", string(r))
		b.WriteRune(r)
	}
	return b.String(), nil
}
