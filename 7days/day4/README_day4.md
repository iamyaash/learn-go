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

#### Function call with single parameter

```go
package main
import ("fmt")
func printMsg (name string) {
    fmt.Println("Hey, ", name)
}
func main() {
    printMsg("Yaash")
}
// Hey, Yaash <-output
```

#### Function call with multiple parameters
##### Same Type

```go
package main
import ("fmt")
func addSum(num1 int, num2 int) {
    var sum int = num1 + num2
    fmt.Println("Sum of ", num1, " & ", num2, "is", sum)
}

func main () {
    addSum(1, 2)
    addSum(4, 4)
} 
```
##### Different Type 

```go
package main
import ("fmt")
func doSome (name string, num int) {
    fmt.Println("Name:", name)
    fmt.Println("Fav:", num)
    // fmt.Println("Name: ", name, "\n", "Favorite Number: ", num)
}
func main() {
    doSome("Yaash", 8)
}
```

### Return Values of Function
> Note: `Printing` & `Returning`, are not the same.

To return a value, you need to define the `return` keyword to return value.
> Return Value: `int`, `string`, `bool`, `float`, ..etc

##### Syntax
```go
func FunctionName (param1 type, param2 type) {
    return output
}
```


#### Function with Return Values
```go
package main
import ("fmt")
func doSomething(num1 int32, num2 int32) {
    return num1 + num2
}
function main() {
    doSomething(4, 4)
}
```
> Note: This only returns the value, but it doesn't print the value on output.


#### Function with Name Return Values
```go
func doSomething(n1 int, n2 int) {
    var result int = n1 + n2
    return result // returns the value of var result

}
func main() {
    doSomething(4, 4)
}
```
> Here we return the value of `var`.

#### Function with Multiple Return Values

```go
package main
import("fmt")
func myFunction(x int, y string) (result int, txt1 string){
    result = x + x
    txt1 = y + " World!"

    // Return the named return values
    return
}

func main() {
    fmt.Println(myFunction(5, "Hello"))
}
```

```go
func main() {
    a, b := myFunction(5, "Hello") // stored in var
    fmt.Println(a, b) // accesed using var
}
// to omit any var (a | b)
func main() {
    _, b = myFunction(5, "Hello")
    fmt.Println(b)
}
```

## Anonymous Function

An `anonymous function` is a function that doesn't have a name. It is particularly useful when you need to create an **inline function** that is only used *temporarily*, often **passed as an argument to higher-order functions** or used for a **one-time operation**.

##### Syntax
```go
// Anonymous Function
func () {
    fmt.Println("It's an Anonymous Function)
}()
```

##### Example
```go
package main
import ("fmt")
func main() {
    result := func(x int , y string)(int, string){
        res := x + x
        txt := y + "Hey"
        return res, txt
    }(5, "Hello")
}
```

#### Key Points
1. **No Name**: The function does not have a name.
2. **Inline Definition and Execution**: It is defined and executed inline.
3. **Temporary Usage**: It is used for temporary purposes where a named function would be unnecessary.
4. **Function Parameters and Return Values**: Like named functions, anonymous functions can accept parameters and return values.

> In simple terms, anonymous function works like regular function but a temporarily used.
