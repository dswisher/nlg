package realizer

// WordElement represents a word and its features.
type WordElement struct {
	baseForm string
}

// NewWordElement constructs a new WordElement.
func NewWordElement(baseForm string) *WordElement {
	// TODO - accept a lexical category
	return &WordElement{baseForm: baseForm}
}

// GetRealization returns the realization for this element.
func (e *WordElement) GetRealization() string {
	return ""
}
