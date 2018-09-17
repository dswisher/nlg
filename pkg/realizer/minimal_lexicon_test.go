package realizer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dswisher/nlg/pkg/realizer"
)

var _ = Describe("Minimal lexicon", func() {
	var lexicon *Lexicon

	BeforeEach(func() {
		lexicon = NewLexicon()
		LoadMinimalWords(lexicon)
	})

	Describe("GetWords", func() {
		Context("Boy", func() {
			It("Is a noun", func() {
				words := lexicon.GetWords("boy", LexicalCategoryNoun)
				Expect(words).To(HaveLen(1))
				Expect(words[0].Category()).To(Equal(LexicalCategoryNoun))
				// TODO - check features
			})
		})

		Context("Rock", func() {
			// TODO - lookup "ANY" and make sure we got two items

			It("Is a noun", func() {
				words := lexicon.GetWords("rock", LexicalCategoryNoun)
				Expect(words).To(HaveLen(1))
				Expect(words[0].Category()).To(Equal(LexicalCategoryNoun))
				// TODO - check features
			})

			It("Is a verb", func() {
				words := lexicon.GetWords("rock", LexicalCategoryVerb)
				Expect(words).To(HaveLen(1))
				Expect(words[0].Category()).To(Equal(LexicalCategoryVerb))
				// TODO - check features
			})
		})
	})
})
