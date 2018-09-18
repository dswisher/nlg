package realizer

// StringElement represents a canned bit of text.
type StringElement struct {
	nlgElementData
	value string
}

// NewStringElement constructs a new StringElement.
func NewStringElement(value string) *StringElement {
	elem := &StringElement{value: value,
		nlgElementData: newElementData(LexicalCategoryAny),
	}
	return elem
}

// GetRealization returns the realization for this element.
func (e *StringElement) GetRealization() string {
	return e.value
}
