package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sshelll/interpreter/part1"
	"github.com/sshelll/interpreter/part2"
	"github.com/sshelll/interpreter/part3"
	"github.com/sshelll/interpreter/part4"
	"github.com/sshelll/interpreter/part5"
	"github.com/sshelll/menuscreen"
)

var (
	expr        = flag.String("e", "", "expression to evaluate")
	interactive = flag.Bool("i", false, "interactive mode")
)

func main() {
	flag.Parse()

	if strings.TrimSpace(*expr) == "" && !*interactive {
		flag.Usage()
		return
	}

	// choose a part, and wrap the func
	fn := selectPart()
	safeFn := func(expr string) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		fmt.Println(fn(expr))
	}

	// expr mode
	if strings.TrimSpace(*expr) != "" {
		safeFn(*expr)
	}

	// interactive mode
	if *interactive {
		interactiveMode(safeFn)
	}
}

func interactiveMode(fn func(string)) {
	reader := bufio.NewReader(os.Stdin)
	for {
		print(">> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		// drop '\n'
		input = input[:len(input)-1]
		if strings.TrimSpace(input) == "" {
			continue
		}
		fn(input)
	}
}

func selectPart() func(string) interface{} {
	// register parts
	type part struct {
		fn   func(string) interface{}
		desc string
	}

	parts := []*part{
		{part1Fn, "part 1: simple calculator, only support addition"},
		{part2Fn, "part 2: support 'minus', 'spaces' and 'long numbers' based on part 1"},
		{part3Fn, "part 3: support 'multi-ops' based on part 2"},
		{part4Fn, "part 4: refactor part3 and replace the ops into '*' and '/'"},
		{part5Fn, "part 5: support 'quadratic operations' based on part 4"},
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
	return parts[idx].fn
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
