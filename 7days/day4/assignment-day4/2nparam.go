package main

import "fmt"

func swap(a int, b int) (int, int) {
    // Store the value of a in a temporary variable
    temp := a
    // Assign the value of b to a
    a = b
    // Assign the value of temp to b
    b = temp
    // Print the swapped values
    return a, b
}


func main() {
    // Call the swap function and print the results
    n, m := swap(4, 8)
    fmt.Println(n, m)
}
