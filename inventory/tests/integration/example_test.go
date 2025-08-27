//go:build integration

package integration

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example Integration Test", func() {

	It("должен успешно запуститься", func() {
		Expect(true).To(BeTrue())
	})
})
