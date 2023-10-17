package router

import (
	"bufio"
	"net"
	"servidor/response"
	"strings"
)

type Request struct {
	Method  string
	Route   string
	Headers map[string]string
	Hosts   string
	Body    []string
}

type Router struct {
	routes map[string]func(*response.Response, *Request)
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]func(*response.Response, *Request)),
	}
}

func (r *Router) AddRoute(route string, handler func(*response.Response, *Request)) {
	r.routes[route] = handler
}

func HandleRequest(conn net.Conn, router *Router) {
	scanner := bufio.NewScanner(conn)
	var header, body []string
	f := true
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			f = false
		}
		if f {
			header = append(header, ln)
		} else {
			body = append(body, ln)
		}
		if !f && ln == "" {
			break
		}
	}
	req := parseRequest(header)
	response := response.Response{Connect: conn}
	handler, found := router.routes[req.Route]
	if !found {
		handler = NotFoundHandler
	}
	handler(&response, &req)
}

func parseRequest(header []string) Request {
	req := Request{}
	req.Headers = make(map[string]string)
	for i, h := range header {
		if i == 0 {
			spl := strings.Split(h, " ")
			req.Method = spl[0]
			req.Route = spl[1]
		} else {
			spl := strings.Split(h, ": ")
			if spl[0] == "Host" {
				req.Hosts = spl[1]
			} else {
				req.Headers[spl[0]] = spl[1]
			}
		}
	}
	return req
}

func NotFoundHandler(response *response.Response, request *Request) {
	response.SendInfo(404, "Not Found", "text/plain", "Not Found page")
}
