package lexer

import (
	"github.com/HT0323/monkey_interpreter/token"
)

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
	l.skipWhitespace()

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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// Token構造体の初期化
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 識別子として取得する(非英字が出て来るまで文字を読み進める)
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 与えられた文字列が識別子(英字）なのかを判定する
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'z' || ch == '_'
}

// 読み込む文字が空欄、改行かを判定し該当すれば次の文字に読み進める
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// 数字として取得する(非数字が出て来るまで文字を読み進める)
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 与えられた文字列が数字なのかを判定する
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
