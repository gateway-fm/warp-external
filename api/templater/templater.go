package templater

import (
	"fmt"
	"github.com/gateway-fm/warp-external/config"
	"go/format"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"go.uber.org/zap"

	"github.com/gateway-fm/scriptorium/logger"
)

type ITemplate interface {
	Generate() error
	Excluded() bool
	GenerateNonGo() error
}

type Template struct {
	CfgSummon                          *config.SummonConfig //config file that holds config's related service info
	Elems                              []string
	Ifaces                             []interface{}
	ConfigTemplatePath, OutPutFilePath string
	FuncMap                            template.FuncMap
	IsExcluded                         bool
}

func (t *Template) DecodeConfig(cfgPath string) (*config.SummonConfig, error) {
	err := hclsimple.DecodeFile(cfgPath, nil, &t.CfgSummon)
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
func (t *Template) Excluded() bool {
	return t.IsExcluded
}

// GoFmt execute go fmt for specific .go file from the code
func GoFmt(path string) error {

	read, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fmted, err := format.Source(read)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, fmted, 0666)
	if err != nil {
		return err
	}
	return nil
}

// Generate is main generation function
func (t *Template) Generate() error {
	err := t.GenerateFile()
	if err != nil {
		return fmt.Errorf(" GenerateScheme returned an error: %w", err)
	}
	err = GoFmt(t.OutPutFilePath)
	if err != nil {
		return fmt.Errorf(" go fmt returned an error: %w", err)
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
