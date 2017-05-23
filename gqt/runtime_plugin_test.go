package gqt_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/gqt/runner"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Runtime Plugin", func() {
	var (
		args         []string
		argsFilepath string
		client       *runner.RunningGarden
	)

	BeforeEach(func() {
		args = []string{}
	})

	JustBeforeEach(func() {
		client = startGarden(args...)
		argsFilepath = filepath.Join(client.Tmpdir, "args")
	})

	AfterEach(func() {
		Expect(os.Remove(argsFilepath)).To(Succeed())
		Expect(client.DestroyAndStop()).To(Succeed())
	})

	Context("when a runtime plugin is provided", func() {
		BeforeEach(func() {
			args = append(
				args,
				"--runtime-plugin", binaries.RuntimePlugin,
				"--network-plugin", binaries.NoopPlugin,
			)
		})

		Context("and a container is successfully created", func() {
			var handle = fmt.Sprintf("some-handle-%d", GinkgoParallelNode())

			JustBeforeEach(func() {
				_, err := client.Create(garden.ContainerSpec{Handle: handle})
				Expect(err).ToNot(HaveOccurred())
			})

			It("executes the plugin, passing the correct args", func() {
				pluginArgsBytes, err := ioutil.ReadFile(argsFilepath)
				Expect(err).ToNot(HaveOccurred())

				pluginArgs := strings.Split(string(pluginArgsBytes), " ")
				Expect(pluginArgs).To(ConsistOf(
					binaries.RuntimePlugin,
					"--debug",
					"--log", HaveSuffix(filepath.Join("containers", handle, "create.log")),
					"--newuidmap", HaveSuffix("newuidmap"),
					"--newgidmap", HaveSuffix("newgidmap"),
					"create",
					"--no-new-keyring",
					"--bundle", HaveSuffix(filepath.Join("containers", handle)),
					"--pid-file", HaveSuffix(filepath.Join("containers", handle, "pidfile")),
					handle,
				))
			})
		})
	})
})
