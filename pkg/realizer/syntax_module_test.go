package realizer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SyntaxModule", func() {
	var module *syntaxModule

	BeforeEach(func() {
		module = new(syntaxModule)
	})

	Describe("Realize nil", func() {
		It("Should yield nil", func() {
			Expect(module.realize(nil)).To(BeNil())
		})
	})

	Describe("Realize WordElement", func() {
		var word *WordElement

		const (
			featureName  = "color"
			featureValue = "red"
		)

		BeforeEach(func() {
			word = NewWordElement("silly", LexicalCategoryAdjective)
			word.setFeature(featureName, featureValue)
		})

		It("Should be inflected with features", func() {
			result := module.realize(word)
			Expect(result).To(BeAssignableToTypeOf(new(InflectedWordElement)))
			Expect(result.getFeature(featureName)).To(Equal(featureValue))
		})
	})
})
