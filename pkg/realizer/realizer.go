package realizer

import "fmt"

// Realizer contains the state used by the realizer.
type Realizer struct {
}

// Realize turns the element tree into a string element.
func (r *Realizer) Realize(element NlgElement) NlgElement {
	// TODO - this is a temporary hack - implement the real version!

	word, ok := element.(*WordElement)
	if ok {
		return NewStringElement(word.baseForm)
	}

	panic(fmt.Sprintf("Unable to realize type '%T'.", element))
}
