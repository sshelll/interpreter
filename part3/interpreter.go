package part3

import (
	"bytes"
	"fmt"
)

type Interpreter struct {
	text         string
	pos          int
	currentToken *Token
	currentChar  byte
}

func NewInterpreter(text string) *Interpreter {
	return &Interpreter{
		text:         text,
		pos:          0,
		currentToken: nil,
		currentChar:  text[0],
	}
}

func (it *Interpreter) GetNextToken() *Token {
	for it.currentChar != 0 {
		// skip white spaces
		if it.currentChar == ' ' {
			it.advance()
			continue
		}

		// plus
		if it.currentChar == '+' {
			it.advance()
			return &Token{Type: PLUS, Value: "+"}
		}

		// minus
		if it.currentChar == '-' {
			it.advance()
			return &Token{Type: MINUS, Value: "-"}
		}

		// integer
		if it.isDigit() {
			return &Token{Type: INTEGER, Value: it.integer()}
		}

		// unknown
		return &Token{Type: UNKNOWN, Value: TokenValue(it.currentChar)}
	}

	return &Token{Type: EOF, Value: ""}
}

func (it *Interpreter) Expr() (result int) {
	it.currentToken = it.GetNextToken()

	result = it.term().Int()

	for it.currentToken.Type != EOF {
		switch it.currentToken.Type {
		case PLUS:
			it.eat(PLUS)
			result += it.term().Int()
		case MINUS:
			it.eat(MINUS)
			result -= it.term().Int()
		default:
			// just panic here
			it.eat(EOF)
		}
	}

	return
}

// term is always a number in this case.
func (it *Interpreter) term() TokenValue {
	token := it.currentToken
	it.eat(INTEGER)
	return token.Value
}

func (it *Interpreter) eat(tokenType TokenType) {
	if it.currentToken.Type == tokenType {
		it.currentToken = it.GetNextToken()
		return
	}
	err := fmt.Sprintf("Error parsing input: %s, expected token type: %v, actual: %v",
		it.text, tokenType, it.currentToken)
	panic(err)
}

func (it *Interpreter) advance() {
	it.pos++
	if it.pos > len(it.text)-1 {
		it.currentChar = 0
	} else {
		it.currentChar = it.text[it.pos]
	}
}

func (it *Interpreter) integer() TokenValue {
	buf := &bytes.Buffer{}
	for it.isDigit() {
		buf.WriteByte(it.currentChar)
		it.advance()
	}
	return TokenValue(buf.String())
}

func (it *Interpreter) isDigit() bool {
	return it.currentChar >= '0' && it.currentChar <= '9'
}
