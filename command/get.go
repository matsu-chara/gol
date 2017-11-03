package command

import (
	"fmt"

	"github.com/matsu-chara/gol/operations"
	"github.com/matsu-chara/gol/util"
	"github.com/urfave/cli"
)

// CmdGet get key
func CmdGet(c *cli.Context) {
	filepath := c.GlobalString("datapath")
	key := c.Args().Get(0)

	entry, err := operations.RunGet(filepath, key)
	if entry == nil {
		return
	}
	util.ExitIfError(err)
	fmt.Println(entry.Value)
}
