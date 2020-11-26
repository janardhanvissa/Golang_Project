// Golang program to find a struct type 
// using type assertions 
package main 
  
import "fmt"
  
// struct Employee definition 
type Employee struct { 
    name        string 
    employee_id int
} 
  
func Teststruct(x interface{}) { 
    // type switch 
    switch x.(type) { 
    case Employee: 
        fmt.Println("Employee type") 
    case int: 
        fmt.Println("int type") 
    default: 
        fmt.Println("Error") 
    } 
} 
  
func main() { 
    // Declaring and initializing a 
    // struct using a struct literal 
    t := Employee{"Janardhan", 1234} 
    Teststruct(t) 
} 