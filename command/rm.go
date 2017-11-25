package command

import (
	"github.com/matsu-chara/gol/operations"
	"github.com/matsu-chara/gol/util"
	"github.com/urfave/cli"
)

// CmdRm rm key
func CmdRm(c *cli.Context) {
	filepath := c.GlobalString("datapath")
	key := c.Args().Get(0)
	name := c.Args().Get(1)

	err := operations.RunRm(filepath, key, name)
	util.ExitIfError(err)
}
