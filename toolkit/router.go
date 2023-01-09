package toolkit

import (
	"net/http"
)

type Router struct {
	handler     http.Handler
	handlerFunc http.HandlerFunc
	path        string
	namedRoutes map[string]*Router
}

func NewRouter() *Router {
	return &Router{namedRoutes: make(map[string]*Router)}
}

// func (route *Router) Handler(w http.ResponseWriter, r http.Response) *Router {
// 	route.handlerFunc =
// }
