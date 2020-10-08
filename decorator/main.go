package main

import (
	"log"

	"github.com/gentra/legosamples/decorator/internal/provide"
)

func main() {
	useCache := true // Let's just pretend that you get this use-cache configuration from config files or cli param
	repo := provide.BookReaderRepo(useCache)

	books, err := repo.ListBook()
	if err != nil {
		log.Println(err)
	}

	log.Printf("%+v", books)
}
