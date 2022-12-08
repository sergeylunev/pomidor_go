package repl

import (
	"bufio"
	"fmt"
	"io"
	"pomidor/pl"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		pl.Run(scanner.Text(), out)
	}
}
