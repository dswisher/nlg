package realizer

import "testing"

func TestSimple(t *testing.T) {
	muggle := "muggle"
	factory := new(NlgFactory)
	realizer := factory.Realizer()
	// word := factory.CreateWord(muggle)
	word := NewWordElement(muggle)

	result := realizer.Realize(word)

	if result.GetRealization() != muggle {
		t.Errorf("Realization incorrect, expected '%s', got '%s'.", muggle, result.GetRealization())
	}
}
