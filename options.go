package fnet

type (
	OptionTypes int

	ResultType[T any] struct {
		Inner T
		Outer OptionTypes
	}
)

const (
	Some OptionTypes = iota
	None
)

func OptionMatch[T comparable](V T) OptionTypes {
	switch V {
	default:
		return Some
	case *new(T):
		return None
	}
}

func Ok[T comparable](V T) ResultType[T] {
	return ResultType[T]{Inner: V, Outer: OptionMatch(V)}
}
