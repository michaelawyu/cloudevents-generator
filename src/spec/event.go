package spec

import (
	"log"

	genspec "github.com/michaelawyu/cloud-events-generator/src/generator/spec"
	"github.com/michaelawyu/cloud-events-generator/src/utils"
)

// Event is
type Event struct {
	Attributes map[string]Attr `json:"attributes" yaml:"attributes"`
	Required   []string        `json:"required" yaml:"required"`
}

// parse is
func (event *Event) parse(name string) (genspec.Kls, []genspec.Kls) {
	if event.Attributes == nil {
		log.Fatalf("Attributes of event %s are missing", name)
	}
	checkReqAttrs(event, name)
	checkEventSpecConformity(event, name)

	var vs []genspec.VarSc
	var childCs []genspec.Kls
	for n, attr := range event.Attributes {
		v, cs := attr.parse(n)
		vs = append(vs, v)
		childCs = append(childCs, cs...)
	}
	for _, v := range vs {
		for _, n := range event.Required {
			if v.Name == utils.FormatName(n, "lowerCamel") {
				v.Required = true
			}
		}
	}
	return genspec.Kls{
		KlsName: utils.FormatName(name, "camel"),
		Vars:    vs,
	}, childCs
}
