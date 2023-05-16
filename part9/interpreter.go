package part9

import (
	"fmt"
	"reflect"
)

type Interpreter struct {
	parser  *Parser
	visitor *astVisitor
}

func NewInterpreter(text string) *Interpreter {
	return &Interpreter{
		parser:  NewParser(NewLexer(text)),
		visitor: newAstVisitor(),
	}
}

func (it *Interpreter) Interpret() float64 {
	tree := it.parser.Parse()
	return it.visitor.visit(tree)
}

func (it *Interpreter) GlobalScope() map[string]float64 {
	return it.visitor.GlobalScope
}

type astVisitor struct {
	// GlobalScope is a symbol table, in this part we use map[string]float64 as impl,
	// cuz all the variables can be represented by float64 for now(we do not consider overflow situation).
	GlobalScope map[string]float64
}

func newAstVisitor() *astVisitor {
	return &astVisitor{
		GlobalScope: make(map[string]float64),
	}
}

func (visitor *astVisitor) visit(node astNode) (result float64) {
	switch node.(type) {
	case *astBinOp:
		return visitor.visitBinOp(node)
	case *astNum:
		return visitor.visitNum(node)
	case *astUnaryOp:
		return visitor.visitUnaryOp(node)
	case *astCompound:
		// compound statement has no return value.
		visitor.visitCompound(node)
	case *astAssign:
		// assign statement has no return value.
		visitor.visitAssign(node)
	case *astVar:
		return visitor.visitVar(node)
	case *astNoOp:
		visitor.visitNoOp(node)
	default:
		panic(fmt.Sprintf("invalid type of astNode: %v", reflect.TypeOf(node)))
	}
	// just return any value here.
	return
}

func (visitor *astVisitor) visitBinOp(node astNode) (result float64) {
	bo := node.(*astBinOp)

	l := visitor.visit(bo.Left)
	r := visitor.visit(bo.Right)

	switch t := bo.Op.Type; t {
	case PLUS:
		return l + r
	case MINUS:
		return l - r
	case MUL:
		return l * r
	case DIV:
		return l / r
	default:
		panic(fmt.Sprintf("invalid op of binOp: %v", t))
	}
}

func (visitor *astVisitor) visitNum(node astNode) (result float64) {
	num := node.(*astNum)
	return num.Value.Float()
}

func (visitor *astVisitor) visitUnaryOp(node astNode) (result float64) {
	uo := node.(*astUnaryOp)
	r := visitor.visit(uo.Right)

	switch t := uo.Op.Type; t {
	case PLUS:
		return r
	case MINUS:
		return -r
	default:
		panic(fmt.Sprintf("invalid op of unaryOp: %v", t))
	}
}

func (visitor *astVisitor) visitCompound(node astNode) {
	cp := node.(*astCompound)
	for i := range cp.Children {
		visitor.visit(cp.Children[i])
	}
}

func (visitor *astVisitor) visitAssign(node astNode) {
	as := node.(*astAssign)
	left := as.Left.(*astVar)
	varName := left.Value.String()
	// set <var_id, val> in to the scope.
	visitor.GlobalScope[varName] = visitor.visit(as.Right)
}

func (visitor *astVisitor) visitVar(node astNode) (result float64) {
	v := node.(*astVar)
	varName := v.Value.String()
	val, ok := visitor.GlobalScope[varName]
	if ok {
		return val
	}
	err := fmt.Sprintf("variable %v not identified", varName)
	panic(err)
}

func (visitor *astVisitor) visitNoOp(node astNode) {
	return
}
