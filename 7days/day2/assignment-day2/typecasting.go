package main

import("fmt")

func main() {
	var numInt int = 10
	var numFloat float64 = 5.5
	conv := float32(float64(numInt) + numFloat)

	fmt.Println(conv)
}
