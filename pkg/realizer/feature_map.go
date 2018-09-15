package realizer

import "fmt"

type featureMap struct {
	features map[string]interface{}
}

const (
	// LexicalFeatureBaseForm defines the base form for words and phrases
	LexicalFeatureBaseForm  = "base_form"
	internalFeatureBaseWord = "base_word"
)

func newFeatureMap() *featureMap {
	fm := new(featureMap)
	fm.features = make(map[string]interface{})
	return fm
}

func (f *featureMap) copyFeatures(other *featureMap) {
	for k, v := range other.features {
		f.features[k] = v
	}
}

func (f *featureMap) setFeature(key string, value interface{}) {
	f.features[key] = value
}

func (f *featureMap) getFeature(key string) interface{} {
	val, ok := f.features[key]
	if ok {
		return val
	}
	return nil
}

func (f *featureMap) getFeatureAsString(key string) string {
	val := f.getFeature(key)
	if val == nil {
		return ""
	}
	return fmt.Sprintf("%v", f.getFeature(key))
}
