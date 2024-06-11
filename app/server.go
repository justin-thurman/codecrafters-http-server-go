package main

import (
	"bufio"
	"errors"
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
			switch path := req.URL.Path; {
			case path == "/":
				res := response.New(200, "OK", "")
				c.Write([]byte(res.String()))
			case strings.HasPrefix(path, "/echo/"):
				echoBody, _ := strings.CutPrefix(req.URL.Path, "/echo/")
				res := response.New(200, "OK", echoBody)
				res.SetHeader("Content-Type", "text/plain")
				res.SetHeader("Content-Length", strconv.Itoa(len(echoBody)))
				c.Write([]byte(res.String()))
			case path == "/user-agent":
				userAgent := req.Header["User-Agent"][0]
				res := response.New(200, "OK", userAgent)
				res.SetHeader("Content-Type", "text/plain")
				res.SetHeader("Content-Length", strconv.Itoa(len(userAgent)))
				c.Write([]byte(res.String()))
			case strings.HasPrefix(path, "/files/"):
				var dir string
				dir = os.Args[2]
				if dir == "" {
					log.Fatal("No directory provided.")
				}
				fileName, _ := strings.CutPrefix(path, "/files/")
				filePath := dir + fileName
				data, err := os.ReadFile(filePath)
				if err != nil {
					if errors.Is(err, os.ErrNotExist) {
						res := response.New(404, "Not Found", "")
						c.Write([]byte(res.String()))
						return
					} else {
						log.Fatal(err.Error())
					}
				}
				dataStr := string(data)
				res := response.New(200, "OK", string(dataStr))
				res.SetHeader("Content-Type", "application/octet-stream")
				res.SetHeader("Content-Length", strconv.Itoa(len(dataStr)))
				c.Write([]byte(res.String()))
			default:
				res := response.New(404, "Not Found", "")
				c.Write([]byte(res.String()))
			}
		}(conn)
	}
}
