package realizer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dswisher/nlg/pkg/realizer"
)

var _ = Describe("Realizer", func() {

	var (
		factory  *NlgFactory
		realizer *Realizer
	)

	BeforeEach(func() {
		factory = new(NlgFactory)
		realizer = factory.Realizer()
	})

	Describe("Single words", func() {
		Context("Made up word, without features", func() {
			It("Should be the same", func() {
				word := NewWordElement("muggle", LexicalCategoryAny)
				Expect(realizer.Realize(word).GetRealization()).To(Equal("muggle"))
			})
		})
	})
})
