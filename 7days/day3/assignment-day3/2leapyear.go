package main

import ("fmt")

func main() {
	year := 2023
	if year%4 == 0 && year%100 != 0 {
		fmt.Println("Leap")
	} else {
		fmt.Println("Non-Leap")
	}
}
