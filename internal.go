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

//
// internal functions below
//

func present[T comparable](V T) bool {
	return !(V == *new(T))
}

func panicField[T comparable](field T) {
	if !present(field) {
		panic("field required")
	}
}
