package realizer

// StringElement represents a canned bit of text.
type StringElement struct {
	*featureMap
	value string
}

// NewStringElement constructs a new StringElement.
func NewStringElement(value string) *StringElement {
	elem := &StringElement{value: value}
	elem.featureMap = newFeatureMap()
	return elem
}

// GetRealization returns the realization for this element.
func (e *StringElement) GetRealization() string {
	return e.value
}
