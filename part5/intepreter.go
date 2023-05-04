package part5

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
// expr   = term((PLUS|MINUS)term)*
// term   = factor((MUL|DIV)factor)*
// factor = INTEGER

// expr returns a float64 calculated by '+' or '-' with term.
func (it *Interpreter) expr() (result float64) {
	result = it.term()

	for it.currentToken.Type != EOF {
		switch it.currentToken.Type {
		case PLUS:
			it.eat(PLUS)
			result += it.term()
		case MINUS:
			it.eat(MINUS)
			result -= it.term()
		default:
			// just panic here
			it.eat(EOF)
		}
	}

	return
}

// term returns a float64 calculated by '*' or '/' with factor.
func (it *Interpreter) term() (result float64) {
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
			// NOTE: do not panic here, cuz it's not the outter rule, it's just a middle proc.
			return
		}
	}

	return
}

// factor returns an INTEGER.
func (it *Interpreter) factor() TokenValue {
	tk := it.currentToken
	it.eat(INTEGER)
	return tk.Value
}
