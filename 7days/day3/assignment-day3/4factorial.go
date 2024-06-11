package main

import ("fmt")

func main() {
	var num int = 4
	var result int = 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	fmt.Println(result)
}
