package modules

import (
	"io/ioutil"

	"github.com/lab42/rocketship/config"
	"github.com/lab42/rocketship/formatter"
	"golang.org/x/mod/modfile"
)

const GOLANG_MODULE = "golang"

func golang_module(config *config.Config) (string, error) {
	if config.Golang.Disabled {
		return "", nil
	}

	modFile, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return "", nil
	}

	f, err := modfile.Parse("go.mod", modFile, nil)
	if err != nil {
		return "", nil
	}

	formatter := formatter.NewFormatter(
		GOLANG_MODULE,
		config.Golang.Format,
		map[string]string{
			"version": f.Go.Version,
			"symbol":  config.Golang.Symbol,
		},
	)

	return formatter.Render()
}

func init() {
	AddModule(NewModule(
		GOLANG_MODULE,
		golang_module,
	))
}
