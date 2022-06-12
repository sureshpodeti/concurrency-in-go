Concurrent programs have to operate individually but with a shared data source

It leads to problems like race conditions

    package main

    import (
      "fmt"
    )

    func deposit(balance *int, amount int) {
      *balance += amount
    }

    func withdraw(balance *int, amount int) {
      *balance -= amount
    }
    func main() {
      balance := 100

      go deposit(&balance, 10)
      withdraw(&balance, 50)

      fmt.Println(balance)
    }

![Tux, the Linux mascot](/1.png)

# How to avoid data races?
using **channels** or **locks**
