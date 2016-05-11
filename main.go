package main

import "os"

func main() {
	command := &Command{
		outStream: os.Stdout,
		errStream: os.Stderr,
	}
	os.Exit(command.Run(os.Args[1:]))
}
