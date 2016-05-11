package command

import (
	"testing"

	"github.com/google/subcommands"
)

func TestFetchCommand_implement(t *testing.T) {
	var _ subcommands.Command = &FetchCommand{}
}
