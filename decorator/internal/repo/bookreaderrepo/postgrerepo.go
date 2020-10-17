package bookreaderrepo

import (
	"github.com/gentra/golang-design-patterns/decorator/internal/entity"
	"log"
)

type PostgreRepo struct {
}

func NewPostgreRepo() *PostgreRepo {
	return &PostgreRepo{}
}

func (b *PostgreRepo) ListBook() ([]entity.Book, error) {
	// Let's just pretend we do all the database fetching here
	log.Println("Fetching data from PostgreSQL Database")
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
