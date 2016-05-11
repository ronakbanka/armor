package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "ronakbanka"
	app.Email = "ronak.banka@rakuten.com"
	app.Usage = "manage openstack security groups"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound
	app.Run(os.Args)
}
