package bookreaderrepo

import (
	"github.com/gentra/legosamples/decorator/internal/entity"
)

type PostgreRepo struct {
}

func NewPostgreRepo() *PostgreRepo {
	return &PostgreRepo{}
}

func (b *PostgreRepo) ListBook() ([]entity.Book, error) {
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
