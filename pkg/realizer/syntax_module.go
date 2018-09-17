package realizer

type syntaxModule struct {
}

func (m *syntaxModule) realize(element NlgElement) NlgElement {

	var realizedElement NlgElement

	if element == nil {
		return nil
	}

	word, ok := element.(*WordElement)
	if ok {
		inflected := newInflectedWordElementFromWord(word)
		inflected.copyFeatures(word.featureMap)
		realizedElement = m.realize(inflected)
	}

	infl, ok := element.(*InflectedWordElement)
	if ok {
		// TODO - syntaxify an inflected word element
		realizedElement = infl
	}

	// TODO - implement syntax module - handle other types

	return realizedElement
}
