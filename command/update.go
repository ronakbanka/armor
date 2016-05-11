package command

import (
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"golang.org/x/net/context"
)

type UpdateCommand struct {
	Meta
}

func (*UpdateCommand) Name() string {
	return "update"
}

func (*UpdateCommand) Synopsis() string {
	return "Update openstack security group"
}

func (c *UpdateCommand) Usage() string {
	return fmt.Sprintf("Usage: %s update [options]\n", c.Meta.Name)
}

func (c *UpdateCommand) SetFlags(f *flag.FlagSet) {
}

func (c *UpdateCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	c.Debugf("Start to run update command")
	return subcommands.ExitSuccess
}
