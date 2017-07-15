package command

import (
	"github.com/codegangsta/cli"
	"github.com/matsu-chara/gol/operations/exec"
	"github.com/matsu-chara/gol/util"
)

// CmdOpen open key
func CmdOpen(c *cli.Context) {
	filepath := c.GlobalString("datapath")
	key := c.Args().Get(0)
	err := exec.RunOpen(filepath, key)
	util.ExitIfError(err)
}
