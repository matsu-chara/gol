package command

import (
	"github.com/codegangsta/cli"
	"github.com/matsu-chara/gol/operations"
	"github.com/matsu-chara/gol/util"
)

// CmdServer server
func CmdServer(c *cli.Context) {
	filepath := c.GlobalString("datapath")
	port := c.Uint("port")

	err := operations.RunServer(filepath, port)
	util.ExitIfError(err)
}
