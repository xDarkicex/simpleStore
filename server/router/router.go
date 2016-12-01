package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/SimpleStore/app/controllers"
	"github.com/xDarkicex/SimpleStore/helper"
	"github.com/xDarkicex/SimpleStore/server/gzip"
	"github.com/xDarkicex/SimpleStore/server/handle"
)

// type routesHandler func(a RouterArgs, user interface{})
func route(controller helper.RoutesHandler) httprouter.Handle {
	return gzip.Middleware(handle.Handle(controller))
}

// GetRoutes func to setup all routes
func GetRoutes() *httprouter.Router {
	router := httprouter.New()
	///////////////////////////////////////////////////////////
	// Main application routes
	///////////////////////////////////////////////////////////

	application := controllers.Application{}
	router.GET("/", route(application.Index))
	router.GET("/api/products", route(application.AllProducts))
	router.GET("/api/products/match", route(application.Match))

	///////////////////////////////////////////////////////////
	// Static routes
	// Caching Static files
	///////////////////////////////////////////////////////////
	fileServer := http.FileServer(http.Dir("public"))
	router.GET("/static/*filepath", gzip.Middleware(func(res http.ResponseWriter, req *http.Request, pm httprouter.Params) {
		res.Header().Set("Vary", "Accept-Encoding")
		res.Header().Set("Cache-Control", "public, max-age=7776000")
		req.URL.Path = pm.ByName("filepath")
		fileServer.ServeHTTP(res, req)
	}))
	return router
}
