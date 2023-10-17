package main

import (
	"log"
	"net"
	"servidor/handlers"
	"servidor/router"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	myRouter := router.NewRouter()
	handlers.AddRoutes(myRouter)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go router.HandleRequest(conn, myRouter)
	}
}
