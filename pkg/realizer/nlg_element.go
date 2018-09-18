package realizer

// NlgElement represents a single item in the tree to be realized.
type NlgElement interface {
	Category() LexicalCategory
	GetRealization() string
	getFeature(key string) interface{} // TODO - export this
}

type nlgElementData struct {
	category LexicalCategory
	*featureMap
}

func newElementData(category LexicalCategory) nlgElementData {
	return nlgElementData{
		category:   category,
		featureMap: newFeatureMap(),
	}
}

// Category retrieves the category of this word
func (e nlgElementData) Category() LexicalCategory {
	return e.category
}
