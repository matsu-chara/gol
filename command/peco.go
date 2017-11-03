package command

import (
	"github.com/matsu-chara/gol/operations/exec"
	"github.com/matsu-chara/gol/util"
	"github.com/urfave/cli"
)

// CmdPeco peco [prefix]
func CmdPeco(c *cli.Context) {
	filepath := c.GlobalString("datapath")
	prefix := c.Args().Get(0)

	err := exec.RunPeco(filepath, prefix)
	util.ExitIfError(err)
}
