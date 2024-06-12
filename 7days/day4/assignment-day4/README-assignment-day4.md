# Day 4 -Assignments

# Functions, Anonymous Functions, and Closures in Go

## Functions

1. **Basic Function Definition and Call:**
   - Write a function `greet` that takes a `name` (string) as a parameter and prints a greeting message `"Hello, <name>!"`.
   - Call the `greet` function with different names and observe the output.

   ```go
   func greet(name string) {
       // Your code here
   }

   func main() {
       // Call the greet function with different names
   }
   ```

2. **Function with Multiple Parameters and Return Values:**
- Write a function `swap` that takes two integers and returns them in swapped order.
- Call the `swap` function and print the results.
```go
func swap(a int, b int) (int, int) {
    // Your code here
}

func main() {
    // Call the swap function and print the results
}
```

3. **Simple Anonymous Functions:**
- Define an anonymous function that takes two integers, adds them, and returns the result.
- Immediately invoke this anonymous function and print the result.

```go
func main() {
    // Define and immediately invoke the anonymous function
}
```

4. **Assigning Anonymous Function to a Variable:**
- Assign an anonymous function that multiplies two integers to a variable `multiply`.
- Use this `multiply` function to multiply different pairs of integers and print the results.
```go
func main() {
    // Define the anonymous function and assign it to a variable
    // Use the multiply function with different pairs of integers
}
```

## Closures

5. **Simple Closure Example:**

- Write a function makeCounter that returns a closure. The closure should increment and return an internal counter each time it is called.
- Create two counters using makeCounter and demonstrate that they maintain independent counts.
```go
func makeCounter() func() int {
    // Your code here
}

func main() {
    // Create two counters and demonstrate their independence
}
```

6. **Closure with External Variables:**
- Write a function makeMultiplier that takes an integer n and returns a closure. The closure should multiply its argument by n.
- Use makeMultiplier to create a multiplier by 2 and a multiplier by 3. Test these multipliers with different inputs.
```go
func makeMultiplier(n int) func(int) int {
    // Your code here
}

func main() {
    // Create multipliers and test them with different inputs
}
```

7. **Combining Functions and Closures:**
- Write a function makeSequence that takes a starting number and a step size. It should return a closure that generates the next number in the sequence each time it is called.
- Use makeSequence to create sequences starting from 0 with step size 1, and from 10 with step size 5. Generate and print several numbers from each sequence.
```go
func makeSequence(start int, step int) func() int {
    // Your code here
}

func main() {
    // Create sequences and generate numbers from them
}
```

8. **Creating a Closure to Calculate Running Average:**
   - Write a function `runningAverage` that returns a closure. The closure should calculate and return the running average of numbers passed to it.
   - The closure should maintain two variables: `sum` to store the sum of all numbers seen so far, and `count` to store the count of numbers seen so far.
   - Test the closure with a series of numbers and print the running average after each addition.

   ```go
   func runningAverage() func(float64) float64 {
       // Your code here
   }

   func main() {
       // Create the runningAverage closure and test it with a series of numbers
   }