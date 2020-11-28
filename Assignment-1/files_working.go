package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {

    mydata := []byte("Hello, i'm learning golang using kloudlearn application\n")

    // the WriteFile method returns an error if unsuccessful
    err := ioutil.WriteFile("myfile.txt", mydata, 0777)
    // handle this error
    if err != nil {
        // print it out
        fmt.Println(err)
    }

    data, err := ioutil.ReadFile("myfile.txt")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(string(data))

    f, err := os.OpenFile("myfile.txt", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    if _, err = f.WriteString("new data that wasn't there originally\n"); err != nil {
        panic(err)
    }

    data, err = ioutil.ReadFile("myfile.txt")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(string(data))

}