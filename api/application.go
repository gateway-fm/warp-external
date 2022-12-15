package api

import (
	"fmt"
	"html/template"

	"github.com/gateway-fm/warp-external/api/templater"
)

type ITemplateFunc interface {
	AddNewTemplate(temp *templater.Template)
	GetExternals() ([]templater.ITemplate, error)
	NewTemplate(elems []string,
		ifaces []interface{},
		cfgPath, configTemplatePath, outPutFilePath string,
		funcMap template.FuncMap) (*templater.Template, error)
}
type TemplateFuncs struct {
	tps  []templater.ITemplate
	temp templater.Template
}

func (t *TemplateFuncs) GetExternals() ([]templater.ITemplate, error) {
	return t.tps, nil
}
func (t *TemplateFuncs) AddNewTemplate(temp *templater.Template) {
	t.tps = append(t.tps, temp)
}

func (t *TemplateFuncs) NewTemplate(elems []string, ifaces []interface{}, cfgPath, configTemplatePath, outPutFilePath string, funcMap template.FuncMap) (*templater.Template, error) {
	t.temp = templater.Template{
		Elems:              elems,
		Ifaces:             ifaces,
		CfgPath:            cfgPath,
		ConfigTemplatePath: configTemplatePath,
		OutPutFilePath:     outPutFilePath,
		FuncMap:            funcMap,
	}
	var err error
	t.temp.CfgSummon, err = t.temp.DecodeConfig()
	if err != nil {
		return nil, fmt.Errorf("error while decoding config:%w", err)
	}
	return &t.temp, nil
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
