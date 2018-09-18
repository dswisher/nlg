package realizer

// Lexicon provides details on words.
type Lexicon struct {
	// TODO - there can be multiple entries with the same base: "rock" -> NOUN, "rock" -> VERB
	baseMap map[string][]*WordElement
}

// LexicalCategory is an enumeration of the different types of words.
type LexicalCategory int

const (
	// LexicalCategoryAny indicates an unspecified category
	LexicalCategoryAny LexicalCategory = iota
	// LexicalCategoryNoun is a noun
	LexicalCategoryNoun
	// LexicalCategoryAdjective is an adjective
	LexicalCategoryAdjective
	// LexicalCategoryAdverb is an adverb
	LexicalCategoryAdverb
	// LexicalCategoryVerb is an verb
	LexicalCategoryVerb
	// LexicalCategoryDeterminer is an determiner
	LexicalCategoryDeterminer
	// LexicalCategoryPronoun is an pronoun
	LexicalCategoryPronoun
	// LexicalCategoryConjunction is an conjunction
	LexicalCategoryConjunction
	// LexicalCategoryPreposition is an preposition
	LexicalCategoryPreposition
	// LexicalCategoryComplementizer is an complementizer
	LexicalCategoryComplementizer
	// LexicalCategoryModal is an modal
	LexicalCategoryModal
	// LexicalCategoryAuxiliary is an auziliary
	LexicalCategoryAuxiliary
)

// LexicalFeature is an enumeration of the features that can be applied to words.
type LexicalFeature int

const (
	// LexicalFeatureDefaultInflection indicates the default inflectional variant for a word
	LexicalFeatureDefaultInflection LexicalFeature = iota
)

// NewLexicon creates and initializes a new, empty lexicon
func NewLexicon() *Lexicon {
	lexicon := new(Lexicon)

	lexicon.baseMap = make(map[string][]*WordElement)

	return lexicon
}

// LookupWord finds a word in the lexicon and returns it.
func (l *Lexicon) LookupWord(word string, lexicalCategory LexicalCategory) *WordElement {
	words := l.GetWords(word, lexicalCategory)
	if len(words) > 0 {
		return words[0].Copy()
	}

	// TODO - better choosing when there are multiple matches
	// TODO - look up by other forms?

	return NewWordElement(word, lexicalCategory)
}

// GetWords returns a list of matching words
func (l *Lexicon) GetWords(baseForm string, lexicalCategory LexicalCategory) []*WordElement {
	match, ok := l.baseMap[baseForm]
	if !ok {
		return []*WordElement{}
	}

	if lexicalCategory == LexicalCategoryAny {
		return match
	}

	for _, word := range match {
		if word.Category() == lexicalCategory {
			return []*WordElement{word}
		}
	}

	return []*WordElement{}
}

// Add puts a word into the lexicon
func (l *Lexicon) Add(word *WordElement) {
	clone := word.Copy()

	if val, ok := l.baseMap[word.baseForm]; ok {
		// TODO - if a word with this category already exists, replace it
		l.baseMap[word.baseForm] = append(val, clone)
	} else {
		l.baseMap[word.baseForm] = []*WordElement{clone}
	}

	// TODO - add to index by variant and whatnot?
}
