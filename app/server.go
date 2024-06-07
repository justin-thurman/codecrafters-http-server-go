package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting...")
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go func(c net.Conn) {
			defer c.Close()
			reader := bufio.NewReader(c)
			req, err := http.ReadRequest(reader)
			if err != nil {
				log.Fatal(err.Error())
			}
			if req.URL.Path == "/" {
				c.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
			} else {
				c.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
			}
		}(conn)
	}
}
