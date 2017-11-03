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
	value := c.Args().Get(1)

	err := operations.RunAdd(filepath, key, value)
	util.ExitIfError(err)
}
