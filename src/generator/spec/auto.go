package spec

import "log"

// AutoFormat is
type AutoFormat struct {
	IsUUIDv4  bool
	IsRFC3339 bool
}

// GetAutoFormat is
func GetAutoFormat(t string, f string) AutoFormat {
	if t != "string" {
		log.Fatalf("format %s works only with attributes of type %s", f, t)
	}

	switch f {
	case "uuid", "UUIDv4":
		return AutoFormat{IsUUIDv4: true}
	case "timestamp", "RFC3339":
		return AutoFormat{IsRFC3339: true}
	default:
		log.Fatalf("format %s cannot be auto generated", f)
	}

	return AutoFormat{}
}
