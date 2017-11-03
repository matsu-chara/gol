package main

import (
	"os"

	"github.com/matsu-chara/gol/util"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "matsu-chara"
	app.Email = "matsuy00@gmail.com"
	app.Usage = "gol(go link) is a url shortner"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	err := app.Run(os.Args)
	util.ExitIfError(err)
}
