package realizer

// Realizer contains the state used by the realizer.
type Realizer struct {
	modules []realizerModule
}

func (r *Realizer) initialize() {
	r.modules = append(r.modules, new(syntaxModule))
	r.modules = append(r.modules, new(morphologyModule))
	r.modules = append(r.modules, new(orthographyModule))
	// TODO - add formatter module
}

// Realize turns the element tree into a string element.
func (r *Realizer) Realize(element NlgElement) NlgElement {
	realized := element

	for _, module := range r.modules {
		realized = module.realize(realized)
	}

	return realized

	// word, ok := element.(*WordElement)
	// if ok {
	// 	return NewStringElement(word.baseForm)
	// }

	// panic(fmt.Sprintf("Unable to realize type '%T'.", element))
}
