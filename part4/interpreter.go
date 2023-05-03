package part4

import "fmt"

type Interpreter struct {
	lx           *Lexer
	currentToken *Token
}

func NewInterpreter(text string) *Interpreter {
	lx := NewLexer(text)
	return &Interpreter{
		lx:           lx,
		currentToken: lx.GetNextToken(),
	}
}

func (it *Interpreter) Expr() (result float64) {
	return it.expr()
}

func (it *Interpreter) eat(tt TokenType) {
	if it.currentToken.Type == tt {
		it.currentToken = it.lx.GetNextToken()
		return
	}
	err := fmt.Sprintf("Error parsing input: %s, expected token type: %v, actual: %v",
		it.lx.text, tt, it.currentToken)
	panic(err)
}

// Grammars:
// factor = INTEGER
// expr = factor((MUL|DIV)factor)*

// factor returns a INTEGER.
func (it *Interpreter) factor() TokenValue {
	tk := it.currentToken
	it.eat(INTEGER)
	return tk.Value
}

// expr returns a calculated INTEGER.
func (it *Interpreter) expr() (result float64) {
	result = it.factor().Float()

	for it.currentToken.Type != EOF {
		switch it.currentToken.Type {
		case MUL:
			it.eat(MUL)
			result *= it.factor().Float()
		case DIV:
			it.eat(DIV)
			result /= it.factor().Float()
		default:
			// just panic here
			it.eat(EOF)
		}
	}

	return
}
