package main_test

import (
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudfoundry/gunk/runner_support"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vito/cmdtest"
)

var _ = Describe("Integration Test for Org Space plugin", func() {

	orgName := "testorg"
	spaceName := "testspace"

	BeforeEach(func() {
		cfCmd := exec.Command("cf", "plugins")
		session := runCommand(cfCmd)

		if strings.Contains(string(session.FullOutput()), "OrgSpace") {
			cfCmd := exec.Command("cf", "uninstall-plugin", "OrgSpace")
			session := runCommand(cfCmd)
			Expect(string(session.FullOutput())).NotTo(ContainSubstring("FAILED"))
		}

		plugin, filePathErr := filepath.Abs("org_space.exe")
		Expect(filePathErr).ToNot(HaveOccurred())

		cfCmd = exec.Command("cf", "install-plugin", plugin)
		session = runCommand(cfCmd)
		Expect(string(session.FullOutput())).NotTo(ContainSubstring("FAILED"))
	})

	AfterEach(func() {
		cfCmd := exec.Command("cf", "uninstall-plugin", "OrgSpace")
		session := runCommand(cfCmd)
		Expect(string(session.FullOutput())).NotTo(ContainSubstring("FAILED"))
	})

	Context("when the org and space does not exist", func() {

		BeforeEach(func() {
			cfCmd := exec.Command("cf", "orgs")
			session := runCommand(cfCmd)

			if strings.Contains(string(session.FullOutput()), orgName) {
				cfCmd := exec.Command("cf", "delete-org", orgName, "-f")
				session := runCommand(cfCmd)
				Expect(string(session.FullOutput())).NotTo(ContainSubstring("does not exist."))
			}
		})

		AfterEach(func() {
			cfCmd := exec.Command("cf", "delete-org", orgName, "-f")
			session := runCommand(cfCmd)
			Expect(string(session.FullOutput())).NotTo(ContainSubstring("does not exist."))
		})

		It("creates the org and space then targets it", func() {
			cfCmd := exec.Command("cf", "org-space", orgName, spaceName)
			session := runCommand(cfCmd)
			Expect(string(session.FullOutput())).To(ContainSubstring("Org testorg and Space testspace is now available and targeted"))

			cfCmd = exec.Command("cf", "orgs")
			session = runCommand(cfCmd)
			Expect(string(session.FullOutput())).To(ContainSubstring(orgName))

			cfCmd = exec.Command("cf", "spaces")
			session = runCommand(cfCmd)
			Expect(string(session.FullOutput())).To(ContainSubstring(spaceName))
		})
	})
})

func runCommand(cmd *exec.Cmd) *cmdtest.Session {
	session, err := cmdtest.StartWrapped(
		cmd,
		runner_support.TeeIfVerbose,
		runner_support.TeeIfVerbose,
	)
	Expect(err).NotTo(HaveOccurred())

	_, err = session.Wait(10 * time.Second)
	Expect(err).NotTo(HaveOccurred())

	return session
}
