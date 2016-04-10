package main

import (
	"fmt"
	"os"
	"time"

	"github.com/codegangsta/cli"
)

const (
	Version = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "t, template",
			Usage: "template to use as generator",
		},
		cli.DurationFlag{
			Name:  "d, duration",
			Usage: "test duration in seconds",
			Value: 60 * time.Second,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "kafka",
			ShortName: "k",
			Usage:     "send payload to kafka",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "zk, zookeeper",
					Usage: "zookeeper connection string, comma separated",
				},
				cli.StringFlag{
					Name:  "topic",
					Usage: "topic to use",
				},
			},
			Action: func(c *cli.Context) {
				execute(&ActionConfig{
					Type:             "kafka",
					ConnectionString: c.String("zookeeper"),
					Topic:            c.String("topic"),
					Template:         c.GlobalString("template"),
					Duration:         c.GlobalDuration("duration"),
				})
			},
		},
		{
			Name:      "stdout",
			ShortName: "s",
			Usage:     "send payload to /dev/stdout",
			Action: func(c *cli.Context) {
				execute(&ActionConfig{
					Type:     "stdout",
					Template: c.GlobalString("template"),
					Duration: c.GlobalDuration("duration"),
				})
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Unexpected failure:", err)
		os.Exit(1)
	}
}
