package dandelion

import (
	"fmt"
	"net/http"
)

type Router struct {
	routes []RouteEntry
}

type RouteEntry struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func (rou *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range rou.routes {
		match := route.Match(r)
		if !match {
			continue
		}

		route.Handler.ServeHTTP(w, r)
		return
	}
	http.NotFound(w, r)
}

func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	e := RouteEntry{
		Path:    path,
		Method:  method,
		Handler: handlerFunc,
	}

	rtr.routes = append(rtr.routes, e)
	fmt.Println(rtr)
}

func (re *RouteEntry) Match(r *http.Request) bool {
	if r.Method != re.Method {
		return false // method same
	}
	if r.URL.Path != re.Path {
		return false // path mismatch
	}

	return true
}
