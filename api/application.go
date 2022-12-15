package api

import (
	"fmt"
	"html/template"

	"github.com/gateway-fm/warp-external/api/templater"
)

type TemplateFuncs struct {
	tps []templater.ITemplate
}

func (t *TemplateFuncs) GetExternals() ([]templater.ITemplate, error) {
	return t.tps, nil
}
func (t *TemplateFuncs) AddNewTemplate(temp templater.ITemplate) {
	t.tps = append(t.tps, temp)
}

func NewTemplate(elems []string, ifaces []interface{}, configTemplatePath, outPutFilePath string, funcMap template.FuncMap) (templater.ITemplate, error) {
	temp := templater.Template{
		Elems:              elems,
		Ifaces:             ifaces,
		ConfigTemplatePath: configTemplatePath,
		OutPutFilePath:     outPutFilePath,
		FuncMap:            funcMap,
	}
	return &temp, nil
}

// SummonExternal is
func (t *TemplateFuncs) SummonExternal() error {
	tempsExt, err := t.GetExternals()
	if err != nil {
		return fmt.Errorf("error while getting external templates %w", err)
	}
	for _, tmp := range tempsExt {
		if err = tmp.GenerateNonGo(); err != nil {
			return fmt.Errorf(" generate  failed: %w", err)
		}
	}
	return nil
}
