package output_test

import (
	"testing"

	"code.cloudfoundry.org/lager/v3"
	"code.cloudfoundry.org/lager/v3/lagertest"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOutput(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Output Suite")
}

var (
	logger lager.Logger
)

var _ = BeforeSuite(func() {
	logger = lagertest.NewTestLogger("output-test")
})
