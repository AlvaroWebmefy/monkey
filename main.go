package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/alwarmra/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the monkey programming language.\n", user.Username)
	fmt.Printf("Fell free to type in command \n")
	repl.Start(os.Stdin, os.Stdout)
}
