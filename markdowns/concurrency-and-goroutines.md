Concurrency and Goroutines
===
Go offers build-in native support for concurrency through the use of **Goroutines** and **Channels**. These synchronization primitives give go the flexibility to utilize the modern multiprocessor architectures with ease and relative safety as opposed to using plain threads and memory barriers which are unsafe and relatively hard to understand if not used correctly.

Goroutines
====
A goroutine is a lightweight thread managed by the Go runtime. It is a function that is capable of running concurrently with other functions. To create a goroutine we use the keyword `go` followed by a function invocation:

```go
package main

import "fmt"

func run(n int) {
  fmt.Println(n, ":", n*n)
}

func main() {
  go run(10)
  var input string
  fmt.Scanln(&input) // note the use of Scanln
}
```

Note the usage of `fmt.Scanln`. This is to ensure that the main function needs to w8 for input thus giving a chance to the goroutine to run.

Normally when we invoke a function our program will execute all the statements in a function and then return to the next line following the invocation. 

However with a goroutine we return immediately to the next line and don't wait for the function to complete.

Channels
===
Channels are a language feature that enable the goroutines to communicate with one another and synchronize their execution.

```go

func printer(c chan string) {
    msg := <- c // this will block until it receives a message. 
    fmt.Println(msg)
}

func main() {
  var c chan string = make(chan string)
  go printer(c) // important to run as a goroutine here

  c <- "ping"
}
```

This program will print **ping** and exit. The way it works is like that. The printer function accepts a unbuffered channel of strings.

This line `msg := <- c` means that the channel will block until it receives a message in the channel. When it eventually do it will return it to the `msg` variable and then print it.

Now in main we run the printer function in a goroutine and pass a message string to `c`. If we didn;t do that we would get a deadlock as the printer function would run on the main thread and block on  `msg := <- c` and would not have a chance to receive the **ping**

Channel Direction
===
We can specify a direction on a channel type thus restricting it to either sending or receiving:

```go
func accept(c chan<- string) // accepts strings in channel

func accept(c <-chan string) // returns or resolves strings from the channel
```

Quiz Time
===
?[How would you create a buffered channel one with a capacity of 20?]
-[ ] make(chan, 20)
-[x] make(chan int, 20)
-[ ] chan &buf{cap: 20}
-[ ] new(chan int, 20)

?[How would you create a function named *accept* that accepts a channel parameter for only sending string values.?]
-[ ] func accept(input <-chan<- string)
-[ ] func accept(input <-chan string)
-[ ] func accept(input chan string)
-[x] func accept(input chan<- string)
