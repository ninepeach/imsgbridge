package contact

type Resolver interface {
	Resolve(
		handle string,
	) Identity
}

type resolver struct {
	cache *Cache
}

func NewResolver() Resolver {

	return &resolver{

		cache: NewCache(),
	}
}

func (r *resolver) Resolve(
	handle string,
) Identity {

	// 1. cache

	if id, ok := r.cache.Get(handle); ok {

		return id
	}

	// 2. fallback

	id := Identity{

		Handle: handle,

		Name: handle,
	}

	r.cache.Set(id)

	return id
}
