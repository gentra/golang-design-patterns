package provide

import (
	"github.com/gentra/legosamples/decorator/internal"
	"github.com/gentra/legosamples/decorator/internal/repo/bookreaderrepo"
)

func BookReaderRepo(useCache bool) internal.BookReaderRepo {
	db := bookreaderrepo.NewPostgreRepo()
	switch useCache {
	case true:
		return bookreaderrepo.NewRedisRepo(db) // we wrap the source-of-truth in the cache class
	default:
		return db
	}
}
