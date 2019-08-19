package main

import (
	"github.com/michaelawyu/cloud-events-generator/src/config"
	"github.com/michaelawyu/cloud-events-generator/src/generator"
	"github.com/michaelawyu/cloud-events-generator/src/utils"
	"github.com/urfave/cli"
)

func generate(c *cli.Context) {
	ip := c.String("input")
	op := c.String("output")
	op = utils.FormatPath(op)
	lang := c.String("language")
	bind := c.String("binding")

	config := config.GetConfig(ip, op, lang, bind)
	generator.Generate(config)
}
