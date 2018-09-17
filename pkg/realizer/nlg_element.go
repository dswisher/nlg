package realizer

// NlgElement represents a single item in the tree to be realized.
type NlgElement interface {
	GetRealization() string
	getFeature(key string) interface{} // TODO - export this
}
