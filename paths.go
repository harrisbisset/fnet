package fnet

type (
	Path int
)

const (
	GET Path = iota
	POST
	UPDATE
	DELETE
	PUT
)

func (p Path) String() string {
	switch p {

	// assume GET
	default:
		return "GET"

	case 1:
		return "POST"

	case 2:
		return "UPDATE"

	case 3:
		return "DELETE"

	case 4:
		return "PUT"

	}
}
