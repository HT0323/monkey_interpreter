package lexer

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
	return l
}
