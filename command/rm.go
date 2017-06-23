package command

import (
	"github.com/codegangsta/cli"
	"github.com/matsu-chara/gol/operations"
	"github.com/matsu-chara/gol/util"
)

// CmdRm rm key
func CmdRm(c *cli.Context) {
	filepath := c.GlobalString("datapath")
	key := c.Args().Get(0)

	err := operations.RunRm(filepath, key)
	util.ExitIfError(err)
}
