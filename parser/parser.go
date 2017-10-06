// Package parser implements parsing a sequence of tokens into an AST.
package parser

import (
	"fmt"

	"github.com/pseidemann/tik/ast"
	"github.com/pseidemann/tik/lexer"
)

var opMap = map[string]ast.OpType{
	"+": ast.OpAdd,
	"-": ast.OpSub,
	"*": ast.OpMul,
	"/": ast.OpDiv,
}

// Parser can parse tokens into an AST.
type Parser struct {
	lex     *lexer.Lexer
	prevTok *lexer.Token
}

// New creates a Parser.
func New(lex *lexer.Lexer) *Parser {
	return &Parser{
		lex: lex,
	}
}

// CreateAST generates an AST.
func (p *Parser) CreateAST() ast.Node {
	return p.parseImplicitBlock("main")
}

func (p *Parser) nextToken() (*lexer.Token, error) {
	if p.prevTok != nil {
		t := p.prevTok
		p.prevTok = nil
		return t, nil
	}
	t, err := p.lex.NextToken()
	// fmt.Println("next token", t)
	return t, err
}

func (p *Parser) unreadToken(t *lexer.Token) {
	p.prevTok = t
}

func (p *Parser) parseImplicitBlock(name string) *ast.Block {
	var stmts []ast.Node
	for {
		s := p.parseStmt()
		if s == nil {
			break
		}
		stmts = append(stmts, s)
	}
	return &ast.Block{
		Name:  name,
		Stmts: stmts,
	}
}

func (p *Parser) parseBlock(name string) *ast.Block {
	p.getToken(lexer.TypeBraceL)
	n := p.parseImplicitBlock(name)
	p.getToken(lexer.TypeBraceR)
	return n
}

func (p *Parser) parseStmt() ast.Node {
	t, err := p.nextToken()
	if err != nil {
		if err != lexer.ErrEOF {
			fmt.Println("ERROR:", err)
		}
		return nil
	}
	switch t.TokenType {
	case lexer.TypeKeyword:
		switch t.Value {
		case lexer.KWFunc:
			return p.parseFuncDef()
		case lexer.KWPrint:
			return p.parseFuncCall(t.Value)
		case lexer.KWReturn:
			expr, ok := p.parseExpr()
			if !ok {
				return &ast.Return{}
			}
			return &ast.Return{
				Value: expr,
			}
		default:
			panic("unknown keyword " + t.Value)
		}
	case lexer.TypeIdent:
		next, err := p.nextToken()
		if err != nil {
			if err != lexer.ErrEOF {
				fmt.Println("ERROR:", err)
			}
		}
		p.unreadToken(next)
		switch next.TokenType {
		case lexer.TypeParenL:
			return p.parseFuncCall(t.Value)
		case lexer.TypeAssign:
			return p.parseAssign(t.Value)
		default:
			panic(fmt.Sprintf("unexpected token %v", t))
		}
	case lexer.TypeNewline:
		return p.parseStmt()
	case lexer.TypeBraceR:
		// end of block
		p.unreadToken(t)
	default:
		panic(fmt.Sprintf("unexpected token %v", t))

	}
	return nil
}

func (p *Parser) parseFuncDef() ast.Node {
	ident := p.getToken(lexer.TypeIdent)
	p.getToken(lexer.TypeParenL)
	params := p.parseParamsList()
	p.getToken(lexer.TypeParenR)
	return &ast.FuncDef{
		Name:   ident.Value,
		Params: params,
		Body:   p.parseBlock("func"),
	}
}

func (p *Parser) parseParamsList() []*ast.Param {
	var params []*ast.Param
	for {
		t, err := p.nextToken()
		if err != nil {
			panic(err)
		}
		switch t.TokenType {
		case lexer.TypeParenR:
			p.unreadToken(t)
			return params
		case lexer.TypeIdent:
			params = append(params, &ast.Param{
				Name: t.Value,
			})
			after := p.getToken(lexer.TypeComma, lexer.TypeParenR)
			if after.TokenType == lexer.TypeParenR {
				p.unreadToken(after)
				return params
			}
		default:
			panic("expected ident, comma or paren right, got " + t.String())
		}
	}
}

func (p *Parser) parseFuncCall(name string) ast.Node {
	p.getToken(lexer.TypeParenL)
	args := p.parseExprList()
	p.getToken(lexer.TypeParenR)
	return &ast.FuncCall{
		Name: name,
		Args: args,
	}
}

func (p *Parser) parseExprList() []ast.Node {
	var exps []ast.Node
	for {
		exp, ok := p.parseExpr()
		if !ok {
			return exps
		}
		exps = append(exps, exp)
		after := p.getToken(lexer.TypeComma, lexer.TypeParenR)
		if after.TokenType == lexer.TypeParenR {
			p.unreadToken(after)
			break
		}
	}
	return exps
}

// parseExpr implements the shunting-yard algorithm.
func (p *Parser) parseExpr() (ast.Node, bool) {
	var outQueue []ast.Node
	var opStack tokenStack
	var nestingLevel int

Loop:
	for {
		t, err := p.nextToken()
		if err != nil {
			if err != lexer.ErrEOF {
				panic(err)
			}
			break
		}

		switch t.TokenType {
		case lexer.TypeComma, lexer.TypeNewline:
			p.unreadToken(t)
			break Loop
		case lexer.TypeNum:
			outQueue = append(outQueue, &ast.Number{
				Num: t.Value,
			})
		case lexer.TypeString:
			outQueue = append(outQueue, &ast.String{
				Str: t.Value,
			})
		case lexer.TypeIdent:
			next, err := p.nextToken()
			if err != nil {
				if err != lexer.ErrEOF {
					fmt.Println("ERROR:", err)
				}
			}
			p.unreadToken(next)
			switch next.TokenType {
			case lexer.TypeParenL:
				outQueue = append(outQueue, p.parseFuncCall(t.Value))
			default:
				outQueue = append(outQueue, &ast.Ident{
					Name: t.Value,
				})
			}
		case lexer.TypeOp:
			for opStack.peek() != nil &&
				opStack.peek().TokenType == lexer.TypeOp &&
				opStack.peek().Precedence >= t.Precedence {
				popped := opStack.pop()
				fmt.Println("move from stack to queue", popped)
				outQueue = queueOp(outQueue, popped)
			}
			opStack.push(t)
		case lexer.TypeParenL:
			nestingLevel++
			opStack.push(t)
		case lexer.TypeParenR:
			if nestingLevel == 0 {
				// we got the closing parenthesis from a function call
				p.unreadToken(t)
				break Loop
			}
			nestingLevel--
			for opStack.peek() != nil && opStack.peek().TokenType != lexer.TypeParenL {
				popped := opStack.pop()
				outQueue = queueOp(outQueue, popped)
			}
			// pop the left parenthesis from the stack
			popped := opStack.pop()
			if popped == nil || popped.TokenType != lexer.TypeParenL {
				panic("unbalanced parenthesis")
			}
		default:
			panic(fmt.Sprintf("unexpected token %v", t))
		}
	}

	for opStack.peek() != nil {
		popped := opStack.pop()
		outQueue = queueOp(outQueue, popped)
	}

	if len(outQueue) == 0 {
		return nil, false
	}

	return outQueue[0], true
}

func queueOp(queue []ast.Node, op *lexer.Token) []ast.Node {
	l := len(queue)
	var left, right ast.Node

	if l > 0 {
		right, queue = queue[l-1], queue[:l-1]
		l = len(queue)
	}

	if l > 0 {
		left, queue = queue[l-1], queue[:l-1]
	}

	return append(queue, &ast.Operation{
		OpType: opMap[op.Value],
		Left:   left,
		Right:  right,
	})
}

func (p *Parser) parseAssign(variable string) ast.Node {
	p.getToken(lexer.TypeAssign)
	exp, ok := p.parseExpr()
	if !ok {
		panic("expected expression on right side of assignment")
	}
	return &ast.Assign{
		Left: &ast.Ident{
			Name: variable,
		},
		Right: exp,
	}
}

func (p *Parser) getToken(tokenTypes ...lexer.TokenType) *lexer.Token {
	t, err := p.nextToken()
	if err != nil {
		panic(err)
	}
	for _, expected := range tokenTypes {
		if t.TokenType == expected {
			return t
		}
	}
	panic(fmt.Sprintf("expected token %v got %s (%s)", tokenTypes, t.TokenType, t.Value))
}
