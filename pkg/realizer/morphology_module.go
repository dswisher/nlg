package realizer

type morphologyModule struct {
}

func (m *morphologyModule) realize(element NlgElement) NlgElement {

	var realized NlgElement
	var baseWord *WordElement

	infl, ok := element.(*InflectedWordElement)
	if ok {
		// TODO - check for NON_MORPH internal feature

		// TODO - get the base word

		switch infl.Category() {
		// TODO - PRONOUN
		case LexicalCategoryNoun:
			realized = doNounMorphology(infl, baseWord)
		// TODO - VERB
		// TODO - ADJECTIVE
		// TODO - ADVERB
		default:
			realized = NewStringElement(infl.getFeatureAsString(LexicalFeatureBaseForm))
			// TODO - set DISCOURSE_FUNCTION feature
		}
	}

	// TODO - implement morphology module - handle additional types

	return realized
}

func doNounMorphology(element *InflectedWordElement, baseWord *WordElement) *StringElement {
	// TODO - doNounMorphology
	return NewStringElement(element.getFeatureAsString(LexicalFeatureBaseForm))
}
