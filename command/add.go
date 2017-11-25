package command

import (
	"github.com/matsu-chara/gol/operations"
	"github.com/matsu-chara/gol/util"
	"github.com/urfave/cli"
)

// CmdAdd add key value
func CmdAdd(c *cli.Context) {
	filepath := c.GlobalString("datapath")
	key := c.Args().Get(0)
	link := c.Args().Get(1)
	registeredBy := c.Args().Get(2)
	isForce := c.Bool("force")

	err := operations.RunAdd(filepath, key, link, registeredBy, isForce)
	util.ExitIfError(err)
}
