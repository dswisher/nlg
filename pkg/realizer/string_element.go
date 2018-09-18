package realizer

// StringElement represents a canned bit of text.
type StringElement struct {
	*featureMap
	category LexicalCategory
	value    string
}

// NewStringElement constructs a new StringElement.
func NewStringElement(value string) *StringElement {
	elem := &StringElement{value: value}
	elem.featureMap = newFeatureMap()
	return elem
}

// Category retrieves the category of this word
func (e *StringElement) Category() LexicalCategory {
	return e.category
}

// GetRealization returns the realization for this element.
func (e *StringElement) GetRealization() string {
	return e.value
}
