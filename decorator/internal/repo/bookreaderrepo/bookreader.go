package bookreaderrepo

import (
	"github.com/gentra/legosamples/decorator/internal"
	"github.com/gentra/legosamples/decorator/internal/entity"
	"log"
)

type RedisRepo struct {
	source internal.BookReaderRepo
}

func NewRedisRepo(source internal.BookReaderRepo) *RedisRepo {
	return &RedisRepo{source: source}
}

func (b *RedisRepo) ListBook() ([]entity.Book, error) {
	books, err := getFromRedis()
	if err == nil { // immediately return when cache-hit
		return books, nil
	}

	// otherwise fetch from deeper data-source
	books, err = b.source.ListBook()
	if err != nil {
		return nil, err
	}

	err = setToRedis(books)
	if err != nil {
		log.Println("failed to write books to redis")
	}

	return books, nil
}

func getFromRedis() ([]entity.Book, error) {
	// Let's just pretend this function fetches object from Redis
	return []entity.Book{
		{
			ID:    1,
			Title: "You got this book from Redis",
		},
		{
			ID:    2,
			Title: "Redis",
		},
	}, nil
}

func setToRedis([]entity.Book) error {
	// Let's just pretend this function stores Books object to Redis
	return nil
}
