package main

import (
	"fmt"
	"os"
	"pomidor/repl"
)

func main() {
	fmt.Print("Это Помидор!\n")
	repl.Start(os.Stdin, os.Stdout)
}
