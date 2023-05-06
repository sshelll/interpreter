package part7

import (
	"fmt"
)

// Parser is a tool which builds an 'IR'.
// 'IR' stands for intermediate representation, and the data struct is AST.
type Parser struct {
	lx           *Lexer
	currentToken *Token
}

func NewParser(lexer *Lexer) *Parser {
	if lexer == nil {
		panic("lexer is nil")
	}
	return &Parser{
		lx:           lexer,
		currentToken: lexer.GetNextToken(),
	}
}

func (ps *Parser) Parse() astNode {
	return ps.expr()
}

// Grammars:
// expr   = term((PLUS|MINUS)term)*
// term   = factor((MUL|DIV)factor)*
// factor = INTEGER|(LPAREN expr RPAREN)

func (ps *Parser) expr() astNode {
	node := ps.term()

	for ps.currentToken.Type != EOF {
		tk := ps.currentToken
		if tk.Type == PLUS {
			ps.eat(PLUS)
		} else if tk.Type == MINUS {
			ps.eat(MINUS)
		} else {
			break
		}
		node = newAstBinOp(node, tk, ps.term())
	}

	return node
}

func (ps *Parser) term() astNode {
	node := ps.factor()

	for ps.currentToken.Type != EOF {
		tk := ps.currentToken
		if tk.Type == MUL {
			ps.eat(MUL)
		} else if tk.Type == DIV {
			ps.eat(DIV)
		} else {
			break
		}
		node = newAstBinOp(node, tk, ps.factor())
	}

	return node
}

func (ps *Parser) factor() astNode {
	tk := ps.currentToken

	// INTEGER
	if tk.Type == INTEGER {
		ps.eat(INTEGER)
		return newAstNum(tk)
	}

	// LPAREN expr RPAREN
	ps.eat(LPAREN)
	node := ps.expr()
	ps.eat(RPAREN)
	return node
}

func (ps *Parser) eat(tt TokenType) {
	if ps.currentToken.Type == tt {
		ps.currentToken = ps.lx.GetNextToken()
		return
	}
	err := fmt.Sprintf("Error parsing input: %s, expected token type: %v, actual: %v",
		ps.lx.text, tt, ps.currentToken)
	panic(err)
}
