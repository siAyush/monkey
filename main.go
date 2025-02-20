package main

import (
	"os"

	"github.com/siAyush/monkey/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
