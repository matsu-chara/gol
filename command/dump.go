package command

import (
	"encoding/json"
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/matsu-chara/gol/operations"
	"github.com/matsu-chara/gol/util"
)

// CmdDump dump
func CmdDump(c *cli.Context) {
	filepath := c.GlobalString("datapath")

	dumped, err := operations.RunDump(filepath)
	util.ExitIfError(err)

	dumpedJSON, err := json.MarshalIndent(dumped, "", "\t")
	util.ExitIfError(err)

	fmt.Println(string(dumpedJSON))
}
