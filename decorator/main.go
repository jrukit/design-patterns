package main

import "fmt"

func main() {
	pizza := &VeggieMania{}

	// Add cheese topping
	pizzaWithChesse := &CheeseTopping{
		pizza: pizza,
	}
	fmt.Printf("Price of veggeMania with cheese topping is %d\n", pizzaWithChesse.getPrice())

	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithChesse,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
