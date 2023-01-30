package dandelion

import (
	"encoding/json"
	"net/http"
)

func NewRouter() *Router {
	return &Router{}
}

type Context interface {
	Response() *http.Response
	Request() *http.Request
}

type context struct {
	writer   http.ResponseWriter
	request  *http.Request
	response *http.Response
}

func (c *context) Request(r *http.Response) *context {
	c.response = r
	return c
}

func JSON(c *context, i interface{}) error {
	w := c.writer
	return json.NewEncoder(w).Encode(i)
}
