package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" //ユーザー定義識別子
	INT   = "INT"   //数値

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION" //言語定義識別子
	LET      = "LET"      //言語定義識別子
	TRUE     = "TRUE"     //言語定義識別子
	FALSE    = "FALSE"    //言語定義識別子
	IF       = "IF"       //言語定義識別子
	ELSE     = "ELSE"     //言語定義識別子
	RETURN   = "RETURN"   //言語定義識別子

	EQ     = "=="
	NOT_EQ = "!="

	STRING = "STRING" //文字列

	LBRACKET = "["
	RBRACKET = "]"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// 渡された識別子がユーザ定義かキーワード(言語定義)識別子かを判定
func LookupIdent(ident string) TokenType {
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
