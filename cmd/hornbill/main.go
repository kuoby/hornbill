package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "hornbill",
		Usage: "The server control plane for Mongoose, written in Go.",
		Commands: []*cli.Command{
			newRunCommand(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Usage: "Load configuration from `FILE`",
				Value: "/etc/hornbill/config.toml",
				Aliases: []string{
					"c",
				},
			},
		},
		EnableBashCompletion: true,
		Suggest:              true,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
