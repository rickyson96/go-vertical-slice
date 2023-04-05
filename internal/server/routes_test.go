package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/rickyson96/go-vertical-slice/internal/server"
)

var _ = Describe("Routes", func() {
	Context("creating routes", func() {
		It("should be able to create new router handler", func() {
			Expect(server.Routes(nil, nil)).NotTo(BeNil())
		})
	})
})
