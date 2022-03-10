package main

import "github.com/urfave/cli"

var ListCommand = cli.Command{
	Name: "run",
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
