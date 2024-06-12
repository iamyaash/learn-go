# Day 4: Defining/calling functions, multiple/named return values, anonymous functions, closures 

A function is block of statements that can be repeatedly used in a program. Function won't be executed unless they're called and it is executed by a `call` to the function.

## Functions
### Rules for naming functions in `Go`

- Must begin with letter **(a-zA-Z)**.

- Function name only contains **alpha-numeric(a-zA-z, 0-9)** and **_(underscore)**.
- No spaces are allowed either.
- Function names are case-sensitive as well.
```go
func hey()  //'h'
func Hey()  //'H'
// both functions are different
```

### Creating a Function
##### Syntax:
```go
func FunctionName(){
    //statements to be executed
}
```
- `func` *is the keyword to create a function.*
- `()-paranthesis` is used to pass parameter, we'll look into it later.
- Function will execute what is inside the`{}-curly braces`.


### Calling a Function
As mention earlier, a function **can't be executed automatically** and it needs to be **called to get executed**.
#### Function call with no parameters
```go
package main
import ("fmt")
func printMsg(){
    fmt.Println("print msg")
}

func main() {
    printMsg() // function call
}
```
In the above example,
- The `printMsg()` function is defined with **no parameters** and it prints the *string*, when called.
- The `printMsg` is called within the `main` function. This triggers the execution of `printMsg`. And it prints the string.
- Calling the `printMsg` function multiple times inside the `main` function will execute the `printMsg` function each time it's called, resulting in it being executed *N times*.


#### Function call with parameters
##### Syntax:
```go
func FunctionName(param1 type, param2 type) {
    //statements to be executed
}
```
- Parameters/Arguments are specified inside `paranthesis()`.
- Parameters act as values for the variables inside the function.