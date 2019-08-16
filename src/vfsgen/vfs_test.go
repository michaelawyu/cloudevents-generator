package vfsgen

import (
	"net/http"
	"testing"

	"github.com/shurcooL/vfsgen"
)

func TestPackagingStaticAssets(t *testing.T) {
	var fs http.FileSystem = http.Dir("../templates")
	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "vfsgen",
		VariableName: "Assets",
	})
	if err != nil {
		t.Fatal(err)
	}
}
