package spec

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/michaelawyu/cloudevents-generator/src/logger"
)

func checkReqAttrs(event *Event, name string) {
	for k, v := range defaultAttrs {
		_, ok := event.Attributes[k]
		if !ok {
			logger.Logger.Warn(fmt.Sprintf("required attribute %s from event %s is missing. a default config will be applied", k, name))
			event.Attributes[k] = v
			event.Required = append(event.Required, k)
		}
	}
	for _, f := range event.Required {
		_, ok := event.Attributes[f]
		if !ok {
			logger.Logger.Fatal(fmt.Sprintf("required attribute %s from event %s is missing", f, name))
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
		"ExclusiveMinimum",
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

	// Warns users if one or more of the unsupported fields is present
	for _, n := range unsupportedFields {
		f := v.FieldByName(n)
		if !f.IsNil() {
			logger.Logger.Warn(fmt.Sprintf("field %s in object attribute %s is not supported. this field will be ignored", n, name))
		}
	}
	// Returns an error if one of the required fields is missing
	for _, n := range requiredFields {
		f := v.FieldByName(n)
		if f.IsNil() {
			logger.Logger.Fatal(fmt.Sprintf("field %s in object attribute %s is required but missing", n, name))
		}
	}
}

func checkStrAttrSpecConformity(attr *Attr, name string) {
	v := reflect.ValueOf(*attr)

	unsupportedFields := []string{
		"Maximum",
		"Minimum",
		"ExclusiveMaximum",
		"ExclusiveMinimum",
		"MaxItems",
		"MinItems",
		"Items",
		"Properties",
		"Required",
	}

	// Warns users if one or more of the unsupported fields is present
	for _, n := range unsupportedFields {
		f := v.FieldByName(n)
		if !f.IsNil() {
			logger.Logger.Warn(fmt.Sprintf("field %s in string attribute %s is not supported. this field will be ignored", n, name))
		}
	}
}

func checkNumAttrSpecConformity(attr *Attr, name string) {
	v := reflect.ValueOf(*attr)

	unsupportedFields := []string{
		"MaxLength",
		"MinLength",
		"Pattern",
		"MaxItems",
		"MinItems",
		"Items",
		"Properties",
		"Required",
	}

	// Warns users if one or more of the unsupported fields is present
	for _, n := range unsupportedFields {
		f := v.FieldByName(n)
		if !f.IsNil() {
			logger.Logger.Warn(fmt.Sprintf("field %s in number attribute %s is not supported. this field will be ignored", n, name))
		}
	}
}

func checkBoolAttrSpecConformity(attr *Attr, name string) {
	v := reflect.ValueOf(*attr)

	unsupportedFields := []string{
		"Maximum",
		"Minimum",
		"ExclusiveMaximum",
		"ExclusiveMinimum",
		"MaxLength",
		"MinLength",
		"Pattern",
		"MaxItems",
		"MinItems",
		"Enum",
		"Items",
		"Format",
		"Auto",
		"Properties",
		"Required",
	}

	// Warns users if one or more of the unsupported fields is present
	for _, n := range unsupportedFields {
		f := v.FieldByName(n)
		if !f.IsNil() {
			logger.Logger.Warn(fmt.Sprintf("field %s in boolean attribute %s is not supported. this field will be ignored", n, name))
		}
	}
}

func checkArrayAttrSpecConformity(attr *Attr, name string) {
	v := reflect.ValueOf(*attr)

	unsupportedFields := []string{
		"Maximum",
		"Minimum",
		"ExclusiveMaximum",
		"ExclusiveMinimum",
		"MaxLength",
		"MinLength",
		"Pattern",
		"Format",
		"Auto",
		"Properties",
		"Required",
	}

	requiredFields := []string{
		"Items",
	}

	// Warns users if one or more of the unsupported fields is present
	for _, n := range unsupportedFields {
		f := v.FieldByName(n)
		if !f.IsNil() {
			logger.Logger.Warn(fmt.Sprintf("field %s in array attribute %s is not supported. this field will be ignored", n, name))
		}
	}
	// Returns an error if one or more of the required fields is missing
	for _, n := range requiredFields {
		f := v.FieldByName(n)
		if f.IsNil() {
			logger.Logger.Fatal(fmt.Sprintf("field %s in array attribute %s is required but missing", n, name))
		}
	}
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
	case "":
		logger.Logger.Fatal(fmt.Sprintf("attribute %s requires a type", name))
	default:
		logger.Logger.Fatal(fmt.Sprintf("unsupported type %s from attribute %s", t, name))
	}
}

func checkMetadataValidity(spec *CEGenSpec) {
	if len(spec.Metadata.PackageName) == 0 {
		logger.Logger.Warn("package name is required; using mypackage instead")
		spec.Metadata.PackageName = "mypackage"
	}

	r, _ := regexp.Compile(`^[A-Za-z][A-Za-z0-9\-\_]*$`)
	match := r.MatchString(spec.Metadata.PackageName)
	if !match {
		logger.Logger.Fatal("package name must start with an alphabetic character and uses only dash, underscore, and alphanumberic charactions")
	}

	if len(spec.Metadata.Version) == 0 {
		logger.Logger.Warn("package version is missing; using 0.0.1 instead")
		spec.Metadata.Version = "0.0.1"
	}
}
