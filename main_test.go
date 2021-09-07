package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main Test", func() {
	Context("With a demo example", func() {
		It("shoud be succeed", func() {
			Expect(true).To(Equal(true))
		})
	})
})
