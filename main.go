package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/sshelll/interpreter/partone"
)

var (
	expr        = flag.String("e", "", "expression to evaluate")
	interactive = flag.Bool("i", false, "interactive mode")
)

func main() {

	flag.Parse()

	fn := func(expr string) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		fmt.Println(partoneFn(expr))
	}

	if strings.TrimSpace(*expr) != "" {
		fn(*expr)
	} else if *interactive {
		for {
			print(">> ")
			var input string
			_, err := fmt.Scanln(&input)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if strings.TrimSpace(input) == "" {
				continue
			}
			fn(input)
		}
	} else {
		flag.Usage()
	}

}

func partoneFn(expr string) interface{} {
	it := partone.NewInterpreter(expr)
	return it.Expr()
}
