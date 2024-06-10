package response

import (
	"fmt"
	"strings"
)

type Response struct {
	Headers map[string]string
	Reason  string
	Body    string
	Status  int
}

func New(status int, reason string, headers map[string]string, body string) *Response {
	return &Response{Headers: headers, Body: body, Status: status, Reason: reason}
}

func (r *Response) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "HTTP/1.1 %d %s\r\n", r.Status, r.Reason)
	for key, value := range r.Headers {
		fmt.Fprintf(&b, "%s: %s\r\n", key, value)
	}
	fmt.Fprintf(&b, "\r\n%s", r.Body)
	return b.String()
}
