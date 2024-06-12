package main

import "fmt"

func makeCounter() func() int {
	// Your code here
	adder := func() func(int) int {
		val := 0
		return func(x int) int{
			val += x
			return val
		}
	}
	val := adder()
	
	return func() int {
		return val(1)
	}

}

func main() {
    // Create two counters and demonstrate their independence
	c1 := makeCounter()
	c2 := makeCounter()
	
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c2())
}
