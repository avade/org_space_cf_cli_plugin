package main_test

import (
	. "github.com/avade/org_space_cf_cli_plugin/main"
	"github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CallOrgSpace", func() {
	Describe(".Run", func() {
		var fakeCliConnection *fakes.FakeCliConnection
		var callCliCommandPlugin *OrgSpace

		BeforeEach(func() {
			fakeCliConnection = &fakes.FakeCliConnection{}
			callCliCommandPlugin = &OrgSpace{}
		})

		Context("when one arguments is provided", func() {
			It("outputs an error message for command usage", func() {
				output := io_helpers.CaptureOutput(func() {
					callCliCommandPlugin.Run(fakeCliConnection, []string{"org-space", "arg1"})
				})
				Expect(len(output)).To(Equal(1))
				Expect(output[0]).To(Equal("You must provide an org and space i.e. org-space org space"))
			})
		})

		Context("when three arguments are provided", func() {
			It("outputs an error message for command usage", func() {
				output := io_helpers.CaptureOutput(func() {
					callCliCommandPlugin.Run(fakeCliConnection, []string{"org-space", "arg1", "arg2", "arg3"})
				})
				Expect(len(output)).To(Equal(1))
				Expect(output[0]).To(Equal("You must provide an org and space i.e. org-space org space"))
			})
		})

		Context("when the org and space are passed in", func() {
			It("ouputs the text returned by the org-space command", func() {
				fakeCliConnection.CliCommandReturns([]string{"Hi", "Mom"}, nil)
				output := io_helpers.CaptureOutput(func() {
					callCliCommandPlugin.Run(fakeCliConnection, []string{"org-space", "testorg", "testspace"})
				})

				Expect(output[0]).To(Equal("Created org"))
			})

			Context("when the org does not exist", func() {
				It("call target org followed by create org, passing the org name argument", func() {
					io_helpers.CaptureOutput(func() {
						callCliCommandPlugin.Run(fakeCliConnection, []string{"org-space", "testorg", "testspace"})
					})
					Expect(fakeCliConnection.CliCommandCallCount()).To(Equal(6))

					Expect(len(fakeCliConnection.CliCommandArgsForCall(0))).To(Equal(3))
					Expect(fakeCliConnection.CliCommandArgsForCall(0)[0]).To(Equal("target"))
					Expect(fakeCliConnection.CliCommandArgsForCall(0)[1]).To(Equal("-o"))
					Expect(fakeCliConnection.CliCommandArgsForCall(0)[2]).To(Equal("testorg"))

					Expect(len(fakeCliConnection.CliCommandArgsForCall(1))).To(Equal(2))
					Expect(fakeCliConnection.CliCommandArgsForCall(1)[0]).To(Equal("create-org"))
					Expect(fakeCliConnection.CliCommandArgsForCall(1)[1]).To(Equal("testorg"))

					Expect(len(fakeCliConnection.CliCommandArgsForCall(2))).To(Equal(3))
					Expect(fakeCliConnection.CliCommandArgsForCall(2)[0]).To(Equal("target"))
					Expect(fakeCliConnection.CliCommandArgsForCall(2)[1]).To(Equal("-o"))
					Expect(fakeCliConnection.CliCommandArgsForCall(2)[2]).To(Equal("testorg"))

					Expect(len(fakeCliConnection.CliCommandArgsForCall(3))).To(Equal(5))
					Expect(fakeCliConnection.CliCommandArgsForCall(3)[0]).To(Equal("target"))
					Expect(fakeCliConnection.CliCommandArgsForCall(3)[1]).To(Equal("-o"))
					Expect(fakeCliConnection.CliCommandArgsForCall(3)[2]).To(Equal("testorg"))
					Expect(fakeCliConnection.CliCommandArgsForCall(3)[3]).To(Equal("-s"))
					Expect(fakeCliConnection.CliCommandArgsForCall(3)[4]).To(Equal("testspace"))

					Expect(len(fakeCliConnection.CliCommandArgsForCall(4))).To(Equal(4))
					Expect(fakeCliConnection.CliCommandArgsForCall(4)[0]).To(Equal("create-space"))
					Expect(fakeCliConnection.CliCommandArgsForCall(4)[1]).To(Equal("testspace"))
					Expect(fakeCliConnection.CliCommandArgsForCall(4)[2]).To(Equal("-o"))
					Expect(fakeCliConnection.CliCommandArgsForCall(4)[3]).To(Equal("testorg"))

					Expect(len(fakeCliConnection.CliCommandArgsForCall(5))).To(Equal(5))
					Expect(fakeCliConnection.CliCommandArgsForCall(5)[0]).To(Equal("target"))
					Expect(fakeCliConnection.CliCommandArgsForCall(5)[1]).To(Equal("-o"))
					Expect(fakeCliConnection.CliCommandArgsForCall(5)[2]).To(Equal("testorg"))
					Expect(fakeCliConnection.CliCommandArgsForCall(5)[3]).To(Equal("-s"))
					Expect(fakeCliConnection.CliCommandArgsForCall(5)[4]).To(Equal("testspace"))
				})
			})
		})
	})
})
