package realizer

type morphologyModule struct {
}

func (m *morphologyModule) realize(element NlgElement) NlgElement {

	var realized NlgElement

	infl, ok := element.(*InflectedWordElement)
	if ok {
		// TODO - doMorphology!
		realized = NewStringElement(infl.getFeatureAsString(LexicalFeatureBaseForm))
	}

	// TODO - implement morphology module - handle additional types

	return realized
}
