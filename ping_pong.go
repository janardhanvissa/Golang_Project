package main

import (
	"fmt"
	"time"
)

func main() {
	pingc := make(chan int, 1)
	pongc := make(chan int, 1)

	go ping(pingc, pongc)
	go pong(pongc, pingc)

	pingc <- 1

	select {}
}

func ping(pingc <-chan int, pongc chan<- int) { //<-chan receiving a channel  // chan<- on a channel sending
	for {
		ball := <-pingc

		fmt.Println("Ping", ball)
		time.Sleep(1 * time.Second)

		pongc <- ball + 1
	}
}

func pong(pongc <-chan int, pingc chan<- int) {
	for {
		ball := <-pongc

		fmt.Println("Pong", ball)
		time.Sleep(1 * time.Second)

		pingc <- ball + 1
	}
}
