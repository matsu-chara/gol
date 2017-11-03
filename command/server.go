package command

import (
	"github.com/matsu-chara/gol/server"
	"github.com/matsu-chara/gol/util"
	"github.com/urfave/cli"
)

// CmdServer server
func CmdServer(c *cli.Context) {
	filepath := c.GlobalString("datapath")
	port := c.Uint("port")

	err := server.RunServer(filepath, port)
	util.ExitIfError(err)
}
