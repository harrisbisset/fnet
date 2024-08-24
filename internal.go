package fnet

//
// internal types below
//

type (
	responseErr int
)

const (
	RenderFail responseErr = iota
	ErrorFail
)

func (r responseErr) View() string {
	switch r {
	default:
		return "view"
	case ErrorFail:
		return "error"
	}
}

//
// internal functions below
//

func match[T comparable](V T) bool {
	switch V {
	default:
		return true
	case *new(T):
		return false
	}
}

func present[T comparable](V T) bool {
	return !(V == *new(T))
}

func panicField[T comparable](field T) {
	if !present(field) {
		panic("field required")
	}
}
