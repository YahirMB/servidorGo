package handlers

import (
	"fmt"
	"servidor/response"
	"servidor/router"
)

func HelloHandler(response *response.Response, request *router.Request) {
	// Implementa la lógica para la ruta /hola

	fmt.Printf("%v\t %v\n", request.Method, request.Route)
	response.SendInfo(200, "OK", "text/html; charset=utf-8", generateHTMLResponse())
}

func generateHTMLResponse() string {
	return `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>Bienvenido</h1>
</body>
</html>
	`
}
