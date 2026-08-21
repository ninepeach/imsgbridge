package contact

import "errors"

type Resolver interface {
	Resolve(
		handle string,
	) (Identity, error)
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
) (Identity, error) {

	if handle == "" {

		return Identity{},
			errors.New("empty handle")
	}

	if id, ok := r.cache.Get(handle); ok {

		return id, nil

	}

	// TODO:
	// macOS Contacts lookup

	id := Identity{

		Handle: handle,

		Name: handle,

		Type: detectType(handle),
	}

	r.cache.Set(id)

	return id, nil
}

func detectType(
	handle string,
) string {

	for _, c := range handle {

		if c == '@' {

			return "email"

		}
	}

	if len(handle) > 0 && handle[0] == '+' {

		return "phone"

	}

	return "unknown"
}
