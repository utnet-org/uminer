package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"uminer/miner-server/cmd"
	"uminer/miner-server/cmd/miner"
	"uminer/miner-server/cmd/utlog"
	"uminer/miner-server/cmd/worker"
)

var StorageMiner storageMiner

type storageMiner struct{}

func main() {

	// set up logs
	utlog.SetupLogLevels()

	app := &cli.App{
		Name:                 "utility miner",
		Usage:                "Utility decentralized network miner",
		Version:              cmd.UserVersion(),
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "storage",
				Usage: "Set storage data",
			},
			&cli.StringFlag{
				Name:  "serverip",
				Usage: "Set miner server IP address",
				Value: "127.0.0.1",
			},
			&cli.StringSliceFlag{
				Name:  "workerip",
				Usage: "Set all workers IP addresses of a miner",
				Value: cli.NewStringSlice(),
			},
			&cli.StringFlag{
				Name:  "node",
				Usage: "Set nodes address",
				Value: cmd.NodeURL,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "startminer",
				Aliases: []string{"s m"},
				Usage:   "Start the miner server",
				Action:  miner.StartMinerServer,
			},
			{
				Name:    "startworker",
				Aliases: []string{"s w"},
				Usage:   "Start the worker server",
				Action:  worker.StartWorkerServer,
			},
		},
	}

	app.Setup()
	app.Metadata["repoType"] = StorageMiner
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
