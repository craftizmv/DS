package domain

import "github.com/craftizmv/DS/pkg/sample-rest-api/cmd"

type RouteHandler interface {
	Handle(req *cmd.Request) (bool, error)
}
