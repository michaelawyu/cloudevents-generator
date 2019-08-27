package main

import (
	"fmt"
	"os"

	"github.com/michaelawyu/cloud-events-generator/src/logger"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Cloud Events Generator"
	app.Usage = "Help you easily produce, consume, and collaborate on Cloud Events"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		cli.Command{
			Name:    "version",
			Aliases: []string{"ver"},
			Usage:   "Returns the version of the Cloud Events Generator.",
		},
		cli.Command{
			Name:    "generate",
			Aliases: []string{"gen"},
			Usage:   "Generates code with specified generator.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input, i",
					Usage: "The `PATH` to the input Cloud Events Generator specification. Required.",
				},
				cli.StringFlag{
					Name:  "output, o",
					Usage: "The `PATH` where the generated event library is saved. Required.",
				},
				cli.StringFlag{
					Name:  "language, lang",
					Usage: "The `LANGUAGE` to use. Required.",
				},
				cli.StringFlag{
					Name:  "binding, bind",
					Usage: "The transport `BINDING` to use. Optional; uses JSON binding if not specified.",
				},
				cli.BoolFlag{
					Name:  "verbose, v",
					Usage: "Enables verbose logging. Optional; disabled if not specified",
				},
			},
			Action: generate,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Logger.Fatal(fmt.Sprintf("%s", err))
	}
}
