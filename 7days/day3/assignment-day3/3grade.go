package main

import ("fmt")

func main() {
	mark := 99
	if mark >= 90 {
		fmt.Println("A")
	} else if mark >= 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
