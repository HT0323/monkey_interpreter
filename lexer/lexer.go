package lexer

import "github.com/HT0323/monkey_interpreter/token"

// 字句解析器(レキサー)の構造体
type Lexer struct {
	input        string // 入力値
	position     int    // 現在の文字を指し示す
	readPosition int    // 次の文字を指し示す
	ch           byte   // 現在精査中の文字
}

// Lexer構造体の初期化
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// 次の文字を取得
func (l *Lexer) readChar() {
	//入力が終端に達しているかを確認
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// 終端に達していなければ次の文字を取得
		l.ch = l.input[l.readPosition]
	}
	// カウンターをインクリメント
	l.position = l.readPosition
	l.readPosition += 1
}

// 現在精査中の文字に対応したトークンを生成
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

// Token構造体の初期化
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
