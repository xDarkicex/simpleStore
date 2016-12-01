package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/SimpleStore/config"
	"github.com/xDarkicex/SimpleStore/helper"
	"github.com/xDarkicex/SimpleStore/server/router"
)

var (
	routes *httprouter.Router
)

// Server for serving site
func Server() {
	routes = router.GetRoutes()
	listen := fmt.Sprintf("%s:%s", config.Host, config.Port)
	fmt.Printf("Listening on %s\n", listen)
	helper.Logger.Fatal(http.ListenAndServe(listen, routes))
}
