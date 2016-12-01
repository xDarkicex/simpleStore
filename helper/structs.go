package helper

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Controller Struct
type Controller struct {
	Controller interface{}
}

// RouterArgs These are the arguments passed in from a router.
type RouterArgs struct {
	Response http.ResponseWriter
	Request  *http.Request
	Params   httprouter.Params
}

// Flash ...
type Flash struct {
	Type    string
	Message string
}

// Users ...
type Users struct {
	name   string
	age    int
	weight int
}

// RoutesHandler for handling padding multiple objects into routes
type RoutesHandler func(a RouterArgs)
