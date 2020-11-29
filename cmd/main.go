package main

import (
	"go-scaffold/internal/apis/grpc"
	"go-scaffold/internal/apis/rest"
	"go-scaffold/pkg/log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "http",
				Usage: "HTTP RESTful Server",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "port,p",
						Value: 80,
						Usage: "port",
					},
				},
				Action: func(c *cli.Context) error {
					port := c.Int("port")
					rest.Serve(port)
					return nil
				},
			},
			{
				Name:  "grpc",
				Usage: "GRPC Server",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "port,p",
						Value: 3000,
						Usage: "port",
					},
				},
				Action: func(c *cli.Context) error {
					port := c.Int("port")
					grpc.Serve(port)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err.Error())
	}
}
