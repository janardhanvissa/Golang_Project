package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}
func main() {
	_, err := os.Open("cmd.txt")
	if err != nil {
		//fmt.Println("error occured", err)
		log.Println("error occured", err)
		//log.Fatalln(err)
		//panic(err)
	}
}
