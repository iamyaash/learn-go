package main

import "fmt"

func main() {
	func(x int, y int) (int) {
		sum := x + y
		fmt.Println(sum)
		return sum
	}(3, 5)
}
