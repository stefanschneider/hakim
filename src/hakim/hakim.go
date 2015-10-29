package main

import (
	"checks"
	"log"

	"github.com/mitchellh/go-ps"
)

var requiredProcesses = []string{"consul.exe", "containerizer.exe", "garden-windows.exe", "rep.exe", "metron.exe"}

func main() {
	processes, err := ps.Processes()
	if err != nil {
		panic(err)
	}
	err = checks.ProcessCheck(processes, requiredProcesses)
	if err != nil {
		log.Fatal(err)
	}
}
