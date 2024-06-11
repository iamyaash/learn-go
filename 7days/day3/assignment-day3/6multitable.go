package main

import ("fmt")

func main() {
	var num int = 10
	var i int

	for i = 1; i <= num; i++ {
        	fmt.Printf("%d x %d = %d\n", i, num, i*num)
	}
}
