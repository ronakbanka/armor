package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"

	"golang.org/x/net/context"

	"github.com/google/subcommands"
	"github.com/ronakbanka/armor/command"
)

type Command struct {
	outStream, errStream io.Writer
}

func (c *Command) Run(args []string) int {
	var (
		version bool
		debug   bool
	)

	// Define global flags
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&version, "version", false, "Show version and quit")
	flags.BoolVar(&debug, "debug", false, "Run as a debug mode")

	if err := flags.Parse(args); err != nil {
		return int(subcommands.ExitFailure)
	}

	// Show version and quit
	if version {
		var buf bytes.Buffer
		fmt.Fprintf(&buf, "%s version %s", Name, Version)
		if len(GitCommit) != 0 {
			fmt.Fprintf(&buf, " (%s)", GitCommit)
		}

		fmt.Fprint(c.outStream, buf.String()+"\n")
		return int(subcommands.ExitSuccess)
	}

	// Create commander and set default commands
	commander := subcommands.NewCommander(flags, Name)
	commander.Register(commander.HelpCommand(), Name)
	commander.Register(commander.FlagsCommand(), Name)

	ctx := context.Background()

	meta := command.Meta{
		Name:      Name,
		Debug:     debug,
		OutStream: c.outStream,
		ErrStream: c.errStream,
	}

	for _, subCmd := range SubCommands(meta) {
		commander.Register(subCmd, Name)
	}

	return int(commander.Execute(ctx))
}
