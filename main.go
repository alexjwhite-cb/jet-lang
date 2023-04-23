package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alexjwhite-cb/track-lang/lexer"
)

func main() {
	out, _ := lexer.Lex(strings.Join(os.Args[1:], "\n"))
	for _, o := range out {
		fmt.Printf("%+v\n", o)
	}
}
