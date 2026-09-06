package main

type Barista struct {
	Orders []Order
}

func (b *Barista) MakeCoffee(order Order) {
	// Make the coffee
}

func (b *Barista) ServeCoffee(order Order) {
	// Serve the coffee
}
