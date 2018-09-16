package realizer

import "testing"

func TestSyntaxModule(t *testing.T) {
	module := new(syntaxModule)

	if module.realize(nil) != nil {
		t.Errorf("Attempting to realize nil should yield nil.")
	}

	featureName := "color"
	featureValue := "red"

	word := NewWordElement("silly", LexicalCategoryAdjective)
	word.setFeature(featureName, featureValue)
	result := module.realize(word)
	infl, ok := result.(*InflectedWordElement)
	if !ok {
		t.Errorf("Attempting to realize word should inflected word, but got %T.", result)
	}

	actualValue := infl.getFeatureAsString(featureName)
	if infl.getFeature(featureName) != featureValue {
		t.Errorf("Inflected result should have feature '%s', with value '%s', but got value '%v'.", featureName, featureValue, actualValue)
	}
}
