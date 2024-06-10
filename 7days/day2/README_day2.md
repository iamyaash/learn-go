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

