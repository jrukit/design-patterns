package main

type TomatoTopping struct {
	pizza IPizza
}

func (p *TomatoTopping) getPrice() int {
	return p.pizza.getPrice() + 7
}
