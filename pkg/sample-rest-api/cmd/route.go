package cmd

import (
	"github.com/craftizmv/DS/pkg/sample-rest-api/pkg/domain"
)

// login
// sign-up

type Router struct {
	handlerMap map[string]domain.RouteHandler
}

func initRouter(handlerMap map[string]domain.RouteHandler) Router {
	return Router{
		handlerMap: handlerMap,
	}
}
