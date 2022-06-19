package modules

import (
	"github.com/lab42/rocketship/config"
	"github.com/lab42/rocketship/formatter"
)

const CHARACTER_MODULE = "character"

func Character_module(config *config.Config) (string, error) {
	if config.Character.Disabled {
		return "", nil
	}

	formatter := formatter.NewFormatter(
		CHARACTER_MODULE,
		config.Character.Format,
		map[string]string{},
	)

	return formatter.Render()
}

func init() {
	AddModule(NewModule(
		CHARACTER_MODULE,
		Character_module,
	))
}
