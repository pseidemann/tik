package lexer

import "regexp"

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

var opMap = map[rune]bool{
	'+': true,
	'-': true,
	'*': true,
	'/': true,
}

func isOp(r rune) bool {
	return opMap[r]
}

var opPrecedence = map[rune]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
}

var keywords = map[string]bool{
	"func":  true,
	"print": true,
}

func isKeyword(ident string) bool {
	return keywords[ident]
}

func isIdent(r rune) bool {
	ok, _ := regexp.MatchString("[a-z]", string(r))
	return ok
}

func insideString(r rune) bool {
	return r != '"'
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t'
}
