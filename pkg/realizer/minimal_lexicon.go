package realizer

// LoadMinimalWords adds a minimal set of words to a lexicon, mainly intended for testing.
func LoadMinimalWords(lexicon *Lexicon) {
	lexicon.Add(NewWordElement("boy", LexicalCategoryNoun))
	lexicon.Add(NewWordElement("chase", LexicalCategoryVerb)) // intransitive, transitive
	lexicon.Add(NewWordElement("monkey", LexicalCategoryNoun))
	lexicon.Add(NewWordElement("rock", LexicalCategoryNoun))
	lexicon.Add(NewWordElement("rock", LexicalCategoryVerb)) // intransitive, transitive
}
