package postgredb

import (
	"github.com/gentra/legosamples/decorator/internal/entity"
)

type BookReader struct {
}

func NewBookReader() *BookReader {
	return &BookReader{}
}

func (b *BookReader) ListBook() ([]entity.Book, error) {
	// Let's just pretend we do all the database fetching here
	return []entity.Book{
		{
			ID:    1,
			Title: "You got this book from DB",
		},
		{
			ID:    2,
			Title: "Database",
		},
	}, nil
}
