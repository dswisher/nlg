package realizer_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRealizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Realizer Suite")
}
