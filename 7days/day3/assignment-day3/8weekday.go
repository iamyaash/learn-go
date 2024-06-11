package main 

import ("fmt")

func main () {
	var week int = 5
	switch week{
	case 1,2,3,4,5:
		fmt.Println("Weekday")
	case 6,7:
		fmt.Println("Weekend")
	default:
		fmt.Println("idk bruh")
	}
}
