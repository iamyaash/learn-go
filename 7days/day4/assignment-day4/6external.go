package main

import "fmt"

func makeMultiplier(n int) func(int) int {
    // Your code here
    return func(x int) int {
	    return x * n
	}

}

func main() {
    // Create multipliers and test them with different inputs
	multiplierBy2 := makeMultiplier(2)
    	multiplierBy3 := makeMultiplier(3)
	
	fmt.Println(multiplierBy2(5))
	fmt.Println(multiplierBy3(9))
}
