package main

import (
	"sync"

	"github.com/kuoby/hornbill/config"
	"github.com/kuoby/hornbill/environment"
	"github.com/kuoby/hornbill/mongoose"
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

			cfg, err := config.ParseFile(cCtx.String("config"))
			if err != nil {
				lgr.Fatal("could not parse configuration from file", zap.Error(err))
			}

			var psr environment.Provisioner
			switch cfg.Environment.Provisioner {
			case "docker":
				psr = environment.NewDocker()
			default:
				lgr.Fatal("unknown provisioner was defined", zap.String("provisioner", cfg.Environment.Provisioner))
			}
			if err := psr.Configure(); err != nil {
				lgr.Fatal("could not configure provisioner", zap.Error(err))
			}

			mgCli := mongoose.NewClient(cfg.Mongoose.BaseURL)
			lgr.Info("fetching a list of servers from mongoose")
			srvs, err := mgCli.Servers(cCtx.Context)
			if err != nil {
				lgr.Fatal("could not fetch a list of servers from mongoose", zap.Error(err))
			}

			var wg sync.WaitGroup
			for _, srv := range srvs {
				wg.Add(1)

				go func(server mongoose.Server) {
					defer wg.Done()

					lgr.Info("processing server configuration", zap.String("uuid", server.UUID))
				}(srv)
			}
			wg.Wait()
			lgr.Info("finished processing server configurations")

			return nil
		},
	}
}
