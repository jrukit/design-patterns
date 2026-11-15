package main

import "math"

type Lfu struct {
	useCount map[string]int
}

func NewLfu() *Lfu { return &Lfu{useCount: map[string]int{}} }

func (l *Lfu) OnGet(key string) {
	l.useCount[key]++
}

func (l *Lfu) OnAdd(key string) {
	l.useCount[key] = 0
}

func (l *Lfu) Evict() string {
	victim, least := "", math.MaxInt
	for k, n := range l.useCount {
		if n < least {
			victim, least = k, n
		}
	}
	delete(l.useCount, victim)

	return victim
}
