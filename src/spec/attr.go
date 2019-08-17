package spec

import (
	"fmt"
	"log"

	"github.com/michaelawyu/cloud-events-generator/src/utils"

	genspec "github.com/michaelawyu/cloud-events-generator/src/generator/spec"
)

// Attr is
type Attr struct {
	Type             string           `json:"type" yaml:"type"`
	Maximum          *float64         `json:"maximum" yaml:"maximum"`
	Minimum          *float64         `json:"minimum" yaml:"minimum"`
	ExclusiveMaximum *bool            `json:"exclusiveMaximum" yaml:"exclusiveMaxmium"`
	ExclusionMinimum *bool            `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	MaxLength        *int             `json:"maxLength" yaml:"maxLength"`
	MinLength        *int             `json:"minLength" yaml:"minLength"`
	Pattern          *string          `json:"pattern" yaml:"pattern"`
	MaxItems         *int             `json:"maxItems" yaml:"maxItems"`
	MinItems         *int             `json:"minItems" yaml:"minItems"`
	Required         *[]string        `json:"required" yaml:"required"`
	Enum             *[]interface{}   `json:"enum" yaml:"enum"`
	Items            *Attr            `json:"items" yaml:"items"`
	Properties       *map[string]Attr `json:"properties" yaml:"properties"`
	Description      *string          `json:"description" yaml:"description"`
	Format           *string          `json:"format" yaml:"format"`
	Default          *interface{}     `json:"default" yaml:"default"`
	Auto             *bool            `json:"auto" yaml:"auto"`
}

func (attr *Attr) parseAsObjAttr(name string) (genspec.VarSc, []genspec.Kls) {
	v := genspec.VarSc{}
	v.BaseName = name
	v.Name = utils.FormatName(name, "lowerCamel")
	klsName := utils.FormatName(name, "camel")
	v.DataType = klsName
	if attr.Description != nil {
		v.Description = *attr.Description
	}
	c := genspec.Kls{}
	c.KlsName = klsName
	dataKs := []genspec.Kls{}
	for k, v := range *attr.Properties {
		childV, childKs := v.parse(k)
		c.Vars = append(c.Vars, childV)
		dataKs = append(dataKs, childKs...)
	}
	if attr.Required != nil {
		for _, elem := range c.Vars {
			for _, n := range *attr.Required {
				if elem.Name == n {
					elem.Required = true
				}
			}
		}
	}
	return v, append(dataKs, c)
}

func (attr *Attr) parseAsNumAttr(name string) genspec.VarSc {
	v := genspec.VarSc{}
	v.BaseName = name
	v.Name = utils.FormatName(name, "lowerCamel")
	v.DataType = attr.Type
	if attr.Description != nil {
		v.Description = *attr.Description
	}
	if attr.Maximum != nil {
		v.HasValidation = true
		v.Maximum = string(fmt.Sprintf("%g", *attr.Maximum))
		if *attr.ExclusiveMaximum {
			v.ExclusiveMaximum = true
		}
	}
	if attr.Minimum != nil {
		v.HasValidation = true
		v.Minimum = string(fmt.Sprintf("%g", *attr.Minimum))
		if *attr.ExclusionMinimum {
			v.ExclusiveMinimum = true
		}
	}
	if attr.Default != nil {
		v.DefaultValue = getValueAsNum(*attr.Default)
	}
	if attr.Enum != nil {
		v.IsEnum = true
		v.AllowableValues = getEnumAsNums(*attr.Enum)
	}
	if attr.Format != nil {
		v.Format = *attr.Format
	}
	if attr.Auto != nil {
		v.Auto = genspec.GetAutoFormat(v.DataType, v.Format)
	}
	return v
}

func (attr *Attr) parseAsIntAttr(name string) genspec.VarSc {
	v := genspec.VarSc{}
	v.BaseName = name
	v.Name = utils.FormatName(name, "lowerCamel")
	v.DataType = attr.Type
	if attr.Description != nil {
		v.Description = *attr.Description
	}
	if attr.Maximum != nil {
		v.HasValidation = true
		v.Maximum = string(fmt.Sprintf("%g", *attr.Maximum))
		if *attr.ExclusiveMaximum {
			v.ExclusiveMaximum = true
		}
	}
	if attr.Minimum != nil {
		v.HasValidation = true
		v.Minimum = string(fmt.Sprintf("%g", *attr.Minimum))
		if *attr.ExclusionMinimum {
			v.ExclusiveMinimum = true
		}
	}
	if attr.Default != nil {
		v.DefaultValue = getValueAsInt(*attr.Default)
	}
	if attr.Enum != nil {
		v.IsEnum = true
		v.AllowableValues = getEnumAsInts(*attr.Enum)
	}
	if attr.Format != nil {
		v.Format = *attr.Format
	}
	if attr.Auto != nil {
		v.Auto = genspec.GetAutoFormat(v.DataType, v.Format)
	}
	return v
}

func (attr *Attr) parseAsBoolAttr(name string) genspec.VarSc {
	v := genspec.VarSc{}
	v.BaseName = name
	v.Name = utils.FormatName(name, "lowerCamel")
	v.DataType = attr.Type
	if attr.Description != nil {
		v.Description = *attr.Description
	}
	if attr.Default != nil {
		v.DefaultValue = getValueAsBool(*attr.Default)
	}
	return v
}

func (attr *Attr) parseAsStrAttr(name string) genspec.VarSc {
	v := genspec.VarSc{}
	v.BaseName = name
	v.Name = utils.FormatName(name, "lowerCamel")
	v.DataType = attr.Type
	if attr.Description != nil {
		v.Description = *attr.Description
	}
	if attr.MaxLength != nil {
		v.HasValidation = true
		v.MaxLength = fmt.Sprintf("%v", *attr.MaxLength)
	}
	if attr.MinLength != nil {
		v.HasValidation = true
		v.MinLength = fmt.Sprintf("%v", *attr.MinLength)
	}
	if attr.Pattern != nil {
		v.HasValidation = true
		v.Pattern = fmt.Sprintf("%s", *attr.Pattern)
	}
	if attr.Default != nil {
		v.DefaultValue = getValueAsStr(*attr.Default)
	}
	if attr.Enum != nil {
		v.IsEnum = true
		v.AllowableValues = getEnumAsStrs(*attr.Enum)
	}
	if attr.Format != nil {
		v.Format = *attr.Format
	}
	if attr.Auto != nil {
		v.Auto = genspec.GetAutoFormat(v.DataType, v.Format)
	}
	return v
}

// FIXME
func (attr *Attr) parseAsArrayAttr(name string) (genspec.VarSc, []genspec.Kls) {
	v := genspec.VarSc{}
	v.BaseName = name
	v.Name = utils.FormatName(name, "lowerCamel")
	if attr.Description != nil {
		v.Description = *attr.Description
	}

	if attr.MaxItems != nil {
		v.HasValidation = true
		v.MaxItems = fmt.Sprintf("%v", *attr.MaxItems)
	}
	if attr.MinItems != nil {
		v.HasValidation = true
		v.MinItems = fmt.Sprintf("%v", *attr.MinItems)
	}

	switch arrayItemType := attr.Items.Type; arrayItemType {
	case "object":
		itemKlsName := utils.FormatName(fmt.Sprintf("%sItem", name), "camel")
		v.DataType = formatArrayDataType(itemKlsName)
		_, cs := attr.Items.parseAsObjAttr(itemKlsName)
		return v, cs
	case "array":
		log.Fatalf("nested array is not supported in array attribute %s", name)
	case "number", "integer", "boolean", "string":
		v.DataType = formatArrayDataType(arrayItemType)
		if attr.Enum != nil {
			v.IsEnum = true
			switch arrayItemType {
			case "number":
				v.IsContainer = true
				v.IsListContainer = true
				v.AllowableValues = getEnumAsNums(*attr.Enum)
			case "integer":
				v.IsContainer = true
				v.IsListContainer = true
				v.AllowableValues = getEnumAsInts(*attr.Enum)
			case "string":
				v.IsContainer = true
				v.IsListContainer = true
				v.AllowableValues = getEnumAsStrs(*attr.Enum)
			default:
				log.Fatalf("unsupported type %s in the enum of array attribute %s", arrayItemType, name)
			}
		}
		return v, []genspec.Kls{}
	default:
		log.Fatalf("unsupported type %s in array attribute %s", arrayItemType, name)
	}

	return genspec.VarSc{}, []genspec.Kls{}
}

// parse is
func (attr *Attr) parse(name string) (genspec.VarSc, []genspec.Kls) {
	checkAttrSpecConformity(attr, name)

	switch attrType := attr.Type; attrType {
	case "object":
		v, cs := attr.parseAsObjAttr(name)
		return v, cs
	case "number":
		v := attr.parseAsNumAttr(name)
		return v, []genspec.Kls{}
	case "integer":
		v := attr.parseAsIntAttr(name)
		return v, []genspec.Kls{}
	case "boolean":
		v := attr.parseAsBoolAttr(name)
		return v, []genspec.Kls{}
	case "string":
		v := attr.parseAsStrAttr(name)
		return v, []genspec.Kls{}
	case "array":
		v, cs := attr.parseAsArrayAttr(name)
		return v, cs
	default:
		log.Fatalf("unsupported type %s from attribute %s", attrType, name)
	}

	return genspec.VarSc{}, []genspec.Kls{}
}
