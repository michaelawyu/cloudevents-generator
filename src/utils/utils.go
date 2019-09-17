package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/michaelawyu/cloudevents-generator/src/logger"
	"github.com/michaelawyu/cloudevents-generator/src/vfsgen"
)

// FormatName is
func FormatName(name string, style string) string {
	switch style {
	case "snake":
		return strcase.ToSnake(name)
	case "screamingSnake":
		return strcase.ToScreamingSnake(name)
	case "kebab":
		return strcase.ToKebab(name)
	case "screamingKebab":
		return strcase.ToScreamingKebab(name)
	case "camel":
		return strcase.ToCamel(name)
	case "lowerCamel":
		return strcase.ToLowerCamel(name)
	default:
		logger.Logger.Warn(fmt.Sprintf("unsupported naming style %s", style))
		return name
	}
}

// FormatPath is
func FormatPath(p string) string {
	if strings.HasSuffix(p, "/") {
		return strings.TrimSuffix(p, "/")
	}

	return p
}

var fs http.FileSystem = vfsgen.Assets

// GetTemplate is
func GetTemplate(p string) string {
	f, err := fs.Open(p)
	if err != nil {
		logger.Logger.Fatal(fmt.Sprintf("template %s is missing", p))
	}
	defer f.Close()
	info, _ := f.Stat()
	s := info.Size()
	t := make([]byte, s)
	_, err = f.Read(t)
	if err != nil && err != io.EOF {
		logger.Logger.Fatal(fmt.Sprintf("cannot read template %s: %s", p, err))
	}
	return string(t)
}

// WriteFile is
func WriteFile(p string, d string) {
	err := ioutil.WriteFile(p, []byte(d), 0777)
	if err != nil {
		logger.Logger.Fatal(fmt.Sprintf("cannot write file %s", p))
	}
}
