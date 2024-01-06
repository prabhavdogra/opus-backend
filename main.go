package main

import (
	"fmt"
	"opus/repl"
	"os"
)

func main() {
	fmt.Printf("Welcome to Opus: A musical programming language!\n")
	// repl - Read, Evaluate, Print, Loop
	repl.Start(os.Stdin, os.Stdout)
}
