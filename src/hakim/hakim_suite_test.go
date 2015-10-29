package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHakim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hakim Suite")
}
