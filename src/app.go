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
	app.Usage = ""
	app.Commands = []cli.Command{
		cli.Command{
			Name:    "version",
			Aliases: []string{"ver"},
			Usage:   "Returns the version of the Cloud Events Generator.",
		},
		cli.Command{
			Name:        "generate",
			Aliases:     []string{"gen"},
			Usage:       "Generates code with specified generator.",
			UsageText:   "",
			Description: "",
			ArgsUsage:   "",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input, i",
					Usage: "Specifies the `PATH` to the event specification. Required.",
				},
				cli.StringFlag{
					Name:  "output, o",
					Usage: "Specifies the `PATH` where generated code is saved. Optional; uses current directory if not specified.",
				},
				cli.StringFlag{
					Name:  "language, lang",
					Usage: "Specifies the `LANGUAGE` to use. Required.",
				},
				cli.StringFlag{
					Name:  "binding, bind",
					Usage: "Specifies the transport `BINDING` to use. Optional; uses JSON binding if not specified.",
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
