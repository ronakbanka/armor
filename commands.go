package main

import (
	"github.com/google/subcommands"
	"github.com/ronakbanka/armor/command"
)

func SubCommands(meta command.Meta) []subcommands.Command {
	return []subcommands.Command{
		&command.FetchCommand{
			meta,
		},
		&command.UpdateCommand{
			meta,
		},
	}
}
