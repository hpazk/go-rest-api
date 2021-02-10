package helper

import "github.com/golangkit/formatime"

type BaseModel struct {
	ID        int
	CreatedAt formatime.Timestamp
	UpdatedAt formatime.Timestamp
	DeletedAt *formatime.Timestamp
}
