package provide

import (
	"github.com/gentra/legosamples/decorator/internal"
	"github.com/gentra/legosamples/decorator/internal/repo/postgredb"
	"github.com/gentra/legosamples/decorator/internal/repo/redis"
)

func BookReaderRepo(useCache bool) internal.BookReaderRepo {
	db := postgredb.NewBookReader()
	switch useCache {
	case true:
		return redis.NewBookReader(db) // we wrap the source-of-truth in the cache class
	default:
		return db
	}
}
