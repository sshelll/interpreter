package part9

import (
	"bytes"
)

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
		if lx.isWhiteSpaces() {
			lx.skipWhiteSpaces()
			continue
		}

		// id
		if lx.isAlaph() {
			return lx.identifier()
		}

		// assign
		if lx.currentChar == ':' && lx.peek() == '=' {
			lx.advance()
			lx.advance()
			return &Token{Type: ASSIGN, Value: ":="}
		}

		// semi
		if lx.currentChar == ';' {
			lx.advance()
			return &Token{Type: SEMI, Value: ";"}
		}

		// dot
		if lx.currentChar == '.' {
			lx.advance()
			return &Token{Type: DOT, Value: "."}
		}

		// plus
		if lx.currentChar == '+' {
			lx.advance()
			return &Token{Type: PLUS, Value: "+"}
		}

		// minus
		if lx.currentChar == '-' {
			lx.advance()
			return &Token{Type: MINUS, Value: "-"}
		}

		// mul
		if lx.currentChar == '*' {
			lx.advance()
			return &Token{Type: MUL, Value: "*"}
		}

		// div
		if lx.currentChar == '/' {
			lx.advance()
			return &Token{Type: DIV, Value: "/"}
		}

		// integer
		if lx.isDigit() {
			return &Token{Type: INTEGER, Value: lx.integer()}
		}

		// lparen
		if lx.currentChar == '(' {
			lx.advance()
			return &Token{Type: LPAREN, Value: "("}
		}

		// rparen
		if lx.currentChar == ')' {
			lx.advance()
			return &Token{Type: RPAREN, Value: ")"}
		}

		// unknown
		return &Token{Type: UNKNOWN, Value: TokenValue(lx.currentChar)}
	}

	// eof
	return &Token{Type: EOF, Value: ""}
}

// peek get the next char without modifying pos.
func (lx *Lexer) peek() byte {
	if peekPos := lx.pos + 1; peekPos > len(lx.text)-1 {
		return 0
	} else {
		return lx.text[peekPos]
	}
}

// advance move the pos to the next and update currentChar.
func (lx *Lexer) advance() {
	lx.pos++
	if lx.pos > len(lx.text)-1 {
		lx.currentChar = 0
	} else {
		lx.currentChar = lx.text[lx.pos]
	}
}

// integer read an num from text.
func (lx *Lexer) integer() TokenValue {
	buf := &bytes.Buffer{}
	for lx.isDigit() {
		buf.WriteByte(lx.currentChar)
		lx.advance()
	}
	return TokenValue(buf.String())
}

// identifier returns reserved keyword or normal id.
func (lx *Lexer) identifier() *Token {
	buf := &bytes.Buffer{}
	for lx.currentChar != 0 && lx.isAlNum() {
		buf.WriteByte(lx.currentChar)
		lx.advance()
	}
	id := buf.String()
	// keyword
	if token, ok := RESERVED_KEYWORDS[id]; ok {
		return token
	}
	// normal id
	return &Token{Type: ID, Value: TokenValue(id)}
}

// skipWhiteSpaces skip all white spaces by moving pos.
func (lx *Lexer) skipWhiteSpaces() {
	for lx.isWhiteSpaces() {
		lx.advance()
	}
}

func (lx *Lexer) isWhiteSpaces() bool {
	return lx.currentChar == ' ' ||
		lx.currentChar == '\n' ||
		lx.currentChar == '\t'
}

func (lx *Lexer) isDigit() bool {
	return lx.currentChar >= '0' && lx.currentChar <= '9'
}

func (lx *Lexer) isAlaph() bool {
	return (lx.currentChar >= 'A' && lx.currentChar <= 'Z') ||
		(lx.currentChar >= 'a' && lx.currentChar <= 'z')
}

func (lx *Lexer) isAlNum() bool {
	return lx.isDigit() || lx.isAlaph()
}
