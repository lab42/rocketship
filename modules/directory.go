package modules

import (
	"os"

	"github.com/lab42/rocketship/config"
	"github.com/lab42/rocketship/formatter"
)

const DIRECTORY_MODULE = "directory"

func directory(config *config.Config) (string, error) {
	if config.Directory.Disabled {
		return "", nil
	}

	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	formatter := formatter.NewFormatter(
		DIRECTORY_MODULE,
		config.Directory.Format,
		map[string]string{
			"path":   path,
			"symbol": config.Directory.Symbol,
		},
	)

	return formatter.Render()
}

func init() {
	AddModule(NewModule(
		DIRECTORY_MODULE,
		directory,
	))
}
