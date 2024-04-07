package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"uminer/miner-server/cmd"
	"uminer/miner-server/cmd/miner"
	"uminer/miner-server/cmd/worker"
)

func main() {
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
				Usage: "Set worker IP addresses",
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

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
