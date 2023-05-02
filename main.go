package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sshelll/interpreter/partone"
	"github.com/sshelll/interpreter/parttwo"
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

	if strings.TrimSpace(*expr) != "" {
		safeFn(*expr)
	}

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
	menu, err := menuscreen.NewMenuScreen()
	if err != nil {
		log.Fatalln(err)
	}
	defer menu.Fini()

	menu.SetTitle("parts")
	menu.AppendLines("part 1: simple calculator (only support addition)")
	menu.AppendLines("part 2: simple calculator (support 'minus', 'spaces' and 'long numbers')")
	menu.Start()

	idx, _, ok := menu.ChosenLine()
	if !ok {
		log.Fatalln("you haven't chosen any lines")
	}

	switch idx {
	case 0:
		return partoneFn
	case 1:
		return parttwoFn
	default:
		return nil
	}
}

func partoneFn(expr string) interface{} {
	it := partone.NewInterpreter(expr)
	return it.Expr()
}

func parttwoFn(expr string) interface{} {
	it := parttwo.NewInterpreter(expr)
	return it.Expr()
}
