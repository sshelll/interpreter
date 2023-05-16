package part9

import (
	"fmt"
	"strconv"
)

var RESERVED_KEYWORDS = map[string]*Token{
	"BEGIN": {Type: BEGIN, Value: "BEGIN"},
	"END":   {Type: END, Value: "END"},
}

const (
	EOF     TokenType = iota
	INTEGER           // num
	PLUS              // +
	MINUS             // -
	MUL               // *
	DIV               // /
	LPAREN            // (
	RPAREN            // )
	BEGIN             // BEGIN
	END               // END
	DOT               // .
	ASSIGN            // =
	SEMI              // ;
	ID                // start with char, contains chars and numbers, such as 'a1'
	UNKNOWN
)

type TokenType int

// String
func (tt TokenType) String() string {
	switch tt {
	case EOF:
		return "EOF"
	case INTEGER:
		return "INTEGER"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case MUL:
		return "MUL"
	case DIV:
		return "DIV"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case BEGIN:
		return "BEGIN"
	case END:
		return "END"
	case DOT:
		return "DOT"
	case ASSIGN:
		return "ASSIGN"
	case SEMI:
		return "SEMI"
	case ID:
		return "ID"
	default:
		return "UNKNOWN"
	}
}

type TokenValue string

type Token struct {
	Type  TokenType
	Value TokenValue
}

func (tv TokenValue) String() string {
	return string(tv)
}

func (tv TokenValue) Int() int {
	i, _ := strconv.Atoi(string(tv))
	return i
}

func (tv TokenValue) Float() float64 {
	// in this case, numbers can only be integer.
	return float64(tv.Int())
}

func (tk *Token) String() string {
	return fmt.Sprintf("Token(%v, %s)", tk.Type, tk.Value)
}
