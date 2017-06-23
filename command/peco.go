package command

import (
	"github.com/codegangsta/cli"
	"github.com/matsu-chara/gol/operations"
	"github.com/matsu-chara/gol/util"
)

// CmdPeco peco [prefix]
func CmdPeco(c *cli.Context) {
	filepath := c.GlobalString("datapath")
	prefix := c.Args().Get(0)

	err := operations.RunPeco(filepath, prefix)
	util.ExitIfError(err)
}
