//Queue program implementation

package main

import (
	"fmt"
)

//Queue represents a queue that holds a slice

type Queue struct {
	items []int
}

//Enqueue adds a value at the end

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)

}

//Dequeue removes a value at the front and returns the removed value

func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

//Main method

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(10)
	myQueue.Enqueue(20)
	myQueue.Enqueue(30)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

}
