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

## Maps
Maps are unordered collection of `key:value` pairs, it provides fast lookups and values that can **retrieve**, **update** or **delete** with the help of keys.

Each element in a map is `key:value` pair and it doesn't allow duplicates. Addtionally, the **default value of map is `nil`**. 
- It allows any data type that supports `==`(equality operator) such as `bool`, `int`, `string`, `array`, `pointers`, `structs`, `interfaces`...etc,
- **Any value type is accepted**
- **Invalid key data types,**
  - Slices
  - Maps
  - Functions
**Note:** It also holds references to an underlying hash table.


#### Ways to Create Maps

1. **Create map using `var` and `:=`(shorthand)**
##### Syntax:
```go
var a = map[KeyType]ValueType{key1:value1, key2:value2, ...}
b := map[KeyType]ValueType{key1:value1, key2:value2, ...}
```
- `[KeyType]ValueType` - Such as `int`, `float`, `string`, `bool`...etc
- `{key1:value1, ..}`  - strong `values` with their respective `keys`.


##### Example:
```go
var a = map[string]string{"steve":"jobs", "bill":"gates"}
b := map[string]string{"steve":"jobs", "bill":"gates"}
                            // %v value
fmt.Printf("a\t%v\n", a)    // \t tab spaces
fmt.Printf("b\t%v\n", b)    // \n new line
```
```
Output:

a   map[steve:jobs bill:gates] 
b   map[steve:jobs bill:gates]
```
<br>

2. **Create using `make()` function**
##### Syntax:
```go
var a = make(map[KeyType]ValueType)
a := make(map[KeyType]ValueType)
```

##### Example:

```go
var a = make(map[string]string{"steve":"jobs", "bill":"gates"})
b := make(map[string]int{"steve":1, "bill":2})
fmt.Printf("a\t%v\n", a) 
fmt.Printf("b\t%v\n", b)
```

```
Output:

a   map[steve:jobs bill:gates]
b   map[steve:1 bill:2]
```

3. **Create an Empty Map:**
    1. Using `var` method:
    ```go
    var a map[string]string
    ```
    2. Using `make()` method:
    ```go
    var a = make(map[string]int)
    ```
##### Right way to create empty map
```go
b := make(map[string]int)
fmt.Println(b == nil)
// true
// trying other way to create map will result in false
```
> **Note:** `make()` is the right way to create an empty map. Creating empty map in other way and writing into it will cause runtime error.

#### Accessing Map Elements
```go
value = map_name[key]
```
##### Example:
```go
b := make(map[string]string)
a["steve"] = "jobs"
a["bill"] = "gates"

fmt.Println(a["bill"])
```
```
Output:

gates
```

#### Updating and Adding map Elements:

```go
map_name[key] = value
```
##### Example:
```go
b := make(map[string]string{"steve":"jobs", "bill":"gates"})
b["Sundar"] = "Pichai"  // Adding an Element
b["bill"] = "ball"
fmt.Println(b)
```

```
Output:
map[steve:jobs ball:gates Sundar:Pichai]
```

#### Removing Element from Map:

```go
delete(map_name, key)
```

```go
b := make(map[string]string{"steve":"jobs", "bill":"gates"})
b["Sundar"] = "Pichai"  // Adding an Element
b["bill"] = "ball"

delete(b, "ball")
fmt.Println(b)
```

```
Output:
map[steve:jobs Sundar:Pichai]
```

#### Specific Element in a Map

```go
val, ok :=map_name[key]
```
##### Example:
You can check if a key exists in a map by using the syntax `val, ok := map_name[key]`, where:

- `val` is the value associated with the key.
- `ok` is a boolean that indicates whether the key exists in the map.
- If you only want to check the existence of the key without caring about the value, you can use the blank identifier `(_)` for val.
```go
    car := map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

    // Check for existing key "brand" and get its value
    val1, ok1 := car["brand"]

    // Check for non-existing key "color" and get its value
    val2, ok2 := car["color"]

    // Check for existing key "day" and get its value (which is an empty string)
    val3, ok3 := car["day"]

    // Check for existing key "model" without getting its value
    _, ok4 := car["model"]

    // Print results
    fmt.Println(val1, ok1)  // Output: Ford true
    fmt.Println(val2, ok2)  // Output:  false
    fmt.Println(val3, ok3)  // Output:  true
    fmt.Println(ok4)        // Output: true
```

**Note:** When checking for specific element, whether it's available or not and it will obviously return `bool` types.


#### Map references to Hash Table

Maps are references type in `Go`. When you assign a variable a map to a variable, both variables refer to  the same underlying data structure `(hash table)`. Hence, changes made to one map affects the others too.

##### Example:
```go
pacakge main

import "fmt"

func main() {
    a := map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
    
    b := a                  //assigning a to map b, 
                            //hence both are pointing to the same map
    fmt.Println("Before!")
    fmt.Println(a)
    fmt.Println(b)
    
    b["year"] = "1970"      //changing value of "year" key in b
    fmt.Println("After!")
    fmt.Println(a)
    fmt.Println(b)
}
```

```
Output:

Before change:
map[brand:Ford model:Mustang year:1964]
map[brand:Ford model:Mustang year:1964]
After change to b:
map[brand:Ford model:Mustang year:1970]
map[brand:Ford model:Mustang year:1970]
```
