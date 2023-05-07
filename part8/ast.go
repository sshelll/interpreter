package part8

type astNode interface {
}

type astUnaryOp struct {
	Op    *Token
	Right astNode
}

func newUnaryOp(op *Token, r astNode) *astUnaryOp {
	return &astUnaryOp{
		Op:    op,
		Right: r,
	}
}

type astBinOp struct {
	Op          *Token
	Left, Right astNode
}

func newAstBinOp(l astNode, op *Token, r astNode) *astBinOp {
	return &astBinOp{
		Left:  l,
		Op:    op,
		Right: r,
	}
}

type astNum struct {
	Token *Token
	Value TokenValue
}

func newAstNum(tk *Token) *astNum {
	if tk == nil {
		return nil
	}
	return &astNum{
		Token: tk,
		Value: tk.Value,
	}
}
