package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]interface{})
	myMap["First Name"] = "Janardhan"
	myMap["Age"] = 25
	fmt.Printf("%+v\n", myMap)
	// fetch value using type assertion
	fmt.Println(myMap["First Name"].(string))
	fetchValue(myMap)

	// Slices dynamically
	s := make([]int, 4, 6)
	fmt.Println(s, len(s), cap(s)) // [0 0 0 0] 4 6
	s = make([]int, 3)
	fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 3
}

func fetchValue(myMap map[string]interface{}) {
	for _, value := range myMap {
		switch v := value.(type) {
		case string:
			fmt.Println("the value is string =", value.(string))
		case float64:
			fmt.Println("the value is float64 =", value.(float64))
		case interface{}:
			fmt.Println(v)
		default:
			fmt.Println("unknown")
		}
	}
}
