package main

import (
	"fmt"
	"os"
	"sundalang/pkg/sundalang" 
)

func run(code string) {
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
		filename := os.Args[1]
		fmt.Printf("Macakeun file: %s...\n", filename)
		
		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Gagal maca file: %s\n", err)
			return
		}

		run(string(content))
		return
	}

	const sundaCode = `
	cetakkeun("=== MODE DEMO SUNDALANG ===")
	cetakkeun("Tip: Jalankeun 'go run cmd/main.go file.sl' pikeun muka file sorangan.")
	
	tanda x = 10
	tanda y = 20
	
	lamun x < y {
		cetakkeun("X leuwih leutik ti Y")
	}
	`
	run(sundaCode)
}