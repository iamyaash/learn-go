package main
import "fmt"

func main() {
	var myslice = []int{11,12,13,14,15,16,17,18,19,20}
	fmt.Println("Before: ",myslice)

	myslice = append(myslice, 21, 22, 23)
	fmt.Println("After: ", myslice)
}

