package main

type Lru struct {
	lastUsed map[string]int
	clock    int
}

func NewLru() *Lru { return &Lru{lastUsed: map[string]int{}} }

func (lr *Lru) OnGet(key string) {
	lr.clock++
	lr.lastUsed[key] = lr.clock
}

func (lr *Lru) OnAdd(key string) {
	lr.clock++
	lr.lastUsed[key] = lr.clock
}

func (lr *Lru) Evict() string {
	victim, oldest := "", lr.clock
	for k, t := range lr.lastUsed {
		if t < oldest {
			victim, oldest = k, t
		}
	}
	delete(lr.lastUsed, victim)
	return victim
}
