package main

type Cache struct {
	storage     map[string]string
	capacity    int
	maxCapacity int
	algo        EvictionAlgo
}

func NewCache(maxCapacity int, algo EvictionAlgo) *Cache {
	return &Cache{
		storage:     map[string]string{},
		capacity:    0,
		maxCapacity: maxCapacity,
		algo:        algo,
	}
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		delete(c.storage, c.algo.Evict())
		c.capacity--
	}

	c.capacity++
	c.storage[key] = value

	c.algo.OnAdd(key)
}

func (c *Cache) Get(key string) string {
	c.algo.OnGet(key)
	return c.storage[key]
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.algo = e
}
