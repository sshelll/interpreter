package partone

import "fmt"

type Interpreter struct {
	text         string
	pos          int
	currentToken *Token
}

func NewInterpreter(text string) *Interpreter {
	return &Interpreter{text: text}
}

// GetNextToken Lexical Analyzer (also known as scanner or tokenizer)
func (it *Interpreter) GetNextToken() *Token {

	text := it.text

	// end
	if it.pos > len(text)-1 {
		return &Token{Type: EOF, Value: ""}
	}

	curCh := text[it.pos]

	// num
	if curCh >= '0' && curCh <= '9' {
		it.pos++
		return &Token{Type: INTEGER, Value: TokenValue(curCh)}
	}

	// plus
	if curCh == '+' {
		it.pos++
		return &Token{Type: PLUS, Value: TokenValue(curCh)}
	}

	// unknown
	err := fmt.Sprintf("Error parsing input: %s, pos: %d, ch: %c", text, it.pos, curCh)
	panic(err)

}

// Expr parser / interpreter, expr -> INTEGER PLUS INTEGER
func (it *Interpreter) Expr() int {

	it.currentToken = it.GetNextToken()

	left := it.currentToken
	it.eat(INTEGER)

	_ = it.currentToken
	it.eat(PLUS)

	right := it.currentToken
	it.eat(INTEGER)

	_ = it.currentToken
	it.eat(EOF)

	leftVal := left.Value
	rightVal := right.Value

	return leftVal.Int() + rightVal.Int()

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
