package internal

import (
	"github.com/gentra/golang-design-patterns/decorator/internal/entity"
)

type BookReaderRepo interface {
	ListBook() ([]entity.Book, error)
}
