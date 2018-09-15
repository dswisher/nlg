package realizer

// NlgElement represents a single item in the tree to be realized.
type NlgElement interface {
	GetRealization() string
}
