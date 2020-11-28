package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}
func worker(jobs <-chan int, results chan<- int) { //jobs <-chan only have a receive from jobs channel
	// results chan<- only have a send on results channel
	for n := range jobs {
		results <- fib(n)
	} // range function is to consume items from the queue
	// it is going to receive on the jobs channel
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
