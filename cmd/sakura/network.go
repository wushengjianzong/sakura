package main

import "github.com/urfave/cli"

var NetworkCommand = cli.Command{
	Name: "network",
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
