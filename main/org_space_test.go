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

				Expect(output[0]).To(Equal("Org testorg and Space testspace is now available and targeted"))
			})

			Context("when the org and space does not exist", func() {
				It("calls the correct commands in sequence", func() {
					io_helpers.CaptureOutput(func() {
						callCliCommandPlugin.Run(fakeCliConnection, []string{"org-space", "testorg", "testspace"})
					})
					Expect(fakeCliConnection.CliCommandWithoutTerminalOutputCallCount()).To(Equal(3))

					Expect(len(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(0))).To(Equal(2))
					Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(0)[0]).To(Equal("create-org"))
					Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(0)[1]).To(Equal("testorg"))

					Expect(len(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(1))).To(Equal(4))
					Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(1)[0]).To(Equal("create-space"))
					Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(1)[1]).To(Equal("testspace"))
					Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(1)[2]).To(Equal("-o"))
					Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(1)[3]).To(Equal("testorg"))

					Expect(len(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(2))).To(Equal(5))
					Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(2)[0]).To(Equal("target"))
					Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(2)[1]).To(Equal("-o"))
					Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(2)[2]).To(Equal("testorg"))
					Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(2)[3]).To(Equal("-s"))
					Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(2)[4]).To(Equal("testspace"))
				})
			})
		})
	})
})
