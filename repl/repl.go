package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ECecillo/redis-interpreter/lexer"
	"github.com/ECecillo/redis-interpreter/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print("redis-custom> ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		if line == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
