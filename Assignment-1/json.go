package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First_name string
	Last_name  string
	Age        int
}

func main() {
	p1 := person{
		First_name: "Janardhan",
		Last_name:  "Vissa",
		Age:        25,
	}

	p2 := person{
		First_name: "Krishna",
		Last_name:  "Vissa",
		Age:        18,
	}
	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error!", err)
	}
	fmt.Println(string(bs))

}
