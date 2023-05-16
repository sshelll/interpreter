package main

import (
	"fmt"
	"log"

	"github.com/sshelll/interpreter/part1"
	"github.com/sshelll/interpreter/part2"
	"github.com/sshelll/interpreter/part3"
	"github.com/sshelll/interpreter/part4"
	"github.com/sshelll/interpreter/part5"
	"github.com/sshelll/interpreter/part6"
	"github.com/sshelll/interpreter/part7"
	"github.com/sshelll/interpreter/part8"
	"github.com/sshelll/interpreter/part9"
	"github.com/sshelll/menuscreen"
)

func selectPart() (fn func(string) interface{}, lexerFn func(string)) {
	// register parts
	type part struct {
		fn      func(string) interface{}
		lexerFn func(string)
		desc    string
	}

	parts := []*part{
		{part1Fn, nil, "part 1: simple calculator, only support addition"},
		{part2Fn, nil, "part 2: support 'minus', 'spaces' and 'long numbers' based on part 1"},
		{part3Fn, nil, "part 3: support 'multi-ops' based on part 2"},
		{part4Fn, nil, "part 4: refactor part3 and replace the ops into '*' and '/'"},
		{part5Fn, nil, "part 5: support 'quadratic operations' based on part 4"},
		{part6Fn, nil, "part 6: support 'parentheses' based on part 5"},
		{part7Fn, nil, "part 7: refactor part6 with ast"},
		{part8Fn, nil, "part 8: support 'unary op' based on part7"},
		{part9Fn, part9LexerFn, "part 9: simple pascal interpreter"},
	}

	if *p > 0 {
		pt := parts[*p-1]
		return pt.fn, pt.lexerFn
	}

	// init menu
	menu, err := menuscreen.NewMenuScreen()
	if err != nil {
		log.Fatalln(err)
	}
	defer menu.Fini()

	menu.SetTitle("parts")
	for _, p := range parts {
		menu.AppendLines(p.desc)
	}
	menu.Start()

	// get chosen part
	idx, _, ok := menu.ChosenLine()
	if !ok {
		log.Fatalln("you haven't chosen any lines")
	}

	p := parts[idx]
	return p.fn, p.lexerFn
}

func part1Fn(expr string) interface{} {
	return part1.NewInterpreter(expr).Expr()
}

func part2Fn(expr string) interface{} {
	return part2.NewInterpreter(expr).Expr()
}

func part3Fn(expr string) interface{} {
	return part3.NewInterpreter(expr).Expr()
}

func part4Fn(expr string) interface{} {
	return part4.NewInterpreter(expr).Expr()
}

func part5Fn(expr string) interface{} {
	return part5.NewInterpreter(expr).Expr()
}

func part6Fn(expr string) interface{} {
	return part6.NewInterpreter(expr).Expr()
}

func part7Fn(expr string) interface{} {
	return part7.NewInterpreter(expr).Interpret()
}

func part8Fn(expr string) interface{} {
	return part8.NewInterpreter(expr).Interpret()
}

func part9Fn(expr string) interface{} {
	intr := part9.NewInterpreter(expr)
	intr.Interpret()
	return intr.GlobalScope()
}

func part9LexerFn(expr string) {
	lx := part9.NewLexer(expr)
	for token := lx.GetNextToken(); token != nil; token = lx.GetNextToken() {
		fmt.Printf("%v\n", token)
		if token.Type == part9.EOF {
			break
		}
	}
}
