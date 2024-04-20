package main

import (
	"fmt"
	"os"

	"github.com/ECecillo/redis-interpreter/repl"
)

func main() {
	fmt.Println("======= DIY REDIS REPL / ECecillo ======")
	fmt.Println("This is a simple Redis-like REPL")
	fmt.Println("Type 'exit' to quit the REPL")
	repl.Start(os.Stdin, os.Stdout)
}
