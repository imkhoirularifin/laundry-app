package main

import "laundry-app/delivery"

func main() {
	/*
		Starting server	with 2 available options:
		1. Chaiining (Method receiving syntax)
		2. Traditional
	*/

	// 1
	// delivery.NewServer().Start()

	// 2
	server := delivery.NewServer()
	delivery.Start(server)

	// End of options
}
