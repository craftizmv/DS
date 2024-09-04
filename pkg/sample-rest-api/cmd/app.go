package cmd

import (
	"github.com/craftizmv/DS/pkg/sample-rest-api/pkg/auth"
	"github.com/craftizmv/DS/pkg/sample-rest-api/pkg/domain"
	"github.com/craftizmv/DS/pkg/sample-rest-api/pkg/login"
)

func InitApp() {
	// init router
	authHandler := &auth.AuthHandler{}

	loginHandler := &login.LoginHandler{}
	handlerMap := map[string]domain.RouteHandler{
		"auth":  authHandler,
		"login": loginHandler,
	}
	router := initRouter(handlerMap)

	// pass router to the http handler
}
