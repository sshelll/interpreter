package part9

import "fmt"

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
	node := ps.program()
	ps.eat(EOF)
	return node
}

// Grammars:
// program              = compound_statement DOT
// compound_statement   = BEGIN statement_list END
// statement_list       = statement | statement SEMI statement_list
// statment             = compound_statement | assignment_statement | empty
// assignment_statement = variable ASSIGN expr
// empty                =
// expr                 = term ((PLUS | MINUS) term)*
// term                 = factor ((MUL | DIV) factor)*
// factor               = PLUS factor | MINUS factor | INTEGER | LPAREN expr RPAREN | variable
// variable             = ID

func (ps *Parser) program() astNode {
	node := ps.compoundStatement()
	ps.eat(DOT)
	return node
}

func (ps *Parser) compoundStatement() astNode {
	ps.eat(BEGIN)
	nodeList := ps.statementList()
	ps.eat(END)
	return newAstCompound(nodeList)
}

func (ps *Parser) statementList() []astNode {
	stmt := ps.statement()

	result := []astNode{stmt}

	for ps.currentToken.Type == SEMI {
		ps.eat(SEMI)
		result = append(result, ps.statement())
	}

	if ps.currentToken.Type == ID {
		panic("unexpected token_type ID after statement_list")
	}

	return result
}

func (ps *Parser) statement() astNode {
	var node astNode

	switch ps.currentToken.Type {
	case BEGIN:
		node = ps.compoundStatement()
	case ID:
		node = ps.assignStatement()
	default:
		node = ps.empty()
	}

	return node
}

func (ps *Parser) assignStatement() astNode {
	left := ps.variable()
	tk := ps.currentToken
	ps.eat(ASSIGN)
	right := ps.expr()
	return newAstAssign(left, tk, right)
}

func (ps *Parser) empty() astNode {
	return newAstNoOp()
}

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
	switch tk := ps.currentToken; tk.Type {
	case PLUS:
		ps.eat(PLUS)
		return newUnaryOp(tk, ps.factor())
	case MINUS:
		ps.eat(MINUS)
		return newUnaryOp(tk, ps.factor())
	case INTEGER:
		ps.eat(INTEGER)
		return newAstNum(tk)
	case LPAREN:
		ps.eat(LPAREN)
		node := ps.expr()
		ps.eat(RPAREN)
		return node
	case ID:
		node := ps.variable()
		return node
	default:
		return nil
	}
}

func (ps *Parser) variable() astNode {
	node := newAstVar(ps.currentToken)
	ps.eat(ID)
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
