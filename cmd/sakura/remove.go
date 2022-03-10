package main

import "github.com/urfave/cli"

var RemoveCommand = cli.Command{
	Name: "rm",
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
