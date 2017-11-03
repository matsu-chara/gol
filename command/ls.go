package command

import (
	"fmt"

	"github.com/matsu-chara/gol/operations"
	"github.com/matsu-chara/gol/util"
	"github.com/urfave/cli"
)

// CmdLs ls
func CmdLs(c *cli.Context) {
	filepath := c.GlobalString("datapath")

	entries, err := operations.RunLs(filepath)
	util.ExitIfError(err)
	for _, entry := range entries {
		fmt.Println(entry)
	}
}
