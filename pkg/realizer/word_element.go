package realizer

// WordElement represents a word and its features.
type WordElement struct {
	*featureMap
	baseForm string
}

// NewWordElement constructs a new WordElement.
func NewWordElement(baseForm string) *WordElement {
	word := &WordElement{baseForm: baseForm}
	// TODO - accept a lexical category
	word.featureMap = newFeatureMap()
	return word
}

// GetRealization returns the realization for this element.
func (e *WordElement) GetRealization() string {
	return ""
}
