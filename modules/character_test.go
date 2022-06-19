package modules_test

import (
	"github.com/lab42/rocketship/config"
	"github.com/lab42/rocketship/modules"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Character", func() {
	var cfg config.Config
	var mod modules.Module

	Context("Character module", func() {
		It("can initialize config", func() {
			cfg, _ = config.NewConfig()
		})
		It("can initialize module with config", func() {
			mod = *modules.NewModule(
				modules.CHARACTER_MODULE,
				modules.Character_module,
			)
		})
		It("can execute module with config", func() {
			out, err := mod.Exec(&cfg)
			Expect(err).To(BeNil())
			Expect(out).ShouldNot(BeEmpty())
		})
	})
})
