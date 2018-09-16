package realizer

// Lexicon provides details on words.
type Lexicon struct {
	baseMap map[string]*WordElement
}

// LexicalCategory is an enumeration of the different types of words.
type LexicalCategory int

const (
	// LexicalCategoryAny indicates an unspecified category
	LexicalCategoryAny LexicalCategory = iota
	// LexicalCategoryAdjective is an adjective
	LexicalCategoryAdjective = iota
	// LexicalCategoryNoun is a noun
	LexicalCategoryNoun = iota
	// TODO - add remaining categories
)

// NewLexicon creates and initializes a new, empty lexicon
func NewLexicon() *Lexicon {
	lexicon := new(Lexicon)

	lexicon.baseMap = make(map[string]*WordElement)

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
	// TODO - need to change baseMap to handle more than one match!
	var list []*WordElement

	match, ok := l.baseMap[baseForm]
	if ok {
		list = append(list, match.Copy())
	}

	return list
}

// Add puts a word into the lexicon
func (l *Lexicon) Add(word *WordElement) {
	l.baseMap[word.baseForm] = word.Copy()
	// TODO - add to index by variant and whatnot?
}
