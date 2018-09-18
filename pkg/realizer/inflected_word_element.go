package realizer

// InflectedWordElement represents a word that needs to be inflected by the morphology module.
type InflectedWordElement struct {
	*featureMap
	category LexicalCategory
}

func newInflectedWordElementFromWord(word *WordElement) *InflectedWordElement {
	inflected := new(InflectedWordElement)
	inflected.featureMap = newFeatureMap()

	// TODO - set the BASE_WORD feature
	// TODO - get the default spelling and use it to set the BASE_FORM feature
	inflected.setFeature(LexicalFeatureBaseForm, word.baseForm)
	// TODO - set the category

	return inflected
}

// Category retrieves the category of this word
func (e *InflectedWordElement) Category() LexicalCategory {
	return e.category
}

// GetRealization returns the realization for this element.
func (e *InflectedWordElement) GetRealization() string {
	return ""
}
