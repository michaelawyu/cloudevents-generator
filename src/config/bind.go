package config

import (
	"fmt"

	genspec "github.com/michaelawyu/cloud-events-generator/src/generator/spec"
	"github.com/michaelawyu/cloud-events-generator/src/logger"
)

// Binding specifies the name of a supported binding
type Binding struct {
	Name string
}

// A list of supported bindings
var bindings = []Binding{
	// JSON
	Binding{
		Name: "JSON",
	},
	// HTTP
	Binding{
		Name: "HTTP",
	},
}

// ToSelector is
func (b Binding) ToSelector() genspec.BindSelector {
	switch b.Name {
	case "JSON":
		return genspec.BindSelector{
			IsJSON: true,
		}
	case "HTTP":
		return genspec.BindSelector{
			IsHTTP: true,
		}
	default:
		logger.Logger.Fatal(fmt.Sprintf("binding %s is not supported", b.Name))
	}
	return genspec.BindSelector{}
}

// GetBinding matches the input with one of the supported bindings
func GetBinding(name string) Binding {
	for _, b := range bindings {
		if name == b.Name {
			return b
		}
	}

	return Binding{}
}
