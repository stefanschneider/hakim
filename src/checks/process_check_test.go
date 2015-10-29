package checks_test

import (
	. "checks"

	"github.com/mitchellh/go-ps"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeProc struct {
	ps.Process
	name string
}

func (p fakeProc) Executable() string {
	return p.name
}

var _ = Describe("Process checks", func() {
	It("returns an error when it doesn't find all its required processes", func() {
		processes := []ps.Process{fakeProc{name: "consul.exe"}}
		requiredProcesses := []string{"Containerizer.exe", "consul.exe"}
		Expect(ProcessCheck(processes, requiredProcesses)).To(MatchError("The following processes are not running: Containerizer.exe"))
	})

	It("returns an error for multiple processes if multiple are missing", func() {
		processes := []ps.Process{fakeProc{name: "foo.exe"}, fakeProc{name: "consul.exe"}}
		requiredProcesses := []string{"Containerizer.exe", "consul.exe", "garden-windows.exe"}
		Expect(ProcessCheck(processes, requiredProcesses)).To(MatchError("The following processes are not running: Containerizer.exe, garden-windows.exe"))
	})

	It("does not return an error when it finds all its required processes", func() {
		processes := []ps.Process{fakeProc{name: "Containerizer.exe"}}
		requiredProcesses := []string{"Containerizer.exe"}
		Expect(ProcessCheck(processes, requiredProcesses)).To(Succeed())
	})
})
