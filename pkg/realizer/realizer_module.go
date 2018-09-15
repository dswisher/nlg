package realizer

type realizerModule interface {
	realize(element NlgElement) NlgElement
}
