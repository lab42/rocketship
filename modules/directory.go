package modules

import (
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/lab42/rocketship/config"
	"github.com/lab42/rocketship/formatter"
)

const DIRECTORY_MODULE = "directory"

func directory_module(config *config.Config) (string, error) {
	if config.Directory.Disabled {
		return "", nil
	}

	if _, err := git.PlainOpen("."); err == nil {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}

		pathSegments := strings.Split(dir, string(os.PathSeparator))

		formatter := formatter.NewFormatter(
			DIRECTORY_MODULE,
			config.Directory.Format,
			map[string]string{
				"path":   pathSegments[len(pathSegments)-1],
				"symbol": config.Directory.Symbol,
			},
		)

		return formatter.Render()
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	formatter := formatter.NewFormatter(
		DIRECTORY_MODULE,
		config.Directory.Format,
		map[string]string{
			"path":   strings.Replace(path, homeDir, "~", 1),
			"symbol": config.Directory.Symbol,
		},
	)

	return formatter.Render()
}

func init() {
	AddModule(NewModule(
		DIRECTORY_MODULE,
		directory_module,
	))
}
