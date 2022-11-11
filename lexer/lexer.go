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
