package fnet

type (
	OptionTypes int

	ResultType[T any] struct {
		Value  T
		Result OptionTypes
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
	return ResultType[T]{Value: V, Result: OptionMatch(V)}
}

func Present[T comparable](V T) bool {
	return !(V == *new(T))
}
