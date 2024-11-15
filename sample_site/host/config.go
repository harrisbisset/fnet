package host

import (
	"database/sql"
)

type (
	CFG struct {
		DB       *sql.DB
		Language *int
	}
)

func CreateConfig() CFG {
	return CFG{
		Language: nil,
	}
}
