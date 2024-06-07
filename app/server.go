package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting...")
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	var request []byte
	_, err = conn.Read(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Request recieved")
	fmt.Printf("%s", request)

	// conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}
