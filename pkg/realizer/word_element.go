package realizer

// WordElement represents a word and its features.
type WordElement struct {
	nlgElementData
	baseForm string
}

// NewWordElement constructs a new WordElement.
func NewWordElement(baseForm string, category LexicalCategory) *WordElement {
	word := &WordElement{
		baseForm:       baseForm,
		nlgElementData: newElementData(category),
	}
	return word
}

// GetRealization returns the realization for this element.
func (e *WordElement) GetRealization() string {
	return ""
}

// Copy makes a copy of the word
func (e *WordElement) Copy() *WordElement {
	copy := NewWordElement(e.baseForm, e.category)
	copy.copyFeatures(e.featureMap)
	return copy
}
