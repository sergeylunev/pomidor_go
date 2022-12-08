package main

import (
	"os"
	"pomidor/pl"
	"pomidor/repl"
)

func main() {
	args := os.Args[1:]

	if len(args) >= 1 {
		fromFile(args[0])
	} else {
		startRepl()
	}
}

func fromFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	pl.Run(string(data), os.Stdout)
}

func startRepl() {
	repl.Start(os.Stdin, os.Stdout)
}
