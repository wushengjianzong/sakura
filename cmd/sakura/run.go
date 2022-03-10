package main

import (
	"errors"
	"fmt"
	"log"
	"sakura/pkg/sandbox"
	"strings"

	"github.com/google/uuid"
	"github.com/urfave/cli"
)

var RunCommand = cli.Command{
	Name: "run",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "tty",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "detach",
		},
		cli.StringFlag{
			Name:  "mem",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpu",
			Usage: "cpu quota",
		},
		cli.StringFlag{
			Name:  "v",
			Usage: "volume",
		},
		cli.StringSliceFlag{
			Name:  "e",
			Usage: "environment varible",
		},
		cli.BoolFlag{
			Name: "name",
		},
		cli.BoolFlag{
			Name: "e",
		},
	},
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return errors.New("missing container command")
		}

		args := make([]string, 0, len(ctx.Args()))
		for _, arg := range ctx.Args() {
			args = append(args, string(arg))
		}

		tty, detach := ctx.Bool("t"), ctx.Bool("d")
		if tty && detach {
			return errors.New("it and d can not both provided")
		}

		var (
			command []string = args[1:]
			name    string   = ctx.String("name")
			volume  string   = ctx.String("v")
			image   string   = args[0]
			envs    []string = ctx.StringSlice("e")
		)
		Run(tty, command, nil, name, volume, image, envs)
		return nil
	},
}

func Run(tty bool, command []string, resource interface{}, name string, volume string, image string, envs []string) {
	containerdId, err := uuid.NewRandom()
	if err != nil {
		log.Fatalln(err)
	}
	if name == "" {
		name = fmt.Sprintf("%s-%s", image, containerdId.String())
	}

	parentCmd, writePipe := sandbox.NewParentProcess(tty, name, volume, image, envs)
	if parentCmd == nil {
		log.Fatalln("new parenet process failed")
	}
	if err := parentCmd.Start(); err != nil {
		log.Println(err)
	}

	// todo: save container info
	// todo: cgroups

	// todo: network

	if _, err := writePipe.WriteString(strings.Join(command, " ")); err != nil {
		log.Fatalln(err)
	}
	if err := writePipe.Close(); err != nil {
		log.Fatalln(err)
	}

	if tty {
		// if err := parentCmd.Wait(); err != nil {
		// 	log.Fatalln(err)
		// }
		_ = parentCmd.Wait()
		// todo: delete container info
		// todo: delete workspace
	}
}
