package healthcheck_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHealthcheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Healthcheck Suite")
}
