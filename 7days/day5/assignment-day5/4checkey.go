package main

import "fmt"

func main() {
	var item = make(map[string]int)
	item["Sugar"] = 8
	item["Cheese"] = 17
	item["Rice"] = 26

	if item, exist := item["Paneer"]; exist {
		fmt.Println(item)
	} else {
		fmt.Println("No dude")
	}
}
