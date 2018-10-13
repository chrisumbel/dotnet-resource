package out_test

import (
	"github.com/miclip/dotnet-resource/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/miclip/dotnet-resource/out"
	"github.com/miclip/dotnet-resource"
)

var _ = Describe("out", func() {
	BeforeEach(func(){
		dotnetresource.ExecCommand = fakes.FakeExecCommand
	})

	It("should output an empty JSON list", func() {
		output, err := out.Execute()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(output).Should(MatchJSON("[]"))
	})
})