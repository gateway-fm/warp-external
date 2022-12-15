package api

import (
	"fmt"
	"github.com/gateway-fm/warp-external/api/service/external"

	"github.com/gateway-fm/warp-external/api/templater"
)

type TemplateFuncs struct {
	tps []templater.ITemplate
	External *external.External
}


// MergeExternal is
func (t *TemplateFuncs) MergeExternal() ([]templater.ITemplate, error) {
	t.External =

	t.tps = append(
		t.tps,
	)
	return t.tps, nil

}

// SummonNewInfra is
func SummonNewInfra() error {
	t := &TemplateFuncs{}
	tempsExt, err := t.MergeExternal()
	if err != nil {
		return fmt.Errorf("error while merging Infra templastes %w", err)
	}
	for _, tmp := range tempsExt {
		if !tmp.Excluded() {
			if err = tmp.GenerateNonGo(); err != nil {
				return fmt.Errorf(" generate  failed: %w", err)
			}
		}
	}
	return nil
}
