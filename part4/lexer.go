package part4

import "bytes"

type Lexer struct {
	text        string
	pos         int
	currentChar byte
}

func NewLexer(text string) *Lexer {
	return &Lexer{
		text:        text,
		pos:         0,
		currentChar: text[0],
	}
}

func (lx *Lexer) GetNextToken() *Token {
	for lx.currentChar != 0 {
		// skip white spaces
		if lx.currentChar == ' ' {
			lx.skipWhiteSpaces()
			continue
		}

		// plus
		if lx.currentChar == '*' {
			lx.advance()
			return &Token{Type: MUL, Value: "*"}
		}

		// minus
		if lx.currentChar == '/' {
			lx.advance()
			return &Token{Type: DIV, Value: "/"}
		}

		// integer
		if lx.isDigit() {
			return &Token{Type: INTEGER, Value: lx.integer()}
		}

		// unknown
		return &Token{Type: UNKNOWN, Value: TokenValue(lx.currentChar)}
	}

	return &Token{Type: EOF, Value: ""}
}

func (lx *Lexer) advance() {
	lx.pos++
	if lx.pos > len(lx.text)-1 {
		lx.currentChar = 0
	} else {
		lx.currentChar = lx.text[lx.pos]
	}
}

func (lx *Lexer) integer() TokenValue {
	buf := &bytes.Buffer{}
	for lx.isDigit() {
		buf.WriteByte(lx.currentChar)
		lx.advance()
	}
	return TokenValue(buf.String())
}

func (lx *Lexer) skipWhiteSpaces() {
	for lx.currentChar == ' ' {
		lx.advance()
	}
}

func (lx *Lexer) isDigit() bool {
	return lx.currentChar >= '0' && lx.currentChar <= '9'
}
