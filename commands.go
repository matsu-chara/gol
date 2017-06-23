package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/matsu-chara/gol/command"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "datapath",
		Value: fmt.Sprintf("%s/.config/gol/links.json", os.Getenv("HOME")),
		Usage: "/path/to/data/file",
	},
}

var Commands = []cli.Command{
	{
		Name:   "add",
		Usage:  "a",
		Action: command.CmdAdd,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "rm",
		Usage:  "r",
		Action: command.CmdRm,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "ls",
		Usage:  "l",
		Action: command.CmdLs,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "get",
		Usage:  "g",
		Action: command.CmdGet,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "open",
		Usage:  "o",
		Action: command.CmdOpen,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "peco",
		Usage:  "p",
		Action: command.CmdPeco,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
