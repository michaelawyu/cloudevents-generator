package spec

import (
	"fmt"

	genspec "github.com/michaelawyu/cloudevents-generator/src/generator/spec"
	"github.com/michaelawyu/cloudevents-generator/src/logger"
	"github.com/michaelawyu/cloudevents-generator/src/utils"
)

// Event is
type Event struct {
	Attributes map[string]Attr `json:"attributes" yaml:"attributes"`
	Required   []string        `json:"required" yaml:"required"`
}

// parse is
func (event *Event) parse(name string) (genspec.Kls, []genspec.Kls) {
	logger.Logger.Info(fmt.Sprintf("parsing event %s", name))
	if event.Attributes == nil {
		logger.Logger.Fatal(fmt.Sprintf("attributes of event %s are missing", name))
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
	for i := range vs {
		for _, n := range event.Required {
			if vs[i].Name == utils.FormatName(n, "lowerCamel") {
				vs[i].Required = true
			}
		}
	}
	return genspec.Kls{
		KlsName: utils.FormatName(name, "camel"),
		Vars:    vs,
	}, childCs
}
