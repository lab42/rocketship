package modules

import (
	"fmt"
	"time"

	"github.com/lab42/rocketship/config"
	"github.com/lab42/rocketship/formatter"
)

const TIME_MODULE = "time"

func time_module(config *config.Config) (string, error) {
	if config.Time.Disabled {
		return "", nil
	}

	dateTime := time.Now()

	formatter := formatter.NewFormatter(
		TIME_MODULE,
		config.Time.Format,
		map[string]string{
			"time":   fmt.Sprintf("%d:%d:%d", dateTime.Hour(), dateTime.Minute(), dateTime.Second()),
			"symbol": config.Time.Symbol,
		},
	)

	return formatter.Render()
}

func init() {
	AddModule(NewModule(
		TIME_MODULE,
		time_module,
	))
}
