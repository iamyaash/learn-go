# Day 3: Conditions, Loops & Switch Statements

**Day 3**, we'll delve deep into **conditions**, **loops** & **switch statement** and only focus on these topics for this day. To make you comfortable working with these topics. I'll explain how these are executed in runtime, how the logic behind it will work and some tips & tricks. 


## What to Expect
  - **Conditions:** `if`, `else if` & `else`.
  - **Loops:** `for` only looping construct.(***(but various method are available to use this)***.
  - **Switch Statements:** `switch` expressing conditions across many branches


## Conditions

Using conditions in `Go` is straightforward and easy to implement. You can combine conditions with logical operators to make more complex checks. Note that in Go, there is no need to use parentheses to declare conditions or initializations.

#### if -only
```go
package main
import ("fmt")
func main() {
    var num int = 5 // num := 5
    if num < 10 {
        fmt.Println("Less than 10")
    }
}
```

#### if-else

```go
package main
import ("fmt")
func main() {
    if 2%2 == 0{
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }
}
```

#### Conditions with Logical operators

```go
num := 8
if num == 8 || num == 17 {
    fmt.Println("Your Num")
}
```

#### Preceding the Conditional Checks
```go
if num := 8; num < 0 {
    // statements
}
```
> When the `if` statement is executed, `num` is declared and initialized to `8`.

> The variable `num` is within the scope of the entire `if-else if-else` block. This means that num can be accessed and used in any of the conditional branches.

#### if -> else-if -> else

```go
package main 
import ("fmt")
func main() {
    if num:= 8; num < 8 {
        fmt.Println("Less than 8")
    } else if num == 8 {
        fmt.Println("Equal to 8")
    } else {
        fmt.Println("Greater than 8")
    }
}
```

## Loops

Like in most programming languages, `Go` allow you to repeat a block of code multiple times until a certain condition is met.

`Go` supports the `for loop`, which is *versatile* and can be *modified* to behave like a `while loop`. Additionally, `Go` introduces the concept of a `range loop`, which uses the range keyword to iterate over arrays, slices, strings, or channels."

#### for loop: 

The for loop in Go is similar to other languages, however it doesn't require paranthesis`()` & curly braces `{}` are mandatory.

##### Basic with single condition:
```go
package main
import ("fmt")
func main() {
    i := 1              //initialized i=1
    for i <= 3 {        //loop checks i<>=3
        fmt.Println(i)  //only print if true,
        i = i + 1       // increment||update
    }
}
```
>Notice: It behaves like `while loop`, hence we implement while loop using `for loop` in `Go`.

##### Enhanced for loop, but standard way
```go
package main
import ("fmt")
func main() {
    // i=0     - initialization
    // i < 10  - condition
    // i++     - update
    for i=0; i < 10; i++ {
        fmt.Println("i")
    }
}
```
> Notice: This can be considered the actual way of implementing `for loop` in other languages.

##### Infinite looping
```go
// loop repeatedly until break
i := 0
for {
        fmt.Println(i)
        i++
        if i == 5 {
            break
        }
}
```

#### range loop:
The `range loop` in Go is used to iterate over elements of a sequence such as an array, slice, string, or map.

```go
package main
import ("fmt")
func main() {
    number := []int{1, 2, 3, 4, 5}
    for index, value := range numbers {
        fmt.Println("Index: %d, Value: %d\n", index, value)
    }
}
```
This loop iterates over each element of the `numbers` slice. `index` represents the index of the current element and  `value` represents the value of the current element.

## Switch Statements

The `switch` statements provides a way to execute different blocks of code based on the value of an expression. It's similar to other languages but slightly different approach.
#### Basic switch conditions
```go
package main
import ("fmt")

func main() {
    fruit := "apple"
    switch fruit {
        case "apple":
        fmt.Println("It's an appl")
        
        case "banana":
        fmt.Println("It's a banana")
        
        default:
        fmt.Println("Idk bruh")
    }
}
```
In the above example,
- `var fruit string = "apple"` is an untyped variable declaration.
- The `switch` statement checks each `case` value again the value of `fruit`.
  - i,e. case == fruit
- If match is found, the respective statement will be executed.
- Otherwise, `default` block statement is executed.

#### Switch statements with multiple conditions
Here we can declare multple conditions in `case`,
```go
switch num {
    case 1, 2:
    fmt.Println("One or two")
    case 3, 4:
    fmt.Println("three or four")
    default:
    fmt.Prinln("idk bruh")
}
```

#### Switch without expression.
Here, we won't specify expression in `switch` statement.

```go
fruit := "apple"
switch {
    case fruit == "apple":
    // do something
    case fruit == "orange":
    // do something
    default:
    // no choice, then execute this
}
```

#### Switch with `Fallthrough`
The `fallthrough` statement is used to execute the next case block even if it doesn't match the expression.

In simple terms, it won't stop after executing one case statement, rather contrinue to execute next case, even if it's false.

```go
num := 1
switch num {
    case 1:
    // do something
    fallthrough 
}
```


# [Day 3 -Assignments [Click Here👆]](./assignment-day3/README-assignment-day3.md)