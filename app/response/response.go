package response

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	Headers map[string]string
	Reason  string
	Body    string
	Status  int
}

func New(status int, reason string) *Response {
	return &Response{Status: status, Reason: reason, Headers: make(map[string]string)}
}

func (r *Response) SetHeader(key, value string) {
	r.Headers[key] = value
}

func (r *Response) SetBody(body string, req *http.Request) {
	r.Body = body
	encoding, ok := req.Header["Accept-Encoding"]
	if ok {
		for _, encodings := range encoding {
			for _, val := range strings.Split(encodings, ",") {
				val = strings.TrimSpace(val)
				if val == "gzip" {
					r.SetHeader("Content-Encoding", "gzip")
					var buff bytes.Buffer
					writer := gzip.NewWriter(&buff)
					writer.Write([]byte(body))
					writer.Close()
					r.Body = buff.String()
					break
				}
			}
		}
	}
	r.SetHeader("Content-Length", strconv.Itoa(len(r.Body)))
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
