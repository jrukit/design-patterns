package main

type Fifo struct {
	order []string
}

func NewFifo() *Fifo { return &Fifo{} }

func (f *Fifo) OnGet(key string) {}

func (f *Fifo) OnAdd(key string) {
	f.order = append(f.order, key)
}

func (f *Fifo) Evict() string {
	victim := f.order[0]
	f.order = f.order[1:]
	return victim
}
