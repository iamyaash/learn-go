package main

import "fmt"

type Car struct{
	brand string
	model string
	year int
}

func main() {
	var veh1 = Car{
		brand:"BMW",
		model:"M5",
		year:2027,
	}
	printCar(veh1)
}

func printCar(veh Car) {
	fmt.Println(veh.brand)
	fmt.Println(veh.model)
	fmt.Println(veh.year)
}
