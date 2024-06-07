package request

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Request struct {
	Method      string
	HTTPVersion string
	Target      string
	Headers     map[string]string
	Body        string
}

func New(r io.Reader) (*Request, error) {
	req := &Request{Headers: make(map[string]string)}
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	requestLine := scanner.Text()

	parts := strings.Split(requestLine, " ")

	if len(parts) != 3 {
		return nil, fmt.Errorf("expected request line to have 3 parts; got %d", len(parts))
	}

	req.Method = parts[0]
	req.Target = parts[1]
	req.HTTPVersion = parts[2]

	for {
		scanner.Scan()
		header := scanner.Text()
		if header == "" {
			break
		}
		parts = strings.Split(header, ": ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("expected header to split into 2 parts; got %v", len(parts))
		}
		req.Headers[parts[0]] = parts[1]
	}
	scanner.Scan()
	body := scanner.Text()
	req.Body = body

	return req, nil
}
