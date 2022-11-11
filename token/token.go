package token

import "fmt"

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" //ユーザー定義識別子
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION" //言語定義識別子
	LET      = "LET"      //言語定義識別子
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// 渡された識別子がユーザ定義かキーワード(言語定義)識別子かを判定
func LookupIdent(ident string) TokenType {
	fmt.Println(ident)
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
