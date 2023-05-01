package partone

import "fmt"

const (
	EOF TokenType = iota
	INTEGER
	PLUS
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
	default:
		return "UNKNOWN"
	}
}

type TokenValue string

func (tv TokenValue) String() string {
	return string(tv)
}

func (tv TokenValue) Int() int {
	// in this part, we only have single digit integers
	return int(tv[0] - '0')
}

type Token struct {
	Type  TokenType
	Value TokenValue
}

func (tk *Token) String() string {
	return fmt.Sprintf("Token(%v, %s)", tk.Type, tk.Value)
}
