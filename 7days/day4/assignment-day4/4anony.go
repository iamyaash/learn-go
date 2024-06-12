package main

import "fmt"

func main() {
	multiply := func(num1 int, num2 int) (int) {
		sum := num1 * num2
		return sum
	}(4, 2)
	fmt.Println(multiply)
}

