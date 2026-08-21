package contact

import "sync"

type Cache struct {
	mu sync.RWMutex

	data map[string]Identity
}

func NewCache() *Cache {

	return &Cache{

		data: make(
			map[string]Identity,
		),
	}

}

func (c *Cache) Get(
	handle string,
) (
	Identity, bool,
) {

	c.mu.RLock()

	defer c.mu.RUnlock()

	id, ok := c.data[handle]

	return id, ok
}

func (c *Cache) Set(
	id Identity,
) {

	c.mu.Lock()

	defer c.mu.Unlock()

	c.data[id.Handle] = id

}

func (c *Cache) Clear() {

	c.mu.Lock()

	defer c.mu.Unlock()

	c.data = make(
		map[string]Identity,
	)

}
