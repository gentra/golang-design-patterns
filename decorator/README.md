# Introduction

Decorator is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors.

[![](https://mermaid.ink/img/eyJjb2RlIjoiY2xhc3NEaWFncmFtXG4gIGNsYXNzIG1haW57XG4gICAgbWFpbigpXG4gIH1cblxuICBjbGFzcyBCb29rUmVhZGVyUmVwbyB7XG4gICAgPDxpbnRlcmZhY2U-PlxuICAgICtMaXN0Qm9vayhpZCkgKFtdQm9vaywgZXJyb3IpXG4gIH1cblxuICBCb29rUmVhZGVyUmVwbzx8Li5SZWRpc1JlcG9cbiAgQm9va1JlYWRlclJlcG88fC4uUG9zdGdyZVJlcG9cbiAgUmVkaXNSZXBvLi4-Qm9va1JlYWRlclJlcG86IGNhbGxzIGRlZXBlciBkYXRhLXNvdXJjZVxuXG4gIGNsYXNzIGJvb2tSZWFkZXJ7XG4gICAgPDxpbnN0YW5jZT4-XG4gIH1cbiAgUmVkaXNSZXBvIC0tbyBib29rUmVhZGVyOiBwYXJ0IG9mXG4gIFBvc3RncmVSZXBvIC4ubyBib29rUmVhZGVyOiBpbmRpcmVjdGx5IHBhcnQgb2ZcblxuICBjbGFzcyBwcm92aWRlQm9va1JlYWRlclJlcG97XG4gICAgPDxmdW5jdGlvbj4-XG4gIH1cblxuICBwcm92aWRlQm9va1JlYWRlclJlcG8uLmJvb2tSZWFkZXJcblxuICBtYWluLi5wcm92aWRlQm9va1JlYWRlclJlcG9cbiAgbWFpbi4uPkJvb2tSZWFkZXJSZXBvIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQiLCJ0aGVtZVZhcmlhYmxlcyI6eyJiYWNrZ3JvdW5kIjoid2hpdGUiLCJwcmltYXJ5Q29sb3IiOiIjRUNFQ0ZGIiwic2Vjb25kYXJ5Q29sb3IiOiIjZmZmZmRlIiwidGVydGlhcnlDb2xvciI6ImhzbCg4MCwgMTAwJSwgOTYuMjc0NTA5ODAzOSUpIiwicHJpbWFyeUJvcmRlckNvbG9yIjoiaHNsKDI0MCwgNjAlLCA4Ni4yNzQ1MDk4MDM5JSkiLCJzZWNvbmRhcnlCb3JkZXJDb2xvciI6ImhzbCg2MCwgNjAlLCA4My41Mjk0MTE3NjQ3JSkiLCJ0ZXJ0aWFyeUJvcmRlckNvbG9yIjoiaHNsKDgwLCA2MCUsIDg2LjI3NDUwOTgwMzklKSIsInByaW1hcnlUZXh0Q29sb3IiOiIjMTMxMzAwIiwic2Vjb25kYXJ5VGV4dENvbG9yIjoiIzAwMDAyMSIsInRlcnRpYXJ5VGV4dENvbG9yIjoicmdiKDkuNTAwMDAwMDAwMSwgOS41MDAwMDAwMDAxLCA5LjUwMDAwMDAwMDEpIiwibGluZUNvbG9yIjoiIzMzMzMzMyIsInRleHRDb2xvciI6IiMzMzMiLCJtYWluQmtnIjoiI0VDRUNGRiIsInNlY29uZEJrZyI6IiNmZmZmZGUiLCJib3JkZXIxIjoiIzkzNzBEQiIsImJvcmRlcjIiOiIjYWFhYTMzIiwiYXJyb3doZWFkQ29sb3IiOiIjMzMzMzMzIiwiZm9udEZhbWlseSI6IlwidHJlYnVjaGV0IG1zXCIsIHZlcmRhbmEsIGFyaWFsIiwiZm9udFNpemUiOiIxNnB4IiwibGFiZWxCYWNrZ3JvdW5kIjoiI2U4ZThlOCIsIm5vZGVCa2ciOiIjRUNFQ0ZGIiwibm9kZUJvcmRlciI6IiM5MzcwREIiLCJjbHVzdGVyQmtnIjoiI2ZmZmZkZSIsImNsdXN0ZXJCb3JkZXIiOiIjYWFhYTMzIiwiZGVmYXVsdExpbmtDb2xvciI6IiMzMzMzMzMiLCJ0aXRsZUNvbG9yIjoiIzMzMyIsImVkZ2VMYWJlbEJhY2tncm91bmQiOiIjZThlOGU4IiwiYWN0b3JCb3JkZXIiOiJoc2woMjU5LjYyNjE2ODIyNDMsIDU5Ljc3NjUzNjMxMjglLCA4Ny45MDE5NjA3ODQzJSkiLCJhY3RvckJrZyI6IiNFQ0VDRkYiLCJhY3RvclRleHRDb2xvciI6ImJsYWNrIiwiYWN0b3JMaW5lQ29sb3IiOiJncmV5Iiwic2lnbmFsQ29sb3IiOiIjMzMzIiwic2lnbmFsVGV4dENvbG9yIjoiIzMzMyIsImxhYmVsQm94QmtnQ29sb3IiOiIjRUNFQ0ZGIiwibGFiZWxCb3hCb3JkZXJDb2xvciI6ImhzbCgyNTkuNjI2MTY4MjI0MywgNTkuNzc2NTM2MzEyOCUsIDg3LjkwMTk2MDc4NDMlKSIsImxhYmVsVGV4dENvbG9yIjoiYmxhY2siLCJsb29wVGV4dENvbG9yIjoiYmxhY2siLCJub3RlQm9yZGVyQ29sb3IiOiIjYWFhYTMzIiwibm90ZUJrZ0NvbG9yIjoiI2ZmZjVhZCIsIm5vdGVUZXh0Q29sb3IiOiJibGFjayIsImFjdGl2YXRpb25Cb3JkZXJDb2xvciI6IiM2NjYiLCJhY3RpdmF0aW9uQmtnQ29sb3IiOiIjZjRmNGY0Iiwic2VxdWVuY2VOdW1iZXJDb2xvciI6IndoaXRlIiwic2VjdGlvbkJrZ0NvbG9yIjoicmdiYSgxMDIsIDEwMiwgMjU1LCAwLjQ5KSIsImFsdFNlY3Rpb25Ca2dDb2xvciI6IndoaXRlIiwic2VjdGlvbkJrZ0NvbG9yMiI6IiNmZmY0MDAiLCJ0YXNrQm9yZGVyQ29sb3IiOiIjNTM0ZmJjIiwidGFza0JrZ0NvbG9yIjoiIzhhOTBkZCIsInRhc2tUZXh0TGlnaHRDb2xvciI6IndoaXRlIiwidGFza1RleHRDb2xvciI6IndoaXRlIiwidGFza1RleHREYXJrQ29sb3IiOiJibGFjayIsInRhc2tUZXh0T3V0c2lkZUNvbG9yIjoiYmxhY2siLCJ0YXNrVGV4dENsaWNrYWJsZUNvbG9yIjoiIzAwMzE2MyIsImFjdGl2ZVRhc2tCb3JkZXJDb2xvciI6IiM1MzRmYmMiLCJhY3RpdmVUYXNrQmtnQ29sb3IiOiIjYmZjN2ZmIiwiZ3JpZENvbG9yIjoibGlnaHRncmV5IiwiZG9uZVRhc2tCa2dDb2xvciI6ImxpZ2h0Z3JleSIsImRvbmVUYXNrQm9yZGVyQ29sb3IiOiJncmV5IiwiY3JpdEJvcmRlckNvbG9yIjoiI2ZmODg4OCIsImNyaXRCa2dDb2xvciI6InJlZCIsInRvZGF5TGluZUNvbG9yIjoicmVkIiwibGFiZWxDb2xvciI6ImJsYWNrIiwiZXJyb3JCa2dDb2xvciI6IiM1NTIyMjIiLCJlcnJvclRleHRDb2xvciI6IiM1NTIyMjIiLCJjbGFzc1RleHQiOiIjMTMxMzAwIiwiZmlsbFR5cGUwIjoiI0VDRUNGRiIsImZpbGxUeXBlMSI6IiNmZmZmZGUiLCJmaWxsVHlwZTIiOiJoc2woMzA0LCAxMDAlLCA5Ni4yNzQ1MDk4MDM5JSkiLCJmaWxsVHlwZTMiOiJoc2woMTI0LCAxMDAlLCA5My41Mjk0MTE3NjQ3JSkiLCJmaWxsVHlwZTQiOiJoc2woMTc2LCAxMDAlLCA5Ni4yNzQ1MDk4MDM5JSkiLCJmaWxsVHlwZTUiOiJoc2woLTQsIDEwMCUsIDkzLjUyOTQxMTc2NDclKSIsImZpbGxUeXBlNiI6ImhzbCg4LCAxMDAlLCA5Ni4yNzQ1MDk4MDM5JSkiLCJmaWxsVHlwZTciOiJoc2woMTg4LCAxMDAlLCA5My41Mjk0MTE3NjQ3JSkifX0sInVwZGF0ZUVkaXRvciI6ZmFsc2V9)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoiY2xhc3NEaWFncmFtXG4gIGNsYXNzIG1haW57XG4gICAgbWFpbigpXG4gIH1cblxuICBjbGFzcyBCb29rUmVhZGVyUmVwbyB7XG4gICAgPDxpbnRlcmZhY2U-PlxuICAgICtMaXN0Qm9vayhpZCkgKFtdQm9vaywgZXJyb3IpXG4gIH1cblxuICBCb29rUmVhZGVyUmVwbzx8Li5SZWRpc1JlcG9cbiAgQm9va1JlYWRlclJlcG88fC4uUG9zdGdyZVJlcG9cbiAgUmVkaXNSZXBvLi4-Qm9va1JlYWRlclJlcG86IGNhbGxzIGRlZXBlciBkYXRhLXNvdXJjZVxuXG4gIGNsYXNzIGJvb2tSZWFkZXJ7XG4gICAgPDxpbnN0YW5jZT4-XG4gIH1cbiAgUmVkaXNSZXBvIC0tbyBib29rUmVhZGVyOiBwYXJ0IG9mXG4gIFBvc3RncmVSZXBvIC4ubyBib29rUmVhZGVyOiBpbmRpcmVjdGx5IHBhcnQgb2ZcblxuICBjbGFzcyBwcm92aWRlQm9va1JlYWRlclJlcG97XG4gICAgPDxmdW5jdGlvbj4-XG4gIH1cblxuICBwcm92aWRlQm9va1JlYWRlclJlcG8uLmJvb2tSZWFkZXJcblxuICBtYWluLi5wcm92aWRlQm9va1JlYWRlclJlcG9cbiAgbWFpbi4uPkJvb2tSZWFkZXJSZXBvIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQiLCJ0aGVtZVZhcmlhYmxlcyI6eyJiYWNrZ3JvdW5kIjoid2hpdGUiLCJwcmltYXJ5Q29sb3IiOiIjRUNFQ0ZGIiwic2Vjb25kYXJ5Q29sb3IiOiIjZmZmZmRlIiwidGVydGlhcnlDb2xvciI6ImhzbCg4MCwgMTAwJSwgOTYuMjc0NTA5ODAzOSUpIiwicHJpbWFyeUJvcmRlckNvbG9yIjoiaHNsKDI0MCwgNjAlLCA4Ni4yNzQ1MDk4MDM5JSkiLCJzZWNvbmRhcnlCb3JkZXJDb2xvciI6ImhzbCg2MCwgNjAlLCA4My41Mjk0MTE3NjQ3JSkiLCJ0ZXJ0aWFyeUJvcmRlckNvbG9yIjoiaHNsKDgwLCA2MCUsIDg2LjI3NDUwOTgwMzklKSIsInByaW1hcnlUZXh0Q29sb3IiOiIjMTMxMzAwIiwic2Vjb25kYXJ5VGV4dENvbG9yIjoiIzAwMDAyMSIsInRlcnRpYXJ5VGV4dENvbG9yIjoicmdiKDkuNTAwMDAwMDAwMSwgOS41MDAwMDAwMDAxLCA5LjUwMDAwMDAwMDEpIiwibGluZUNvbG9yIjoiIzMzMzMzMyIsInRleHRDb2xvciI6IiMzMzMiLCJtYWluQmtnIjoiI0VDRUNGRiIsInNlY29uZEJrZyI6IiNmZmZmZGUiLCJib3JkZXIxIjoiIzkzNzBEQiIsImJvcmRlcjIiOiIjYWFhYTMzIiwiYXJyb3doZWFkQ29sb3IiOiIjMzMzMzMzIiwiZm9udEZhbWlseSI6IlwidHJlYnVjaGV0IG1zXCIsIHZlcmRhbmEsIGFyaWFsIiwiZm9udFNpemUiOiIxNnB4IiwibGFiZWxCYWNrZ3JvdW5kIjoiI2U4ZThlOCIsIm5vZGVCa2ciOiIjRUNFQ0ZGIiwibm9kZUJvcmRlciI6IiM5MzcwREIiLCJjbHVzdGVyQmtnIjoiI2ZmZmZkZSIsImNsdXN0ZXJCb3JkZXIiOiIjYWFhYTMzIiwiZGVmYXVsdExpbmtDb2xvciI6IiMzMzMzMzMiLCJ0aXRsZUNvbG9yIjoiIzMzMyIsImVkZ2VMYWJlbEJhY2tncm91bmQiOiIjZThlOGU4IiwiYWN0b3JCb3JkZXIiOiJoc2woMjU5LjYyNjE2ODIyNDMsIDU5Ljc3NjUzNjMxMjglLCA4Ny45MDE5NjA3ODQzJSkiLCJhY3RvckJrZyI6IiNFQ0VDRkYiLCJhY3RvclRleHRDb2xvciI6ImJsYWNrIiwiYWN0b3JMaW5lQ29sb3IiOiJncmV5Iiwic2lnbmFsQ29sb3IiOiIjMzMzIiwic2lnbmFsVGV4dENvbG9yIjoiIzMzMyIsImxhYmVsQm94QmtnQ29sb3IiOiIjRUNFQ0ZGIiwibGFiZWxCb3hCb3JkZXJDb2xvciI6ImhzbCgyNTkuNjI2MTY4MjI0MywgNTkuNzc2NTM2MzEyOCUsIDg3LjkwMTk2MDc4NDMlKSIsImxhYmVsVGV4dENvbG9yIjoiYmxhY2siLCJsb29wVGV4dENvbG9yIjoiYmxhY2siLCJub3RlQm9yZGVyQ29sb3IiOiIjYWFhYTMzIiwibm90ZUJrZ0NvbG9yIjoiI2ZmZjVhZCIsIm5vdGVUZXh0Q29sb3IiOiJibGFjayIsImFjdGl2YXRpb25Cb3JkZXJDb2xvciI6IiM2NjYiLCJhY3RpdmF0aW9uQmtnQ29sb3IiOiIjZjRmNGY0Iiwic2VxdWVuY2VOdW1iZXJDb2xvciI6IndoaXRlIiwic2VjdGlvbkJrZ0NvbG9yIjoicmdiYSgxMDIsIDEwMiwgMjU1LCAwLjQ5KSIsImFsdFNlY3Rpb25Ca2dDb2xvciI6IndoaXRlIiwic2VjdGlvbkJrZ0NvbG9yMiI6IiNmZmY0MDAiLCJ0YXNrQm9yZGVyQ29sb3IiOiIjNTM0ZmJjIiwidGFza0JrZ0NvbG9yIjoiIzhhOTBkZCIsInRhc2tUZXh0TGlnaHRDb2xvciI6IndoaXRlIiwidGFza1RleHRDb2xvciI6IndoaXRlIiwidGFza1RleHREYXJrQ29sb3IiOiJibGFjayIsInRhc2tUZXh0T3V0c2lkZUNvbG9yIjoiYmxhY2siLCJ0YXNrVGV4dENsaWNrYWJsZUNvbG9yIjoiIzAwMzE2MyIsImFjdGl2ZVRhc2tCb3JkZXJDb2xvciI6IiM1MzRmYmMiLCJhY3RpdmVUYXNrQmtnQ29sb3IiOiIjYmZjN2ZmIiwiZ3JpZENvbG9yIjoibGlnaHRncmV5IiwiZG9uZVRhc2tCa2dDb2xvciI6ImxpZ2h0Z3JleSIsImRvbmVUYXNrQm9yZGVyQ29sb3IiOiJncmV5IiwiY3JpdEJvcmRlckNvbG9yIjoiI2ZmODg4OCIsImNyaXRCa2dDb2xvciI6InJlZCIsInRvZGF5TGluZUNvbG9yIjoicmVkIiwibGFiZWxDb2xvciI6ImJsYWNrIiwiZXJyb3JCa2dDb2xvciI6IiM1NTIyMjIiLCJlcnJvclRleHRDb2xvciI6IiM1NTIyMjIiLCJjbGFzc1RleHQiOiIjMTMxMzAwIiwiZmlsbFR5cGUwIjoiI0VDRUNGRiIsImZpbGxUeXBlMSI6IiNmZmZmZGUiLCJmaWxsVHlwZTIiOiJoc2woMzA0LCAxMDAlLCA5Ni4yNzQ1MDk4MDM5JSkiLCJmaWxsVHlwZTMiOiJoc2woMTI0LCAxMDAlLCA5My41Mjk0MTE3NjQ3JSkiLCJmaWxsVHlwZTQiOiJoc2woMTc2LCAxMDAlLCA5Ni4yNzQ1MDk4MDM5JSkiLCJmaWxsVHlwZTUiOiJoc2woLTQsIDEwMCUsIDkzLjUyOTQxMTc2NDclKSIsImZpbGxUeXBlNiI6ImhzbCg4LCAxMDAlLCA5Ni4yNzQ1MDk4MDM5JSkiLCJmaWxsVHlwZTciOiJoc2woMTg4LCAxMDAlLCA5My41Mjk0MTE3NjQ3JSkifX0sInVwZGF0ZUVkaXRvciI6ZmFsc2V9)

# Example
## Problem
Suppose we have built an ecommerce book store. Our website is getting bigger and have traffic much more than ever. We,
of course, need to scale our application to handle new demands coming in. The team has decided that we should add cache
when we're fetching the book-list. Let's just say the team decided to use Redis.

Our team is still not really sure though if our cache choice is the best one. So we'd need to ensure that when we need
to change our cache in the future, it won't break any other logics. We want to separate the caching logics and processes
from everything else.

Let's take a step back and see the big picture here. We currently already have a module that fetches book-list from
PostgreSQL DB. Our current intention is just adding another layer of cache-logic to ease our Postgres load. If only we can
somehow wrap our DB logic with additional process without interfering with it.

## Solution
Yes we can, that's what Decorator pattern does. First, we need to define a common contract which unites both Postgres
and Redis process. If we look carefully at both Redis and Postgres, their business intention is actually the same, it is to
fetch the list of book. That's set then, we'll build an interface around that intention.
```go
type BookReaderRepo interface {
	ListBook() ([]entity.Book, error)
}
```

Then we'll start coding our Postgres processes like usual to fulfill the interface [here](https://github.com/gentra/golang-design-patterns/blob/master/decorator/internal/repo/bookreaderrepo/postgrerepo.go).

With Postgres done, now we just need to think about the Redis. Our Redis has **the same intention of fetching list of book**,
but since it's not the source-of-truth, **it needs to fetch list of book from Postgres if data doesn't exist** yet on Redis.
You see from both bold texts before that the intentions are actually the same. We can use the same contract `BookReaderRepo`.
We just need to structure it in a way that Redis implementation calls the Postgres one when needed.

```go
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
	log.Print("Checking if data is cached on Redis")
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
```

With all core logics done, we just need to set up a constructor which constructs both Redis and Postgres in a combined
manner.

```go
func BookReaderRepo(useCache bool) internal.BookReaderRepo {
	db := bookreaderrepo.NewPostgreRepo()
	switch useCache {
	case true:
		return bookreaderrepo.NewRedisRepo(db) // we wrap the source-of-truth in the cache class
	default:
		return db
	}
}
```
Then execute it
```go
useCache := true // Let's just pretend that you get this use-cache configuration from config files or cli param
repo := provide.BookReaderRepo(useCache)

books, err := repo.ListBook()
if err != nil {
    log.Println(err)
}
```

As you can see from the code above, we are even now able to turn on/off our cache dynamically without polluting our core
codes, all the toggle happens on a constructor. Much cleaner. In real world case, the on/off configuration might reside
on a configuration file, ENV, or feature-toggle. It's usually useful in some surge scenarios or debugging purposes.

Let's get back to our problem statement above a bit.
> Our team is still not really sure, though if our cache choice is the best one.

Let's say we find a better cache option, we can implement a new cache-option just by simply implementing `BookReaderRepo`
interface. Then if we want to dynamically switch the caches for experiment, we can put a [strategy pattern](https://github.com/gentra/golang-design-patterns/tree/master/strategy)
on it.