package main

import "github.com/urfave/cli"

var CommitCommand = cli.Command{
	Name: "commit",
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
