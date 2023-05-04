package part6

import (
	"fmt"
	"strconv"
)

const (
	EOF TokenType = iota
	INTEGER
	PLUS
	MINUS
	MUL
	DIV
	LPAREN
	RPAREN
	UNKNOWN
)

type TokenType int

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
	default:
		return "UNKNOWN"
	}
}

type TokenValue string

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

type Token struct {
	Type  TokenType
	Value TokenValue
}

func (tk *Token) String() string {
	return fmt.Sprintf("Token(%v, %s)", tk.Type, tk.Value)
}
