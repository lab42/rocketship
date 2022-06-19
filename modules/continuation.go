package modules

import (
	"github.com/lab42/rocketship/config"
	"github.com/lab42/rocketship/formatter"
)

const CONTINUATION_MODULE = "continuation"

func Continuation_module(config *config.Config) (string, error) {
	if config.Continuation.Disabled {
		return "", nil
	}

	formatter := formatter.NewFormatter(
		CONTINUATION_MODULE,
		config.Continuation.Format,
		map[string]string{},
	)

	return formatter.Render()
}

func init() {
	AddModule(NewModule(
		CONTINUATION_MODULE,
		Continuation_module,
	))
}
