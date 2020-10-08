package internal

import (
	"github.com/gentra/legosamples/decorator/internal/entity"
)

type BookReaderRepo interface {
	ListBook() ([]entity.Book, error)
}
