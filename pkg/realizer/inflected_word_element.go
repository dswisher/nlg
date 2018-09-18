package realizer

// InflectedWordElement represents a word that needs to be inflected by the morphology module.
type InflectedWordElement struct {
	nlgElementData
}

func newInflectedWordElementFromWord(word *WordElement) *InflectedWordElement {
	inflected := &InflectedWordElement{
		nlgElementData: newElementData(word.Category()),
	}

	// TODO - set the BASE_WORD feature
	// TODO - get the default spelling and use it to set the BASE_FORM feature
	inflected.setFeature(LexicalFeatureBaseForm, word.baseForm)
	// TODO - set the category

	return inflected
}

// GetRealization returns the realization for this element.
func (e *InflectedWordElement) GetRealization() string {
	return ""
}
