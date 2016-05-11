package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/ronakbanka/armor/command"
)

//GlobalFlags returns flags for global options
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "config, c",
		Usage:  "Openstack auth config path",
		Value:  command.GetBaseDir(),
		EnvVar: "ARMOR_CONFIG_PATH",
	},
}

// Commands returns cli commands
var Commands = []cli.Command{
	{
		Name:   "get-sec-groups",
		Usage:  "Fetch openstack security groups",
		Action: command.CmdGetSecGroups,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "tenant, t",
				Usage:  "Openstack tenant name",
				EnvVar: "OS_TENANT_NAME",
			},
			cli.StringFlag{
				Name:  "sec-group, sg",
				Usage: "Openstack Security group name",
			},
		},
	},
	{
		Name:   "update-sec-groups",
		Usage:  "Update openstack security groups",
		Action: command.CmdUpdateSecGroups,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "dry-run",
				Usage: "Prints output without updating openstack sec groups",
			},
		},
	},
}

//CommandNotFound returns error for invalid commands
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

// Exists check for file existance
func Exists(name string) (bool, error) {
	_, err := os.Stat(name)

	if os.IsNotExist(err) {
		return false, err
	} else if err == nil {
		return true, nil
	}

	return false, err
}
