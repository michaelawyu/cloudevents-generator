package spec

import (
	"fmt"

	"github.com/michaelawyu/cloudevents-generator/src/logger"
)

// AutoFormat is
type AutoFormat struct {
	IsUUIDv4  bool
	IsRFC3339 bool
}

// GetAutoFormat is
func GetAutoFormat(t string, f string) *AutoFormat {
	if t != "string" {
		logger.Logger.Fatal(fmt.Sprintf("format %s works only with attributes of type %s", f, t))
	}

	var af AutoFormat

	switch f {
	case "uuid", "UUIDv4":
		af = AutoFormat{IsUUIDv4: true}
		return &af
	case "timestamp", "RFC3339":
		af = AutoFormat{IsRFC3339: true}
		return &af
	default:
		logger.Logger.Fatal(fmt.Sprintf("format %s cannot be auto generated", f))
	}

	return nil
}
