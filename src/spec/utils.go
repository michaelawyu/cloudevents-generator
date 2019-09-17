package spec

import (
	"fmt"
	"reflect"
	"strconv"

	genspec "github.com/michaelawyu/cloudevents-generator/src/generator/spec"
	"github.com/michaelawyu/cloudevents-generator/src/logger"
)

func getValueAsNum(d interface{}) string {
	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	switch k := t.Kind(); k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := v.Int()
		return fmt.Sprintf("%v", i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := v.Uint()
		return fmt.Sprintf("%v", u)
	case reflect.Float32, reflect.Float64:
		f := v.Float()
		return fmt.Sprintf("%v", f)
	case reflect.String:
		s := v.String()
		i, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			return fmt.Sprintf("%v", i)
		}
		u, err := strconv.ParseUint(s, 10, 64)
		if err == nil {
			return fmt.Sprintf("%v", u)
		}
		f, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return fmt.Sprintf("%g", f)
		}
		logger.Logger.Fatal(fmt.Sprintf("default or enum value %s is not a number (looks like a string)", v))
	default:
		logger.Logger.Fatal(fmt.Sprintf("default or enum value %s is not a number", v))
	}
	return ""
}

func getEnumAsNums(es []interface{}) []genspec.AllowableValue {
	as := []genspec.AllowableValue{}
	for _, e := range es {
		a := genspec.AllowableValue{
			Value: getValueAsNum(e),
		}
		as = append(as, a)
	}
	return as
}

func getValueAsInt(d interface{}) string {
	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	switch k := t.Kind(); k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := v.Int()
		return fmt.Sprintf("%v", i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := v.Uint()
		return fmt.Sprintf("%v", u)
	case reflect.String:
		s := v.String()
		i, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			return fmt.Sprintf("%v", i)
		}
		u, err := strconv.ParseUint(s, 10, 64)
		if err == nil {
			return fmt.Sprintf("%v", u)
		}
		logger.Logger.Fatal(fmt.Sprintf("default or enum value %s is not a number (looks like a string)", v))
	default:
		logger.Logger.Fatal(fmt.Sprintf("default or enum value %s is not a number", v))
	}
	return ""
}

func getEnumAsInts(es []interface{}) []genspec.AllowableValue {
	as := []genspec.AllowableValue{}
	for _, e := range es {
		a := genspec.AllowableValue{
			Value: getValueAsInt(e),
		}
		as = append(as, a)
	}
	return as
}

func getValueAsBool(d interface{}) string {
	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	switch k := t.Kind(); k {
	case reflect.Bool:
		b := v.Bool()
		return fmt.Sprintf("%t", b)
	case reflect.String:
		s := v.String()
		b, err := strconv.ParseBool(s)
		if err == nil {
			return fmt.Sprintf("%t", b)
		}
		logger.Logger.Fatal(fmt.Sprintf("default or enum value %s is not boolean (seems like a string)", v))
	default:
		logger.Logger.Fatal(fmt.Sprintf("default or enum value %s is not boolean", v))
	}
	return ""
}

func getValueAsStr(d interface{}) string {
	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	switch k := t.Kind(); k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := v.Int()
		return fmt.Sprintf("\"%v\"", i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := v.Uint()
		return fmt.Sprintf("\"%v\"", u)
	case reflect.Float32, reflect.Float64:
		f := v.Float()
		return fmt.Sprintf("\"%v\"", f)
	case reflect.Bool:
		b := v.Bool()
		return fmt.Sprintf("\"%t\"", b)
	case reflect.String:
		s := v.String()
		return fmt.Sprintf("\"%s\"", s)
	default:
		logger.Logger.Fatal(fmt.Sprintf("default or enum value %s is not a string", v))
	}
	return ""
}

func getEnumAsStrs(es []interface{}) []genspec.AllowableValue {
	as := []genspec.AllowableValue{}
	for _, e := range es {
		a := genspec.AllowableValue{
			Value: getValueAsStr(e),
		}
		as = append(as, a)
	}
	return as
}

func formatArrayDataType(itemKlsName string) string {
	return fmt.Sprintf("array/%s", itemKlsName)
}
