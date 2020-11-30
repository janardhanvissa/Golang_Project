package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	fmt.Println("Inside function")

	message, _ := bufio.NewReader(conn).ReadString('\n')

	fmt.Print("Message Received:", string(message))

	newmessage := strings.ToUpper(message)

	conn.Write([]byte(newmessage + "\n"))
	conn.Close()

}

func main() {
	fmt.Println("Launching server...")
	fmt.Println("Listen on port")
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Accept connection on port")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Calling handleConnection")
		go handleConnection(conn)
	}

}
