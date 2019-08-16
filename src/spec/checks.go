package spec

import (
	"fmt"
	"log"
	"reflect"
)

func checkReqAttrs(event *Event, name string) {
	for k, v := range defaultAttrs {
		_, ok := event.Attributes[k]
		if !ok {
			fmt.Printf("required attribute %s from event %s is missing. a default config will be applied.\n", k, name)
			event.Attributes[k] = v
			event.Required = append(event.Required, k)
		}
	}
	for _, f := range event.Required {
		_, ok := event.Attributes[f]
		if !ok {
			log.Fatalf("required attribute %s from event %s is missing", f, name)
		}
	}
}

func checkEventSpecConformity(event *Event, name string) {
}

func checkObjAttrSpecConformity(attr *Attr, name string) {
	v := reflect.ValueOf(*attr)

	unsupportedFields := []string{
		"Maximum",
		"Minimum",
		"ExclusiveMaximum",
		"ExclusionMinimum",
		"MaxLength",
		"MinLength",
		"Pattern",
		"MaxItems",
		"MinItems",
		"Enum",
		"Items",
		"Format",
		"Default",
		"Auto",
	}
	requiredFields := []string{
		"Properties",
	}

	// Warns users if one of the unsupported fields is present
	for _, n := range unsupportedFields {
		f := v.FieldByName(n)
		if !f.IsNil() {
			fmt.Printf("field %s in object attribute %s is not supported. this field will be ignored.\n", n, name)
		}
	}
	// Returns an error if one of the required fields is missing
	for _, n := range requiredFields {
		f := v.FieldByName(n)
		if f.IsNil() {
			log.Fatalf("field %s in object attribute %s is required but missing", n, name)
		}
	}
}

func checkStrAttrSpecConformity(attr *Attr, name string) {
}

func checkNumAttrSpecConformity(attr *Attr, name string) {
}

func checkBoolAttrSpecConformity(attr *Attr, name string) {
}

func checkArrayAttrSpecConformity(attr *Attr, name string) {
}

func checkAttrSpecConformity(attr *Attr, name string) {
	switch t := attr.Type; t {
	case "object":
		checkObjAttrSpecConformity(attr, name)
	case "array":
		checkArrayAttrSpecConformity(attr, name)
	case "string":
		checkStrAttrSpecConformity(attr, name)
	case "number", "integer":
		checkNumAttrSpecConformity(attr, name)
	case "boolean":
		checkBoolAttrSpecConformity(attr, name)
	default:
		log.Fatalf("unsupported type %s from attribute %s", t, name)
	}
}

func checkMetadataValidity(spec *CEGenSpec) {}
