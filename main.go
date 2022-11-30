package main

import (
	"fmt"
	"os"
	"os/user"
	"pomidor/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Monkey PL!\n", user.Username)
	fmt.Printf("Feel free to type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
