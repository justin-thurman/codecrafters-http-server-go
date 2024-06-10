package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/response"
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
			} else if strings.HasPrefix(req.URL.Path, "/echo/") {
				fmt.Printf("Got echo: %v", req.URL.Path)
				echoBody, _ := strings.CutPrefix(req.URL.Path, "/echo/")
				res := response.New(200, "OK", map[string]string{"Content-Type": "text/plain", "Content-Length": strconv.Itoa(len(echoBody))}, echoBody)
				c.Write([]byte(res.String()))
			} else {
				c.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
			}
		}(conn)
	}
}
