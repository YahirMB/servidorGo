package handlers

import "servidor/router"

func AddRoutes(myRouter *router.Router) {
	myRouter.AddRoute("/hola", HelloHandler)
}
