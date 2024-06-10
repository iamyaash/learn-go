# Day 2: Basic Syntax, Data Types, and Type Conversion

In day 2, we'll cover only `Basic Syntax`, `Data Types` and `Type Conversion`. 

## What to Expect
- **Basic Syntax:** *Understanding the structure of a Go program.*
- **Data Types:** *Introduction to the available data types in Go and how to use them.*
- **Type Conversion:** *Exploring how Go handles type conversion.*

**Day 2** module will be more focussed on you getting comfortable with `Go`, adpating yourself to program and understanding the syntax.

You will also be introduced to the data types available in `Go`, and at the end of this `Day 2` you will know *how and where* to use those data types.

Additionally, we'll look into type `conversion` and how `Go` programming language handles  the type conversion.

## Details to Know

- `Go` is a **high-level programming language**, designed by Google in 2007.
- `Go` was created to overcome the complex syntax of `C++` & `C` and be simple & easy to read like `Python`.
- It became an **open-source** project since **2009**, and publicly released in **2012**.
- This is by far the most used programming language in **`Cloud Technologies`**


## Basic Syntax

Let's take a look at this `hello-world.go` example:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

Here, 
- Every `go` program is part of the `package`.  That is, we define it using `package` keyword. 
  - *Line 1: `package main` means, this program belongs to main package.*
- `import "fmt"`, let use import file inclued in `fmt` package. It is standard package provided  in `go` itself, when installing it.
- `func main() {}` is a function. 
  - `main()` inside paranthesis, parameters & arguments can be passed
  - Function inside `{}` curly braces will be executed.
- `fmt.Println()` is a statement inside the `main` function.
  - It used `fmt` pacakge, which specializes in formating string or text, printing in console.
  - `Println()`, is a function present in `fmt` which is used to print a line of text to the console, followed by a newline.

### Additional Points of Basic Syntax:
- **Comments:** Go supports both single-line & multi-line comments.
```go
// single line comment
package main // single-line comment alongside code
/*
multi-line comment
*/
```
- **Importing Package:** Mulitple packages can be imported by using paranthesis and mentioning package names in each line.
```go
import (
    "fmt"
    "packagename1"
    "packagename2"
)
```
- **No need for semicolon:** Unline `Java`, `C++` and `C`, `Go` doesn't require *semicolons* at the end of statements.

- **Variable Declaration**: In `Go`, variables can be declared using both `var` keyword or `:=` shorthand operator.
```go
// Declaration without initialization
var age int   
// Initialization later
age = 30      

// Declaration with initialization 
var rolno int = 69 

// Shorthand declaration with initialization
name := "John"
```
## Data Types:
## Basic Data Types:
### 1. Numeric Data Types:

#### Integer Data Types
Signed integer data type, it's size can vary however it'll atleast be `32-bits`.

| dtype  | Definition              | Example |
| ------ | ----------------------- | ------- |
| int8   | 8-bit Signed Integer    | None    |
| int16  | 16-bit Signed Integer   | None    |
| int32  | 32-bit Signed Integer   | None    |
| int64  | 64-bit Signed Integer   | None    |
| uint8  | 8-bit Unsigned Interger | None    |
| uint16 | 16-bit Unsigned Integer | None    |
| uint32 | 32-bit Unsigned Integer | None    |
| uint64 | 64-bit Unsigned Integer | None    |
> Note: Both `int` & `uint` contain same size, that's either 32-bits or 64-bits.
#### Floating Data Types
It's used to represent **floating-point numbers**. The floating-point numbers are **numbers that has a decimal point**, so that it can represent both **integer** & **fractional** parts.

##### float32 - occupies 32-bits in memory
- `float32` aka single-precision floating-point format.
- 7 decimal digits of precision.
- Suitable if precision is not critical.
##### float64 - occupies 64-bits in memory
- `float64` aka double-precision floating-point format.
- 16 decimal digits of precision.
- Suitable in mathematical operations where precision is important.

```go
var num1 float32 = 3.14
var num2 float64 = 3.141592653589793
```

**Usage**:
- `float32`: Used when memory usage is a concern or when precision up to **7 decimal places** is sufficient.
- `float64`: Provides higher precision and is generally preferred for most calculations requiring floating-point numbers.
   
#### Complex Data Types

- `complex64`: Represents complex numbers with float32 real and imaginary parts (each 32 bits).
- `complex128`: Represents complex numbers with float64 real and imaginary parts (each 64 bits).

These data types are used when you need to deal with **both real and imaginary numbers simultaneously**. They are often used in mathematical **calculations involving complex operations**.

```go
var c1 complex64 = 5 + 7i   // 32-bit real and imaginary parts
var c2 complex128 = 10 - 3i // 64-bit real and imaginary parts
```

**Usage**:
- The `complex64` type is useful when memory usage is a concern or when dealing with complex numbers that fit within 32 bits.
- The `complex128` type provides higher precision for more demanding calculations.

### 2. Boolean Data Type:
Boolearn is a data type, which only take `true` or `false`. It's declared using `bool` keyword.
> Note: Default value of `bool` is `false`.

```go
package main
import ("fmt")
func main(){
    var b1 bool = true // typed declaration with initial value
    var b2 = true // untyped declaration with initial value
    var b3 bool // typed declaration without initial value
    b4 := true // untyped declaration with initial value

    fmt.Println(b1)
    fmt.Println(b2)
    fmt.Println(b3)
    fmt.Println(b4)
}
```

### 3. String Data Type:
String data type is used to store a collection of characters. Strings must be surrounded by **double quotes**

```go
package main
import ("fmt")
func main() {
    var text1 string = "Hello World"
    var text2 := "Shorthand Hello World"
    var text3 string

    fmt.Println("Type: %T, Value: %v \n", text1)
    fmt.Println("Type: %T, Value: %v \n", text2)
    fmt.Println("Type: %T, Value: %v \n", text3)
}
```
> %T - Format Specifier used to print Type

> %v - Format Specifier used to print value

## Type Casting:
Type casting, or type conversion, is the **process of changing the data type of a variable to another basic data type**, allowing operations between different types, such as *converting an integer to a floating-point number or vice versa*.


**Example**:
```go
var x int = 10
var y float64 = float64(x)
```
Here, `float64(x)` converts the value of `x` from `int` to `float64`.

**Example**:
```go
var x int = 10
var y float64 = 5.5

sum := float64(x) + y // converting x to float64 before addition
```
In this case, `x` converted to `float64`  right before adding to `y`.