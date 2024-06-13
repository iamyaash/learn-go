package main

import "fmt"

func main() {
	categories := make(map[string][]string)

	categories["Laptop"] = []string{"Compact", "Travel", "Less Power"}
	categories["Desktop"]= []string{"Bigger", "Not Out", "More Power"}
	for category, items := range categories {
        	fmt.Printf("Category: %s\n", category)
        	for _, item := range items {
            	fmt.Println(item)
        	}
        fmt.Println()
    }	
}
