package modules_test

import (
	"github.com/lab42/rocketship/config"
	"github.com/lab42/rocketship/modules"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Git branch", func() {
	var cfg config.Config
	var mod modules.Module

	Context("Git branch module", func() {
		It("can initialize config", func() {
			cfg, _ = config.NewConfig()
		})
		It("can initialize module with config", func() {
			mod = *modules.NewModule(
				modules.GIT_BRANCH_MODULE,
				modules.Git_branch_module,
			)
		})
		It("can execute module with config", func() {
			_, err := mod.Exec(&cfg)
			Expect(err).To(BeNil())
		})
	})
})
