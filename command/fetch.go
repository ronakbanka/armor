package command

import (
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"golang.org/x/net/context"
)

type FetchCommand struct {
	Meta
}

func (*FetchCommand) Name() string {
	return "fetch"
}

func (*FetchCommand) Synopsis() string {
	return "Fetch openstack security group"
}

func (c *FetchCommand) Usage() string {
	return fmt.Sprintf("Usage: %s fetch [options]\n", c.Meta.Name)
}

func (c *FetchCommand) SetFlags(f *flag.FlagSet) {
}

func (c *FetchCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	c.Debugf("Start to run fetch command")
	return subcommands.ExitSuccess
}
