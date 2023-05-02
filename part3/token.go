package part3

import (
	"fmt"
	"strconv"
)

const (
	EOF TokenType = iota
	INTEGER
	PLUS
	MINUS
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

type Token struct {
	Type  TokenType
	Value TokenValue
}

func (tk *Token) String() string {
	return fmt.Sprintf("Token(%v, %s)", tk.Type, tk.Value)
}
