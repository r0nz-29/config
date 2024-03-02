package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if !validArgs(args) {
		println("use '<cmd> help' for usage instructions")
		return
	}
	handle(args)
}
