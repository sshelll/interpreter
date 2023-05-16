package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	expr        = flag.String("e", "", "expression to evaluate")
	interactive = flag.Bool("i", false, "interactive mode")
	file        = flag.String("f", "", "use file content as input")
	debugLexer  = flag.Bool("dl", false, "debug lexer")
	p           = flag.Int("p", 0, "select part")
)

func main() {

	flag.Parse()

	if strings.TrimSpace(*expr) == "" &&
		strings.TrimSpace(*file) == "" &&
		!*interactive &&
		!*debugLexer {
		flag.Usage()
		return
	}

	// choose a part, and wrap the func
	fn, lexerFn := selectPart()
	safeFn := func(expr string) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		fmt.Println(fn(expr))
	}

	// debug lexer
	if *debugLexer {
		lexerFn(*expr)
		return
	}

	// expr mode
	if strings.TrimSpace(*expr) != "" {
		safeFn(*expr)
	}

	// interactive mode
	if *interactive {
		interactiveMode(safeFn)
	}

	// file mode
	if strings.TrimSpace(*file) != "" {
		fileMode(safeFn)
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

func fileMode(fn func(string)) {
	fn(readFile(*file))
}

func readFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("read file %s failed, err = %v\n", *file, err)
	}
	return string(content)
}
