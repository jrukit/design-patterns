package main

type EvictionAlgo interface {
	OnAdd(key string)
	Evict() string
	OnGet(key string)
}
