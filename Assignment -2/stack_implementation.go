//Stack Program implementation

package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

//push will add a value at the end

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//Pop will remove a value at the end and returns the removed value

func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

//Main method

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)

}
