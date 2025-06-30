package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var cache string
	var storage string
	var output string
	var mode string

	cmd := &cli.Command{
		Name:  "new",
		Usage: "create new bot",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "mode",
				Value:       "polling",
				Usage:       "tg bot mode webhook or (polling)",
				Destination: &mode,
			},
			&cli.StringFlag{
				Name:        "cache",
				Value:       "memory",
				Usage:       "cache type redis or (memory)",
				Destination: &cache,
			},
			&cli.StringFlag{
				Name:        "storage",
				Value:       "file",
				Usage:       "storage type mongo or (file)",
				Destination: &storage,
			},
			&cli.StringFlag{
				Name:        "output",
				Value:       "",
				Aliases:     []string{"o"},
				Usage:       "путь создания",
				Destination: &output,
			},
		},
		Action: func(context.Context, *cli.Command) error {
			generateFiles(&GenConfig{mode, cache, storage, output})
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
