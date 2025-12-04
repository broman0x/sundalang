package main

import (
	"fmt"
	"os"
	"sundalang/pkg/sundalang" 
)

func runFile(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Gagal maca file: %s\n", err)
		return
	}

	code := string(content)
	l := sundalang.NewLexer(code)
	p := sundalang.NewParser(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Printf("Parser Errors:\n")
		for _, msg := range p.Errors() {
			fmt.Printf("\t%s\n", msg)
		}
		return
	}

	env := sundalang.NewEnvironment()
	evalResult := sundalang.Eval(program, env)
	if evalResult != nil && evalResult.Type() == sundalang.ERROR_OBJ {
		fmt.Println(evalResult.Inspect())
	}
}

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg == "-v" || arg == "--version" {
			fmt.Println("SundaLang v1.0.1")
			fmt.Println("Basa Pemrograman Sunda Pandeglang")
			fmt.Println("Immersive Environment Build")
			return
		}

		runFile(arg)
		return
	}
	sundalang.StartREPL(os.Stdin, os.Stdout)
}