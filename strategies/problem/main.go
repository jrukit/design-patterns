package main

import (
	"fmt"
	"math"
)

type Cache struct {
	storage      map[string]string
	evictionAlgo string
	capacity     int
	maxCapacity  int

	// lru
	lastUsed map[string]int
	clock    int

	// lfu
	useCount map[string]int

	//fifo
	order []string
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}

	c.capacity++
	c.storage[key] = value

	switch c.evictionAlgo {
	case "lfu":
		c.useCount[key] = 0
	case "fifo":
		c.order = append(c.order, key)
	case "lru":
		c.clock++
		c.lastUsed[key] = c.clock
	}
}

func (c *Cache) evict() {
	var victim string
	switch c.evictionAlgo {
	case "lfu":
		least := math.MaxInt
		for k, n := range c.useCount {
			if n < least {
				victim, least = k, n
			}
		}
		delete(c.useCount, victim)
	case "fifo":
		victim = c.order[0]
		c.order = c.order[1:]
	case "lru":
		oldest := c.clock
		for k, t := range c.lastUsed {
			if t < oldest {
				victim, oldest = k, t
			}
		}
		delete(c.lastUsed, victim)
	}
	delete(c.storage, victim)
	c.capacity--
}

func (c *Cache) get(key string) string {
	switch c.evictionAlgo {
	case "lfu":
		c.useCount[key]++
	case "lru":
		c.clock++
		c.lastUsed[key] = c.clock
	}
	return c.storage[key]
}

func (c *Cache) setEvictionAlgo(e string) {
	c.evictionAlgo = e
}

func main() {
	cache := Cache{
		storage:      make(map[string]string),
		evictionAlgo: "fifo",
		capacity:     0,
		maxCapacity:  2,
		lastUsed:     make(map[string]int),
		clock:        0,
		useCount:     make(map[string]int),
	}

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")
	fmt.Println("[fifo] data: ", cache.storage)

	cache.setEvictionAlgo("lru")
	cache.get("b")
	cache.get("b")
	cache.get("b")
	cache.get("b")
	cache.get("c")
	cache.add("d", "4")
	fmt.Println("[lru] data: ", cache.storage)

	cache.setEvictionAlgo("lfu")
	cache.get("c")
	cache.get("c")
	cache.get("d")
	cache.add("e", "5")
	fmt.Println("[lfu] data: ", cache.storage)
}
