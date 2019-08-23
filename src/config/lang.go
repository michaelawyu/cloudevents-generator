package config

import (
	"fmt"

	"github.com/michaelawyu/cloud-events-generator/src/logger"
)

// Language specifies the name and template paths of a supported languages
type Language struct {
	Name string
}

// A list of supported languages
var langs = []Language{
	// Python
	Language{
		Name: "python",
	},
}

// GetLanguage matches the input with one of the supported languages
func GetLanguage(name string) Language {
	for _, l := range langs {
		if name == l.Name {
			return l
		}
	}

	logger.Logger.Fatal(fmt.Sprintf("language %s is not supported", name))
	return Language{}
}
