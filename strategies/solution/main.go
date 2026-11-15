package main

import "fmt"

func run(name string, algo EvictionAlgo) {
	cache := NewCache(3, algo)

	cache.Add("a", "1")
	cache.Add("b", "2")
	cache.Get("b")
	cache.Get("b")
	cache.Add("c", "3")
	cache.Get("a")

	cache.Add("d", "4")

	fmt.Printf("%-4s → %v\n", name, cache.storage)
}

func main() {
	run("LRU", NewLru())
	run("LFU", NewLfu())
	run("FIFO", NewFifo())
}
