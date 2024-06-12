package main

import "fmt"

func runningAverage() func(float64) float64 {
    // Initialize sum and count
    sum := 0.0
    count := 0.0

    // Return a closure that calculates the running average
    return func(num float64) float64 {
        // Add the new number to the sum
        sum += num
        // Increment the count
        count++
        // Calculate the running average
        avg := sum / count
        // Return the running average
        return avg
    }
}

func main() {
    // Create the runningAverage closure
    avg := runningAverage()

    // Test the closure with a series of numbers and print the running average after each addition
    fmt.Println("Running Average:")
    fmt.Println("Add 5:", avg(5))
    fmt.Println("Add 10:", avg(10))
    fmt.Println("Add 15:", avg(15))
}

