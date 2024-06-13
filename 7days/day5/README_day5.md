# Day 5: Data Structures: Arrays, slices, maps , structs

## Arrays:
- Arrays are used to store sequence of multiple values of same type in a single variable. Instead of declaring separate variables for each value.
- In `Go`, slices are most commonly used data structure rather than arrays. Since it's only used in special occasions.

#### Creating/Declaring an Array:
1. With `var` keyword, which is used to declare a variable.
```go
var arrName = [length]dataType{Values,......nValues} 
```
2. With `:=` shorthand sign, which automatically identifies the variable type.
```go
var arrName := [...]dataType{values}
```
##### Example:
```go
var arrName [8]int = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
arrName := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
arrName := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
```

#### Declaring an Array

```go
arr1 = [4]int{1,2,3,4}          // []int{1,2,...}
arr2 = [8]int{1,2,3,4,5,6,7,8}   // []int{1,...} 
arr1Str = [3]string{"here", "and", "there"}
arr2Str = []string{"Here", "And", "There"}
fmt.Println(arr1)
fmt.Println(arr2)
```
Here, 
- We basically create an `array` with type `int` assigned to a variable `arr1`.
- We can also do the same for `string` type array.
- At the end, the output is printed using `fmt` package.
- **Note:** We also declare nothing inside`[length]`, and providing `values{}` can be defined. Here the is `inferred`.

#### Accessing Element of an Array:

To access an array, we must learn the positions of array. Just like the lists in other programming languages, the **first element of an array is always 0**.

The **first** or the **Nth** element is accessed using,
- [0] -> for **first** element.
- [1] -> for **second** element.
- [n] -> for **Nth** element.

##### Example:
```go
array_name := [4]int{10, 20, 30, 40}
first := prices[0]
second := prices[1]
fmt.Println(first)
fmt.Println(second)
```
Here, we access the `Nth` element by `array_name[position-1]`.

#### Changing Element of an Array:
```go
array_name := [4]int{10, 20, 30, 40}
array_name[1] = 88 // [1] = {20}
fmt.Println(array_name)
//{10, 88, 30, 40}
```
**Note :**  We've changed the `element of an array` by assigning value to it by using `=(equal)`, not `:=` shorthand. Keep this in mind, and don't make this mistake.

#### Changing Specific Element of an Array:
```go
arr := [4]int{1:8,2:17}
// 0 8 17 0
```
**Note:** *Here, the value of an element get assigned specifically by defining them with position of the index. Hence, we have defined the `length[] as 4` and `values{}` are partially assigned to them with respect to the position of the index and the rest of the undefined elements get assigned with `value 0`*len()

#### How value are Assigned?

Before working on Arrays, we should know what is happening and how the values are declared inside. We'll take some cases and understand this while looking into it.

**The default value for `int = 0` & `string=""`**.
##### Case 1
```go
array1 := [4]int{}
// [0 0 0 0]
```
**Note:** *we have defined the `length[]`, but not the `values{}`.
Just because the array is filled with 0's doesn't mean it's empty, but the array is declared and by default it assigns 0's to every elements.*


##### Case 2
```go
array2 := [4]int{1, 2}
// [1 2 0 0]
```
**Note:** *This is a partially initialized array. Here, we did defined the `length[]` and partially defined the `values{1, 2}` as well. Hence, the first 2 elements are defined, the array assigns 1 & 2 to it's elements but assigns the default value to the rest of the elements*

##### Case 3
```go
array3 := [4]int{1, 2, 3, 4}
// 1 2 3 4
```
**Note:** *Here, both `length[]` and `values{}` are defined, and array itself assigns the value to it elements.*

> **Tip:** using `len()` will return the length of the data-type.


## Slices
Unline arrays, slices are most commanly used and important data type. Providing more **powerful interface to sequence than arrays**.

The length of a slice can grow and shrink as you se fit.

### Ways to create slice:
####  1. Creating slice from an `array`:
```go
var myarray := [length]datatype{values} // an array
myslice := myarray[start:end]           // slice made from array
```
##### Example:
```go
arr := [4]int{1, 2, 3, 4}
myslice := arr[1:2]
fmt.Println(myslice)
// [2 3]
```

#### 2.Using `make()` function:
```go
sliceName := make([]type, length, capacity)
```

#### Example:
```go
mySlice := make([]int, 4, 8)
```
Here, 
- `4` is the length of the slice
- `8` is the capacity of the slice
**Note:** If capacity parameter is not defined, it will equal to the length. i.e, `if capacity==null, then capacity == length`.


#### 3. Using the `[]datatype{values}` format.
```go
sliceName := []datatype{values}
```

##### Example:
```go
myslice := []int{}
fmt.Println(len(myslice))
fmt.Println(cap(myslice))
fmt.Println(myslice)
```

In `Go`, there are two functions that can be used to return the length and capacity of a slice:

- `len()` *function* - returns the length of the slice (the number of elements in the slice)
- `cap()` *function* - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

### Append Elements to Slice:

In `Go` we can use the function `append()` to add elements at the end of the slice.

```go
sliceName := append(sliceName, element1, element2, element3, ...)
```

#### Appending Elements only
```go
mySlice := [4]int{1, 2, 3, 4}
mySlice := append(mySlice, 5, 6, 7)
fmt.Println(mySlice)
// [1 2 3 4 5 6 7]
```
Here, 
1. `mySlice := [4]int{1,2,3,4}`
   - `mySlice` is a variable holding slice data type
   - `:=` shorthand symbol that explicitly assigns the data type based on the right-hand side value.
   - `[4]` is the lenght of the slice.
   - `int` is the datatype.
   - `{}` is the values of the slice.
2. `mySlice := append(mySlice, 5, 6, 7)`
   - `mySlice` is again the variable hoding slice data type.
   - `append()` is the function that appends the element at the end.
     - `(mySlice)` - is the slice which elements are appended
     - `(5, 6 ,7)` these are the values, that are going to be added at the end of the slice.

#### Appending One Slice to Another Slice:

```go
slice1 := [4]int{1,2,3,4}
slice2 := [4]int{5,6,7,8}
myslice := append(slice1, slice2)
// [1 2 3 4 5 6 7 8]
```


#### Changing Length of the Slice:
Unlike arrays, in `Go` it is possible to change the length of the slice.
```go
arr1 := [6]int{9, 10, 11, 12, 13, 14}
myslice1 := arr1[1:5] // Slice array
myslice1 = arr1[1:3] // Change length by re-slicing the array
myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
```

### Copying Elements to Slice:

The `copy()` is a function that takes two slices`dest` & `src` and copies data from `src` to `dest`. It returns the number if elements copied.
```go
copy(dest, src)
```

##### Example:

```go
package main

import "fmt"

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
    fmt.Println("Original slice:", nums)
    fmt.Println("Length:", len(nums))
    fmt.Println("Capacity:", cap(nums))

    // Create a copy with only needed numbers
    newnums := nums[:len(nums)-3] // Exclude last 3 numbers
    numscpy := make([]int, len(newnums))
    copy(numscpy, newnums)      // (dest, src)

    fmt.Println("Copied slice:", numscpy)
    fmt.Println("Length:", len(numscpy))
    fmt.Println("Capacity:", cap(numscpy))
}
```

**Here's what this code does:**

1. Initializes a slice `nums` with numbers from 1 to 8.
2. Prints original slice `nums`, its length, and capacity.
3. Creates a copy of `nums` with only the first 5 elements.
4. Prints the copied slice, its length, and capacity.

```
Original slice: [1 2 3 4 5 6 7 8]
Length: 8
Capacity: 8
Copied slice: [1 2 3 4 5]
Length: 5
Capacity: 5
```