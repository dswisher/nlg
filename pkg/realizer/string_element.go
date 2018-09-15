package realizer

// StringElement represents a canned bit of text.
type StringElement struct {
	value string
}

// NewStringElement constructs a new StringElement.
func NewStringElement(value string) *StringElement {
	return &StringElement{value: value}
}

// GetRealization returns the realization for this element.
func (e *StringElement) GetRealization() string {
	return e.value
}
