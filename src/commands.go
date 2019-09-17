package main

import (
	"github.com/michaelawyu/cloudevents-generator/src/config"
	"github.com/michaelawyu/cloudevents-generator/src/generator"
	"github.com/michaelawyu/cloudevents-generator/src/logger"
	"github.com/michaelawyu/cloudevents-generator/src/utils"
	"github.com/urfave/cli"
)

func generate(c *cli.Context) {
	ip := c.String("input")
	op := c.String("output")
	op = utils.FormatPath(op)
	lang := c.String("language")
	bind := c.String("binding")
	vLogEnabled := c.Bool("verbose")

	if ip == "" || op == "" || lang == "" {
		logger.Logger.Fatal("Required parameters --input, --ouptut, and/or --language are missing; wiew help message with --help")
	}

	config := config.GetConfig(ip, op, lang, bind)

	if vLogEnabled {
		logger.Logger.VLogEnabled = true
	}

	generator.Generate(config)
}
