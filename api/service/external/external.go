package external

import (
	"fmt"
	"github.com/gateway-fm/warp-external/api/templater"
)

type External struct {
	external         templater.Template
	ExternalFunction []*ExternalFunc
	configPath       string
}
type ExternalFunc func() (templater.ITemplate, error)

func (ex *External) ExternalFiller(extFunc []*ExternalFunc, path string, extTemp ...templater.Template) (*External, error) {
	ext := &External{
		configPath: path,
	}
	config, err := ext.external.DecodeConfig(ext.configPath)
	if err != nil {
		return nil, fmt.Errorf("decode config failed: %w", err)
	}
	ext.external.CfgSummon = config

	for i := range extTemp {
		ext.external = extTemp[i]
	}
	ext.ExternalFunction = extFunc
	return ext, nil
}
