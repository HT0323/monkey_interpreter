package parser

import (
	"github.com/HT0323/monkey_interpreter/ast"
	"github.com/HT0323/monkey_interpreter/lexer"
	"github.com/HT0323/monkey_interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// Parser構造体を作成する
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

// curToken,peekTokenにtokenを設定する
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
