package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	gopher1 := make(chan string)
	gopher2 := make(chan string)

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Microsecond)
		gopher1 <- "Gopher 1 Buzzed"
	}()

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Microsecond)
		gopher2 <- "Gopher 2 Buzzed"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-gopher1:
			fmt.Println(m1)
		case m2 := <-gopher2:
			fmt.Println(m2)
		}
	}

}
