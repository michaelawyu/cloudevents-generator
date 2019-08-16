package spec

// GetDefaultIDAttr returns the default ID attribute for cloud events
func GetDefaultIDAttr() Attr {
	format := "uuid"
	auto := true
	minLength := 1
	return Attr{
		Type:      "string",
		Format:    &format,
		Auto:      &auto,
		MinLength: &minLength,
	}
}

var defaultIDAttr = GetDefaultIDAttr()

// GetDefaultSourceAttr returns the default source attribute for cloud events
func GetDefaultSourceAttr() Attr {
	format := "uri"
	return Attr{
		Type:   "string",
		Format: &format,
	}
}

var defaultSourceAttr = GetDefaultSourceAttr()

// GetDefaultSpecVersionAttr returns the default specversion attribute for cloud events
func GetDefaultSpecVersionAttr() Attr {
	var defaultValue interface{}
	defaultValue = "0.3"
	minLength := 1
	return Attr{
		Type:      "string",
		Default:   &defaultValue,
		MinLength: &minLength,
	}
}

var defaultSpecVersionAttr = GetDefaultSpecVersionAttr()

// GetDefaultTypeAttr returns the default type attribute for cloud events
func GetDefaultTypeAttr() Attr {
	minLength := 1
	return Attr{
		Type:      "string",
		MinLength: &minLength,
	}
}

var defaultTypeAttr = GetDefaultTypeAttr()

var defaultAttrs = map[string]Attr{
	"id":          defaultIDAttr,
	"source":      defaultSourceAttr,
	"specversion": defaultSpecVersionAttr,
	"type":        defaultTypeAttr,
}
