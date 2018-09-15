package realizer

// NlgFactory is used to construct other objects used in the realization process.
type NlgFactory struct {
	realizer *Realizer
}

// Realizer returns the realizer, constructing one if needed.
func (f *NlgFactory) Realizer() *Realizer {
	if f.realizer == nil {
		f.realizer = new(Realizer)
	}
	return f.realizer
}
