package main

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func newRunCommand() *cli.Command {
	return &cli.Command{
		Name: "run",
		Aliases: []string{
			"r",
		},
		Usage: "Runs an API server allowing programmatic control over game servers.",
		Action: func(cCtx *cli.Context) error {
			lgr, err := zap.NewProduction()
			if err != nil {
				return errors.Wrap(err, "could not create production logger")
			}
			defer lgr.Sync()

			return nil
		},
	}
}
