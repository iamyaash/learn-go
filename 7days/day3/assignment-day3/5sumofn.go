package main

import ("fmt")

func main() {
	var num int = 4
	var results int

	for i := 1; i <= num; i++ {
		results += i
	}
	fmt.Println(results)
}
