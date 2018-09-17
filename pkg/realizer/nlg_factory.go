package realizer

// NlgFactory is used to construct other objects used in the realization process.
type NlgFactory struct {
	realizer *Realizer
	lexicon  *Lexicon
}

// SetLexicon sets the lexicon to use for the factory.
func (f *NlgFactory) SetLexicon(lexicon *Lexicon) {
	f.lexicon = lexicon
}

// Lexicon returns the current lexicon, constructing a default if needed.
func (f *NlgFactory) Lexicon() *Lexicon {
	if f.lexicon == nil {
		f.lexicon = NewLexicon()
		// TODO - build a better default lexicon
		LoadMinimalWords(f.lexicon)
	}
	return f.lexicon
}

// Realizer returns the realizer, constructing one if needed.
func (f *NlgFactory) Realizer() *Realizer {
	if f.realizer == nil {
		f.realizer = new(Realizer)
		f.realizer.initialize()
	}
	return f.realizer
}

// CreateWordFromString creates a new element representing a word.
func (f *NlgFactory) CreateWordFromString(word string, lexicalCategory LexicalCategory) NlgElement {
	wordElement := f.Lexicon().LookupWord(word, lexicalCategory)
	// TODO - if the word is a pronoun, set pronoun features
	return wordElement
}
