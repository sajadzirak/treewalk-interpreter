package main

import (
	"fmt"
	"os"
	"treewalk-interpreter/repl"
)

func main() {
	fmt.Println("This is the TreeWalk Interpreter.")
	repl.Start(os.Stdin, os.Stdout)
}
