package python

import (
	"fmt"
	"log"
	"os"
	"strings"

	utils "github.com/michaelawyu/cloud-events-generator/src/utils"

	"github.com/cbroglie/mustache"

	genspec "github.com/michaelawyu/cloud-events-generator/src/generator/spec"
)

const prefix = "/python"

func matchDataType(t string) (string, bool) {
	p := fmt.Sprintf("%s/%s", prefix, "typing.mustache")
	tpl := utils.GetTemplate(p)

	tcs := strings.Split(t, "/")
	if tcs[0] == "array" && len(tcs) > 1 {
		it, _ := matchDataType(tcs[1])
		return fmt.Sprintf("List[%s]", it), true
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
		return i, false
	}
	pyt, err := mustache.Render(tpl, s)
	if err != nil {
		log.Fatalf("unsupported type %s", t)
	}
	return pyt, true
}

func genKls(k genspec.Kls) string {
	deps := []map[string]string{}
	for i := range k.Vars {
		var builtInTypeFlag bool
		// Match data types with their python counterparts
		k.Vars[i].DataType, builtInTypeFlag = matchDataType(k.Vars[i].DataType)
		// If not a built-in type, imports the class separately
		if !builtInTypeFlag {
			deps = append(deps, map[string]string{
				"KlsName": k.Vars[i].DataType,
			})
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
		log.Fatalf("failed to generate model from template: %s", err)
	}
	return d
}

func genFile(tp string, p string, fn string, context map[string]interface{}) {
	t := utils.GetTemplate(tp)
	d, err := mustache.Render(t, context)
	if err != nil {
		log.Fatalf("failed to render template %s to %s: %s", tp, fn, err)
	}
	fp := fmt.Sprintf("%s/%s", p, fn)
	utils.WriteFile(fp, d)
}

func genMod(p string, m genspec.Mod) {
	n := m.ModName

	// Create the folder
	dp := fmt.Sprintf("%s/%s", p, n)
	err := os.MkdirAll(dp, os.FileMode(0777))
	if err != nil {
		log.Fatalf("cannot create folder for mod %s: %s", n, err)
	}
	p = p + fmt.Sprintf("/%s", n)

	// Generate __init__.py
	kls := []map[string]string{}
	kls = append(kls, map[string]string{
		"KlsName": m.Event.KlsName,
	})
	for _, v := range m.DataClasses {
		kls = append(kls, map[string]string{
			"KlsName": v.KlsName,
		})
	}

	tp := fmt.Sprintf("%s/%s", prefix, "__init__mod.mustache")
	genFile(tp, p, "__init__.py", map[string]interface{}{
		"Kls": kls,
	})

	// Generate the event class
	d := genKls(m.Event)
	fp := fmt.Sprintf("%s/%s.py", p, m.Event.KlsName)
	utils.WriteFile(fp, d)

	// Generate the data class(es)
	for _, v := range m.DataClasses {
		d = genKls(v)
		fp = fmt.Sprintf("%s/%s.py", p, v.KlsName)
		utils.WriteFile(fp, d)
	}
}

// GenPkg is
func GenPkg(p string, ms []genspec.Mod, b genspec.BindSelector, meta genspec.Metadata) {
	// Prepare files for distribution purposes
	// Generate setup.py
	tp := fmt.Sprintf("%s/%s", prefix, "setup.mustache")
	genFile(tp, p, "setup.py", map[string]interface{}{
		"Metadata": meta,
	})

	// Generate requirements.txt
	tp = fmt.Sprintf("%s/%s", prefix, "requirements.mustache")
	genFile(tp, p, "requirements.txt", map[string]interface{}{
		"Deps": b,
	})

	// Add README.md
	tp = fmt.Sprintf("%s/%s", prefix, "README.md")
	genFile(tp, p, "README.md", map[string]interface{}{})

	// Prepare package files
	pkgName := utils.FormatName(meta.PackageName, "snake")
	p = fmt.Sprintf("%s/%s", p, pkgName)
	err := os.MkdirAll(p, os.FileMode(0777))
	if err != nil {
		log.Fatalf("cannot create package folder %s (%s)", pkgName, p)
	}

	// Generate __init__.py
	mods := []map[string]string{}
	for _, m := range ms {
		mods = append(mods, map[string]string{
			"ModName":      m.ModName,
			"EventKlsName": m.Event.KlsName,
		})
	}

	tp = fmt.Sprintf("%s/%s", prefix, "__init__.mustache")
	genFile(tp, p, "__init__.py", map[string]interface{}{
		"Mods": mods,
	})

	// Generate util.py
	tp = fmt.Sprintf("%s/%s", prefix, "util.mustache")
	genFile(tp, p, "util.py", map[string]interface{}{})

	// Generate base_model.py
	tp = fmt.Sprintf("%s/%s", prefix, "base_model.mustache")
	genFile(tp, p, "base_model.py", map[string]interface{}{
		"Metadata": meta,
		"Binding":  b,
	})

	// Generate the mods
	for _, m := range ms {
		genMod(p, m)
	}
}
