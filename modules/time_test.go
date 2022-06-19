package modules_test

import (
	"github.com/lab42/rocketship/config"
	"github.com/lab42/rocketship/modules"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Time", func() {
	var cfg config.Config
	var mod modules.Module

	Context("Time module", func() {
		It("can initialize config", func() {
			cfg, _ = config.NewConfig()
		})
		It("can initialize module with config", func() {
			mod = *modules.NewModule(
				modules.TIME_MODULE,
				modules.Time_module,
			)
		})
		It("can execute module with config", func() {
			out, err := mod.Exec(&cfg)
			Expect(err).To(BeNil())
			Expect(out).ShouldNot(BeEmpty())
		})
	})
})
