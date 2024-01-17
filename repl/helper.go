package repl

import (
	"bufio"
	"fmt"
	"io"
	"opus-backend/evaluator"
	"opus-backend/lexer"
	"opus-backend/object"
	"opus-backend/parser"
	"opus-backend/token"
	"os"
)

const PROMPT = "opus >> "

func Start(line string) {
	// scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	// for {
	// 	fmt.Print(PROMPT)
	// 	scanned := scanner.Scan()
	// 	if !scanned {
	// 		return
	// 	}

	// line := scanner.Text()
	l := lexer.New(line)
	out := bufio.NewWriter(os.Stdout)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
	fmt.Println(evaluated)
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
	fmt.Print("\n")
	// }
}

func printParserErrors(out io.Writer, errors []string) {
	// io.WriteString(out, MONKEYFACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	fmt.Println(" parser errors:")
	for _, msg := range errors {
		fmt.Println("\t" + msg + "\n")
	}
}
