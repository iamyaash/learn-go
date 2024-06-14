# Day 6: Pointers and Methods: Understanding pointers, defining methods on types, interfaces

## Pointers

Pointers are variables that store the memory address of another variable in `Go`. Pointers can be used to directly access and modify the value stored at that memory address.

#### Important Factors
- **\* (Dereferencing Operator)** - Used to declare a pointer and to access the value stored at the memory address that the pointer holds.

- **\& (Address Operator)** - Used to get the memory address of a variable.
    > Used to **return the memory address** or **access the memory address** of a pointer to a variable.
- Default value of pointer is `nil`.

### Declaring a Pointer
**Syntax:**
```go
var pointer_name *data_type
```
**Example: Defined Type**
```go
package main 
import "fmt"
func main() {
    var a int = 8
    //init p pointer - p *int
    //return mem addr of a - &a
    var p *int = &a
    fmt.Println(p) //this'll print mem addr
}
```
**Example: Undefinded Type**
```go
package main 
import "fmt"
func main() {
    var a = 8
    //not initialized pointer p with *int
    //return mem addr of a - &a
    var p = &a
    fmt.Println(p) //this'll print mem addr
    fmt.Println(p*)//print the value at mem addr
    *p = 18        //val modifying
    fmt.Println(a)
}
```

**Example: Shorthand**
```go
package main 
import "fmt"
func main() {
    var a := 8
    //not initialized pointer p with *int
    //return mem addr of a - &a
    var p := &a
    fmt.Println(p) //this'll print mem addr
    fmt.Println(p*)//print the value at mem addr
    *p = 18        //val modifying
    fmt.Println(a)
}
```

## Defining Methods on Types
Let's understand **defining method on types** by looking at an example:

Let's create simple `Person` struct and a method to print the person's name.

```go
package main
import "fmt"

// **Defined Struct:** Struct defined with `Person`
type Person struct {
    name string
}

// **Defined Method:** Created a method with receiver of the `Person` type.
func (p Person) PrintName() {
    // Print the name of the person
    fmt.Println(p.name) // Corrected the typo from p.mame to p.name
}

func main() {
    // **Instance Creation:** Create an instance of the Person struct and call the PrintName method on it.
    person := Person{name: "Yaash"}
    person.PrintName() // This will print: Yaash
}

```

Here, in the above code:
- **Defined Struct:** Struct defined with `Person`
- **Defined Method:** Created a method of receiver of the `struct` type.
- **Instance Creation:** Create instance of the method and call the method on it.

## Interfaces
Interfaces in `Go` is different from all the other programming languages. 

An interface in `Go` is a type that specifies a set of **method signatures** (a *method's name, parameters, and return type*). Any type that provides definitions for all the methods in an interface is said to implement that interface.
#### Why Use Interfaces?
Interfaces allow you to write more flexible and reusable code. They enable you to define functions that can operate on different types that implement the same interface, without knowing the exact type.




**Syntax:**
```go
type interface_name{
    //method signatures
}
```

**Let's understand interfaces by looking into an example:**
1. **Define an Interface**: Create an interface with method signatures.
2. **Implement the Interface**: Define types that implement the interface.
3. **Use the Interface**: Write a function that accepts the interface type as an argument and call the method.

```go
package main 

import "fmt"

// Define an interface with a Speak() method
type Speaker interface {
    Speak()
}
// Define a struct with a field
type Person struct {
    name string
}
// Implement the Speak method for the Person struct
func (p Person) Speak() {
    fmt.Println(p.name) // Corrected the typo from fmt.Prntln to fmt.Println
}
func main() {
    // Create an instance of the Person struct
    person := Person{name: "Yaash"}

    // Call the Speak method on the Person instance
    person.Speak() // Output: Yaash
}
```
#### Explanation
1. **Define an Interface**: The `Speaker` interface is defined with a `Speak()` method.
2. **Implement the Interface**: The `Person` struct implements the `Speak()` method.
3. **Use the Interface**: An instance of `Person` is created and the `Speak()` method is called on it, demonstrating the implementation.
