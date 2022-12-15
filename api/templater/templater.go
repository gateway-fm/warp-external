package templater

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"go.uber.org/zap"

	"github.com/gateway-fm/scriptorium/logger"
)

type ITemplate interface {
	GenerateNonGo() error
}

type WarpConfig struct {
	CfgSumm *any
}

type Template struct {
	CfgSummon                                   *WarpConfig
	Elems                                       []string
	Ifaces                                      []interface{}
	CfgPath, ConfigTemplatePath, OutPutFilePath string
	FuncMap                                     template.FuncMap
	IsExcluded                                  bool
}

func (w *WarpConfig) CreateWarpConfig(config any) *any {
	config = w.CfgSumm
	return w.CfgSumm
}
func (t *Template) DecodeConfig() (*WarpConfig, error) {

	err := hclsimple.DecodeFile(t.CfgPath, nil, &t.CfgSummon)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	return t.CfgSummon, nil
}

// GenerateFile is generates single *.go file in project's dir
// according to values parsed from .hcl file
// .hcl -> templater.go -> *.gotmpl -> *.go
func (t *Template) GenerateFile() error {

	elems := t.Elems
	ifaces := t.Ifaces
	// as an empty .go file and just "filled up" in this func
	file, _ := os.Create(t.OutPutFilePath) //
	defer file.Close()

	// path to template file is absolute here, but it doesn't make any sense :D
	pattern, _ := filepath.Abs(t.ConfigTemplatePath) // .gotmpl is used because of IDE's supports only :D

	// template final preparation. Template must parse given pattern (which is our scheme.gotmpl file)
	tmpl := template.Must(template.New("").Funcs(t.FuncMap).ParseFiles(pattern))

	var wg sync.WaitGroup
	for i := range elems {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := tmpl.ExecuteTemplate(file, elems[i], ifaces[i]) // first arg is output, second is the data we want to pass to this config. It could also be nil.
			if err != nil {
				logger.Log().Error("An error occurred", zap.Error(err))
				return
			}
		}()
		wg.Wait()
	}

	return nil
}

// GenerateNonGo is main generation function for non .go files
func (t *Template) GenerateNonGo() error {
	err := t.GenerateFile()
	if err != nil {
		return fmt.Errorf(" GenerateScheme returned an error: %w", err)
	}
	return nil
}
