This is my implementation for Codecrafter's HTTP Server course. Includes the gzip compression extension.

# Learnings
## The HTTP Spec
This was the biggest takeaway for me. I'm sure I've just barely scratched the surface, but before working on this server, I didn't really know what an HTTP request or response even was. I had only worked with higher level frameworks (e.g., Django), where you're handed a `request` object and don't really have to think about it. But just knowing the basic shape of an HTTP response (e.g., `HTTP/1.1 200 OK\r\nMy-Header: foo\r\n\r\nHello world!`) has given me a better understanding of how higher level web frameworks work -- and given me some intuition about the kinds of things that might go wrong.

## Buffers/IO
While not strictly related to HTTP servers, I also got more hands on experience working with buffers, IO, and string/byte conversion in Go. I'm still new to Go, with most of my professional experience in Python, where a lot of these lower level concerns are abstracted away. But I got much more experience with them here, and how to work with Go's ubiquitous `io.Reader` and `io.Writer` interfaces. The biggest lightbulb moment here was when working on the gzip compression feature. Noting that `gzip.NewWriter(w io.Writer) *Writer` can turn any writer into a gzip-compressed writer gave me a better appreciation of the power of Go's "accept interfaces, return structs" paradigm. It felt a lot like a Python decorator with less syntactical sugar, which is something I can get behind.
