package pl

import (
	"io"
	"pomidor/evaluator"
	"pomidor/lexer"
	"pomidor/object"
	"pomidor/parser"
)

func Run(code string, out io.Writer) {
	env := object.NewEnvironment()
	l := lexer.New(code)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		panic(1)
	}

	evaluated := evaluator.Eval(program, env)

	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
