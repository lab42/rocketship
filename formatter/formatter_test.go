package formatter_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/lab42/rocketship/formatter"
)

var _ = Describe("Formatter", func() {
	formatter := formatter.NewFormatter(
		"directory",
		"<fg=f44336>Hello</> {{.name}}",
		map[string]string{"name": "rocketship"},
	)

	result, err := formatter.Render()

	Context("with formatter instance", func() {
		It("name should be Directory", func() {
			Expect(formatter.Name).To(Equal("directory"))
		})
		It("output should be in color", func() {
			Expect(result).To(Equal("\x1b[38;2;244;67;54mHello\x1b[0m rocketship"))
		})
		It("there should be no error'", func() {
			Expect(err).To(BeNil())
		})
	})
})
