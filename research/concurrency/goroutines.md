# Goroutines
Goroutines are used to create concurrent programs. A concurrent programs are able to **run multiple processes at the same time**. Also goroutines can execute the function _independently and simultaneously_. Additionally goroutines are **ligthweight threads** in `Go`.
#### Real-time example
To understand this behaviour, we have to take a look on how the functions in `Go` are executed. The functions are executed asynchronously, that is when we have 2 functions in program, the second program gets executed after finishing the execution of first program. This is called synchronous execution.

Goroutines executes the functions asynchronously, that is both the function can be **executed at the same time**. *And this goroutines help us achieve concurrency in programming.*

### Things to know:
To understand this, we need to look into **Processes** & **Threads**.

**States in Process:**
|     State      | Description                                                   |
| :------------: | :------------------------------------------------------------ |
|    **new**     | The process is being created                                  |
|   **ready**    | The process is ready to run and waiting in CPU for allocation |
|  **running**   | The process is currently being executed by CPU                |
|  **waiting**   | The process is waiting for some event (I/O prompt).           |
| **terminated** | The process has finished execution.                           |
| **suspended**  | The process is paused, typically moved out of memory to disk. |

#### Definition of Process
A **process** is a program that is dispatched from **ready** state and are scheduled in CPU for **execution**. This process can create other processes, and they are known as **child processes**.

By default, the process takes more time to terminate and it is **isolated** from other process, meaning it **doesn't share memory** with other processes.

**States in Thread:**
|    State    | Description                                               |
| :---------: | :-------------------------------------------------------- |
| **running** | The thread is currently executing.                        |
|  **ready**  | The thread is ready to run and waiting in CPU allocation. |
| **blocked** | The thread is waiting for some event to occur             |

#### Definition of Thread
A **thread** is a segment of a process, meaning a process can have multiple threads and these threads are contained within the process itself.

A thread takes **less time to terminate itself** and it can't be isolated from other threads(they share the same memory under the same process).

### Creation of Goroutine:
We need a regular function to execute goroutine, calling the function with **`go` keyword** is enough to start the goroutine.

```go
func funcName() {
    //statements block
}
//start goroutine
go func()
```

#### Example: Concurrent programming using goroutine
Without
```go
package main
import (
    "fmt"
    "time"
)
func display(message string) {
    fmt.Println(message)
}
func main() {
    go display("Goroutine")
    display("Regular!")
}
```
##### Output:
```
Regular!
```
The reason is that, `main()` executes the goroutine which runs concurrently thus it moves the next function and it executes and after executing the `main()` will terminates itself even if other goroutines have not finished executed yet.

#### Example: Using Synchronization
```go
package main
import (
    "fmt"
    "time"
)

func display(message string){
    fmt.Println(message)
}

func main() {
    var wg sync.WaitGroup
    wg.add(1) //increment the waitgroup beforehand

    go func() {
        defer wg.Done() // decrement goroutine counter
        display("Goroutine")
    }()
    display("Regular!")
    wg.Wait()
}
```
##### Output:
```
Goroutine
Regular!
```

1. **Import** the `sync` package.
2. **Increment the WaitGroup counter**: Use `wg.Add(1)` to indicate a goroutine to wait for.
3. **Start the goroutine**: Use `go` to run a function concurrently, and use `defer wg.Done()` to signal its completion.
4. **Execute the regular function**: Call the function normally in `main()`.
5. **Wait for the goroutine to finish**: Use `wg.Wait()` to block until the goroutine completes.

#### Example: Running Multiple Goroutines
```go
package main
import "fmt"
import "sync"

func display(message string){
    defer wg.done()
    fmt.Println(message)
}

func main() {
    var wg sync.WaitGroup
    wg.add(3) //incremented as per 3 goroutines
    go display("hey1")
    go display("hey2")
    go display("hey3")

    wg.Wait()   //wait until it finishes goroutines
}
```
1. **Import** the `sync` package.
2. Initialize `display()` method with string type param.
3. In `main()`, create var with sync type and declare increment for 3 goroutines
4. When each goroutine executes, the count val get's decremented by -1.
5. After executing, it `wg.Wait()` for other goroutines to finish.


## Advantages of Goroutines
- It helps to execute two or more functions **independently & simultaneously**.
- Used to run programs in background.
- It communicate through **private channels**, indicates **safer communication**.
- We break one task into **multiple tasks to perform better**. 