package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/iancoleman/strcase"
	"github.com/michaelawyu/cloud-events-generator/src/vfsgen"
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
		fmt.Printf("unsupported naming style %s\n", style)
		return name
	}
}

var fs http.FileSystem = vfsgen.Assets

// GetTemplate is
func GetTemplate(p string) string {
	f, err := fs.Open(p)
	if err != nil {
		log.Fatalf("template %s is missing", p)
	}
	defer f.Close()
	info, _ := f.Stat()
	s := info.Size()
	t := make([]byte, s)
	_, err = f.Read(t)
	if err != nil && err != io.EOF {
		log.Fatalf("cannot read template %s: %s", p, err)
	}
	return string(t)
}

// WriteFile is
func WriteFile(p string, d string) {
	err := ioutil.WriteFile(p, []byte(d), 0777)
	if err != nil {
		log.Fatalf("cannot write file %s", p)
	}
}
