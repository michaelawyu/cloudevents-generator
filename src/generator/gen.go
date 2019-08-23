package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/michaelawyu/cloud-events-generator/src/generator/python"
	"github.com/michaelawyu/cloud-events-generator/src/logger"

	"github.com/michaelawyu/cloud-events-generator/src/config"
	"github.com/michaelawyu/cloud-events-generator/src/spec"
	"gopkg.in/yaml.v2"
)

// Generate is
func Generate(cfg config.GenConfig) {
	d, err := ioutil.ReadFile(cfg.Input)
	if err != nil {
		logger.Logger.Fatal(fmt.Sprintf("cannot read cloud event specification %s: %s", cfg.Input, err))
	}

	var spec spec.CEGenSpec
	switch ext := filepath.Ext(strings.ToLower(cfg.Input)); ext {
	case ".json":
		err = json.Unmarshal(d, &spec)
		if err != nil {
			logger.Logger.Fatal(fmt.Sprint("cannot unmarshal JSON file %s: %s", cfg.Input, err))
		}
	case ".yaml":
		err = yaml.Unmarshal(d, &spec)
		if err != nil {
			logger.Logger.Fatal(fmt.Sprintf("cannot unmarshal YAML file %s: %s", cfg.Input, err))
		}
	default:
		logger.Logger.Fatal(fmt.Sprintf("unsupported file extension %s (%s); requires a JSON or YAML file", ext, cfg.Input))
	}

	ms, meta := spec.Parse()

	bs := cfg.Binding.ToSelector()

	err = os.MkdirAll(cfg.Output, os.FileMode(0777))
	if err != nil {
		logger.Logger.Fatal(fmt.Sprintf("cannot create folder(s) in path %s: %s", cfg.Output, err))
	}
	switch n := cfg.Language.Name; n {
	case "python":
		logger.Logger.Info("generating python package")
		python.GenPkg(cfg.Output, ms, bs, meta)
	}

	logger.Logger.Success(fmt.Sprintf("successfully generated package at %s", cfg.Output))
}
