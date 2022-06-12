Occurs when all processes are blocked while waiting for each other and program cannot proceed further

### Conditions for deadlock to occur
There are four conditions, known as the coffman conditions, that must be present simultaneously for a deadlock to occur
#### Mutual exclusion
A concurrent process holds atleast one resource at any one time making it non-sharable 
![Tux, the Linux mascot](/2.png)

#### Hold and wait
A concurrent process holds a resource and is waiting for an additional resource
![Tux, the Linux mascot](/3.png)

#### No Preemption
A resource held by a concurrent process cannot be taken away by the system. It can only be freed by the process holding it
![Tux, the Linux mascot](/4.png)

#### Circular wait
A concurrent process must be waiting on a chain of other concurrent processes such that P1 is waiting on P2, P2 on P3, and so on,
and there exists a Pn which is waiting on P1. This forms a ciruclar loop.
![Tux, the Linux mascot](/5.png)

In order to prevent deadlocks, we need to **make sure that atleast one of the condtions stated above should not hold**.

Possible deadlock code:

    package main

    import (
      "fmt"
    )
    func main() {
      ch := make(chan int)

      ch <- 10

      fmt.Println(<-ch)
    }

This creates deadlock because the problem is stuck on sending a value to the channel: **mychannel <- 10**. 
The sending operation is a blocking operation and requires the receive channel to be ready before sending data to the channel.

How to fix then?

    func main() {
      ch := make(chan int)

      go func() {
        ch <- 10
      }()

      fmt.Println(<-ch)
    }

Go only detects deadlocks when the program is stuck as a whole and not when a few of the goroutines are blocked. Also goroutines are mostly blocked while waiting for a channel operation or for the locks which belong to the **sync** package.

Following program runs successfully, although it is stuck. But Go only detects if the program stuct as a whole.

    func main() {
      ch1 := make(chan int)
      ch2 := make(chan int)
      ch3 := make(chan int)

      go func() {
        ch1 <- 10
      }()

      go func() {
        ch2 <- 20
      }()

      go func() {
        ch3 <- 30
      }()

      fmt.Println(<-ch2)
    }


