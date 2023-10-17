package response

import (
	"fmt"
	"net"
)

type Response struct {
	Connect net.Conn
}

func (r Response) SendInfo(statusCode int, statusText, contentType, content string) {

	fmt.Fprintf(r.Connect, "HTTP/1.1 %d %s\r\n", statusCode, statusText)
	fmt.Fprintf(r.Connect, "Content-Type: %s\r\n", contentType)
	fmt.Fprint(r.Connect, "Access-Control-Allow-Origin: *\r\n")
	fmt.Fprint(r.Connect, "Connection: Keep-Alive\r\n")
	fmt.Fprint(r.Connect, "Keep-Alive: timeout=5, max=997\r\n")
	fmt.Fprint(r.Connect, "Server: Apache\r\n")
	fmt.Fprintf(r.Connect, "Content-Length: %d\r\n", len(content))
	fmt.Fprint(r.Connect, "\r\n")
	fmt.Fprint(r.Connect, content)
}
