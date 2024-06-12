package main

import "fmt"

func makeCounter() func() int {
	adder := func() func(int) int {
		val := 0
		return func(x int) int {
			val += x
			return val
		}
	}
	// Create a new counter by calling adder
	val := adder()

	// Return a function that returns the current value of the counter
	return func() int {
		return val(1) // Increment the counter by 1 and return the new value
	}
}

func main() {
	// Create two counters and demonstrate their independence
	counter1 := makeCounter()
	counter2 := makeCounter()

	fmt.Println("Counter 1:", counter1())
	fmt.Println("Counter 1:", counter1())
	fmt.Println("Counter 2:", counter2())
	fmt.Println("Counter 1:", counter1())
}

