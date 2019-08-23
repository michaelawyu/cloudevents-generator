package nodejs

import (
	"fmt"
	"os"
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/michaelawyu/cloud-events-generator/src/logger"
	"github.com/michaelawyu/cloud-events-generator/src/utils"

	genspec "github.com/michaelawyu/cloud-events-generator/src/generator/spec"
)

const prefix = "/nodejs"

func matchDataType(t string) (string, bool, string) {
	p := fmt.Sprintf("%s/%s", prefix, "typing.mustache")
	tpl := utils.GetTemplate(p)

	tcs := strings.Split(t, "/")
	if tcs[0] == "array" && len(tcs) > 1 {
		it, btf, _ := matchDataType(tcs[1])
		return fmt.Sprintf("Array(%s)", it), btf, it
	}

	s := map[string]bool{}
	switch i := tcs[0]; i {
	case "string":
		s["IsString"] = true
	case "number":
		s["IsNumber"] = true
	case "integer":
		s["IsInteger"] = true
	case "boolean":
		s["IsBoolean"] = true
	default:
		return i, false, ""
	}
	jst, err := mustache.Render(tpl, s)
	if err != nil {
		logger.Logger.Fatal(fmt.Sprintf("unsupported type %s", t))
	}
	return jst, true, ""
}

func genKls(k genspec.Kls) string {
	deps := []map[string]string{}
	for i := range k.Vars {
		var builtInTypeFlag bool
		var itemType string
		// Match data types with their node.js/javascript counterparts
		k.Vars[i].DataType, builtInTypeFlag, itemType = matchDataType(k.Vars[i].DataType)
		// If not a built-in type, import the class separately
		if !builtInTypeFlag {
			if itemType == "" {
				deps = append(deps, map[string]string{
					"KlsName": k.Vars[i].DataType,
				})
			} else {
				deps = append(deps, map[string]string{
					"KlsName": itemType,
				})
			}
		}
		// Set HasMore flags
		if i != len(k.Vars)-1 {
			k.Vars[i].HasMore = true
		}
		for ai := range k.Vars[i].AllowableValues {
			if ai != len(k.Vars[i].AllowableValues)-1 {
				k.Vars[i].AllowableValues[ai].HasMore = true
			}
		}
	}

	p := fmt.Sprintf("%s/%s", prefix, "model.mustache")
	t := utils.GetTemplate(p)

	d, err := mustache.Render(t, map[string]interface{}{
		"Model":   k,
		"Imports": deps,
	})
	if err != nil {
		logger.Logger.Fatal(fmt.Sprintf("failed to generate model from template: %s", err))
	}
	return d
}

func genFile(tp string, p string, fn string, context map[string]interface{}) {
	t := utils.GetTemplate(tp)
	d, err := mustache.Render(t, context)
	if err != nil {
		logger.Logger.Fatal(fmt.Sprintf("failed to render template %s to %s: %s", tp, fn, err))
	}
	fp := fmt.Sprintf("%s/%s", p, fn)
	utils.WriteFile(fp, d)
}

func genMod(p string, m genspec.Mod) {
	n := m.ModName
	logger.Logger.Info(fmt.Sprintf("preparing event module %s", n))

	// Create the folder
	dp := fmt.Sprintf("%s/%s", p, n)
	err := os.MkdirAll(dp, os.FileMode(0777))
	if err != nil {
		logger.Logger.Fatal(fmt.Sprintf("cannot create folder for mod %s: %s", n, err))
	}
	p = p + fmt.Sprintf("/%s", n)

	// Generate the event class
	logger.Logger.Info(fmt.Sprintf("preparing event class %s.js", m.Event.KlsName))
	d := genKls(m.Event)
	fp := fmt.Sprintf("%s/%s.js", p, m.Event.KlsName)
	utils.WriteFile(fp, d)

	// Generate the data class(es)
	for _, v := range m.DataClasses {
		logger.Logger.Info(fmt.Sprintf("preparing data class %s.js", v.KlsName))
		d = genKls(v)
		fp = fmt.Sprintf("%s/%s.js", p, v.KlsName)
		utils.WriteFile(fp, d)
	}
}

// GenPkg is
func GenPkg(p string, ms []genspec.Mod, b genspec.BindSelector, meta genspec.Metadata) {
	// Prepare package files
	logger.Logger.Info("preparing package")
	pkgName := utils.FormatName(meta.PackageName, "snake")
	p = fmt.Sprintf("%s/%s", p, pkgName)
	err := os.MkdirAll(p, os.FileMode(0777))
	if err != nil {
		logger.Logger.Fatal(fmt.Sprintf("cannot create folder %s at %s", pkgName, p))
	}

	// Add README.md
	logger.Logger.Info("preparing README.md")
	tp := fmt.Sprintf("%s/%s", prefix, "README.md")
	genFile(tp, p, "README.md", map[string]interface{}{})

	// Generate Base.js
	logger.Logger.Info("preparing Base.js")
	tp = fmt.Sprintf("%s/%s", prefix, "base.mustache")
	genFile(tp, p, "Base.js", map[string]interface{}{
		"Binding": b,
	})

	// Generate index.js
	logger.Logger.Info("preparing index.js")
	tp = fmt.Sprintf("%s/%s", prefix, "index.mustache")
	genFile(tp, p, "index.js", map[string]interface{}{
		"Mods": ms,
	})

	// Generate package.json
	logger.Logger.Info("preparing package.json")
	tp = fmt.Sprintf("%s/%s", prefix, "package.mustache")
	genFile(tp, p, "package.json", map[string]interface{}{
		"Metadata": meta,
	})

	// Generate the mods
	for _, m := range ms {
		genMod(p, m)
	}
}
