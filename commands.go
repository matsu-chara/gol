package main

import (
	"fmt"
	"os"

	"github.com/matsu-chara/gol/command"
	"github.com/urfave/cli"
)

// GlobalFlags are Flags
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "datapath",
		Value: fmt.Sprintf("%s/.config/gol/links.json", os.Getenv("HOME")),
		Usage: "/path/to/data/file",
	},
}

// Commands is a list of commands
var Commands = []cli.Command{
	{
		Name:      "add",
		ShortName: "a",
		Action:    command.CmdAdd,
		Flags:     []cli.Flag{},
	},
	{
		Name:      "rm",
		ShortName: "r",
		Action:    command.CmdRm,
		Flags:     []cli.Flag{},
	},
	{
		Name:      "ls",
		ShortName: "l",
		Action:    command.CmdLs,
		Flags:     []cli.Flag{},
	},
	{
		Name:      "get",
		ShortName: "g",
		Action:    command.CmdGet,
		Flags:     []cli.Flag{},
	},
	{
		Name:      "dump",
		ShortName: "d",
		Action:    command.CmdDump,
		Flags:     []cli.Flag{},
	},
	{
		Name:      "open",
		ShortName: "o",
		Action:    command.CmdOpen,
		Flags:     []cli.Flag{},
	},
	{
		Name:      "peco",
		ShortName: "p",
		Action:    command.CmdPeco,
		Flags:     []cli.Flag{},
	},
	{
		Name:      "server",
		ShortName: "s",
		Action:    command.CmdServer,
		Flags: []cli.Flag{
			cli.UintFlag{
				Name:  "port",
				Value: 5656,
				Usage: "server port default = 5656",
			},
		},
	},
}

// CommandNotFound is handler for undefined command
func CommandNotFound(c *cli.Context, command string) {
	_, err := fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(2)
}
