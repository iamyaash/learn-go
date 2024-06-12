package main

import "fmt"

func makeSequence(start int, step int) func() int {
    // Your code here
    current := start
    return func() int {
	    result := current
	    current += step
	    return result
    }
}

func main() {
    // Create sequences and generate numbers from them
    s1 := makeSequence(0,1)
    s2 := makeSequence(10, 5)

    for i := 0; i < 5; i++ {
    	fmt.Println(s1())
    }

    for i := 0; i < 5; i++ {
	fmt.Println(s2())
    }

}
