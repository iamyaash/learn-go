# Day 7: Concurrency Basics: Goroutines, channels, select statement

## Goroutines:

Goroutines are a special feature of golang. It is **function** or **method**, in which executes **independently** and **simultaneously** in connection with any other `Goroutines` present in your program.

In simple terms, every concurrently executing activity is known as **Goroutine**. It is both _light-weight and efficient_ when it comes to the cost of creation.

**Workflow** is simply defined, that every Goroutine will **work under a main Goroutine**. If the main Goroutine gets _terminates_, then all the other Goroutines _will be terminated as well_.

**Syntax:**
```go
func function_name() {
    //statements
}
go function_name()
```
> In order to execute the Goroutine, the method must be defined with `go` keyword beside the method name.

### Understanding Goroutines with Examples

#### Example 1:
```go
package main 
import "fmt"

func display(word string) {
    for i:=1; i>=4; i++ {
        fmt.Println(word)
    }
}

func main() {
    //Call Goroutine
    go display("go guy")
    //Call Function
    display("Hello")
}
```
##### Output:
```
Hello
Hello
Hello
Hello
```
#### Explanation for Example 1:
- You can see that, go program only executed normal function, but not the `Goroutine`.
- The reason is that when you start a goroutine with `go display("go guy")`, it runs concurrently with the main function. However, the main function does not wait for the goroutine to complete.

- The main function calls `display("Hello")`, which runs and completes its loop, printing "Hello" four times. By the time the `display` function completes, the main function also completes and exits, terminating the program along with any running goroutines.
- Therefore, the `display("go guy")` goroutine may not get a chance to execute before the program exits.

#### Example 2:
```go
package main 
import "fmt"
import "time"

func display(word string) {
    for i:=1; i<=4; i++ {
        time.Sleep(1 * time.Second)
        fmt.Println(word)
    }
}

func main() {
    //Call Goroutine
    go display("go guy")
    //Call Function
    display("Hello")
}
```

##### Output:
```
Hello
Hello
Hello
Hello
go guy
go guy
go guy
go guy
```

#### Explanation:
- In this example, you use `time.Sleep(1 * time.Second)` to introduce a delay in each iteration of the loop inside the display function.

- The main function starts the goroutine `go display("go guy")` and then immediately calls `display("Hello")`.
- Because `display("Hello")` takes around 4 seconds to complete (due to the sleep calls), the main function remains running, giving the `go display("go guy")` goroutine time to execute.

- The `display("Hello")` function prints "Hello" four times, and concurrently, the `display("go guy")`goroutine starts printing "go guy". The output can be interleaved due to the concurrent execution.

- The goroutine and the function **both get executed**, and their outputs **may mix depending on the scheduling** of the goroutines by the Go runtime.

## Anonymous Goroutines
Starting a goroutine with **anonymous function**. you can create **anonymous goroutine** using `go` keyword as a prefix of that function.

**Syntax:**
```go
go func (parameter) {
    //statements block
} (arguments)
```
Understanding anonymous goroutine with an example would be great,
#### Example:
```go
package main
import "fmt"
import "time"

func main() {
    fmt.Println("Main function")
    go func () {
        fmt.Println("Anonymous function")
    } ()
    time.Sleep(1 * time.Second)
    fmt.Println("Ta-ta")
}
```
##### Output:
```
Welcome!! to Main function
Welcome!! to GeeksforGeeks
GoodBye!! to Main function
```

##  Channels

A channel is a medium used to **pass data from one goroutine to another goroutine.** It's also bidirectional, meaning that the goroutine can send or receive data from the same channel.

**Syntax:**
1. Creating `Channel` using `chan` keyword.
```go
var Channel_name chan Type
```
> Only **transfer** the **data of the same type**.

2. Creating `Channel` using `make()` function with shorthand.
```go
channel_name := make(chan Type)
```
> Transfers **different types of data**.

#### Example:
```go
package main
import "fmt"
func main() {
    //channel creation
    var mychannel chan int
    newchannel := make(chan int)
    fmt.Println("Value of Channel:", mychannel)
    fmt.Println("Value of Channel:", newchannel)
    fmt.Println("Type of Channel: %T",mychannel)
    fmt.Println("Type of Channel: %T", newchannel)
}
```
##### Output:
```
Value of Channel: 
Value of Channel: 0x432112
Type of Channel: chan int
Type of Channel: chan int
```

## Channel Operations
### 1. Send Operation
The send operation is used to send data from one goroutine to another goroutine with the help of the channel.

- Sending values like `int`, `float` & `bool` are safe and easy to send, hence the values are copied and there is no risk of losing data. 
- String are easy to send as well, because they are immutable.
- However, sending pointers or references like map, slice, etc.. through a channel is **not safe** because the because the **values of pointers or reference may change by sending or receiving goroutine** and also the result is unpredicted.
```
mychannel <- element
```

### 2. Receive Operation
```
element := <-mychannel
```
Can also write a receive statement like:
```
<-mychannel
```

#### Example:
```go
package main
import "fmt"
func myfunc(ch chan int) {
    fmt.Println(234 + <-ch)
}
func main() {
    fmt.Println("Start")
    ch := make(chan int)
    go myfunc(ch)
    ch <- 23
    fmt.Println("End")
}
```
##### Output:
```
Start
256
End
```


#### Simple Example for Channel:
```go
package main
import "fmt"
import "time"

func main() {
    //create new channel of type string
    messages := make(chan string)
    //start new goroutine that sends message
    go func() {
        time.Sleep(1*time.Sleep) //delay
        messages <- "Hello from goroutine"
    }()
    //receive the message from channel
    msg := <-messages // wait for it to be received
    fmt.Println(msg)  // print the message
}
```
<br>

1. **Creating a new Channel:**
```go
messages := make(chan string)
```
Initiated new channel with type string.

2. **Sending data to Channel:**
```go
go func() {
    time.Sleep(1*time.Second)
    messages <- "Hello from goroutine"
}()
```
- A new goroutine gets started.
- With a time delay of 1 second.
- Sends data to `messages` channel with `string` type and `value` of `"Hello from goroutine"`.

3. **Receiving message from goroutine:**
```go
msg := <- messages
```
- The main function waits to receive data from the `messages` channel.
- `msg` receives the data from `messages` channel and stores it inside.


## Select Statement

The `select` statement allows a goroutine to wait on multiple communication operations(channels). It lets you perform different actions based on which channel operations becomes ready first, similar to a `switch` statement but for channels.

**Syntax:**
``` go
select {
    case casename1: //statement
    case casename2: //statement
    case casename3: //statement
    case casename4: //statement

    default:        //statement
}
```
>**Note:** The select statement always **waits until the sent or receive operation** is prepared for some cases to begin.

#### Example:
```go
package main
import "fmt"
import "time"

func main() {
    ch1 :=make(chan string)
    ch2 :=make(chan string)

    //starting goroutine for ch1
    go func() {
        time.Sleep(1*time.Second)
        ch1 <- "Message from ch1"
    }()

    //starting goroutine for ch2
    go func() {
        time.Sleep(1*time.Second)
        ch2 <- "Message from ch2"
    }()
    //using select to wait for message
    //from either ch1 or ch2 to get 
    //executed
    select {
        case msg1 := <- ch1:
        fmt.Println(msg1)
        case msg2 := <- ch2:
        fmt.Println(msg2)
        default:
        fmt.Println("Sorry bruh!")
    }
}
```

#### Explanation
1. **Creating Channels:**
```go
ch1 := make(chan string)
ch2 := make(chan string)
```
Creating two channels

2. **Sending Data to Channels:**
```go
go func() {
    time.Sleep(1*time.Second)
    ch1 <- "Message from ch1"
}()

go func() {
    time.Sleep(1*time.Second)
    ch2 <- "Message from ch2"
}()
```
Two of the goroutines are started. Each goroutine sends a message to its respective channel after a delay of 1 second

3. **Using select statement:**
```go
select {
    case msg1:= <-ch1:
    fmt.Println(msg1)
    case msg2:= <-ch2:
    fmt.Println(msg2)

    default:
    fmt.Println("Sorry bruh!")
}
```
- The `select` statement waits for one of the channel to become ready
- If one get communication, it will execute the statement block right away.
- Else, the default statement gets executed.