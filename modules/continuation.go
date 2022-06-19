package modules

import (
	"github.com/lab42/rocketship/config"
	"github.com/lab42/rocketship/formatter"
)

const CONTINUATION_MODULE = "continuation"

func continuation_module(config *config.Config) (string, error) {
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
		continuation_module,
	))
}
