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
    t := Employee{"Janardhan", 1234} 
    Teststruct(t) 
} 
