package realizer

import "testing"

func TestMinimalLexicon(t *testing.T) {
	lexicon := NewLexicon()
	loadMinimalWords(lexicon)

	words := lexicon.GetWords("rock", LexicalCategoryNoun)
	if len(words) != 1 {
		t.Errorf("Lookup of 'rock', expected 1 item, got %d.", len(words))
	}
}
