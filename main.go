package main

import (
	"Guerilla/repl"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is Guerilla.")
	repl.Start(os.Stdin, os.Stdout)
}
