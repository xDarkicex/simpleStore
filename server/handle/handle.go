package handle

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/SimpleStore/helper"
)

func Handle(fn helper.RoutesHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		a := helper.RouterArgs{Request: r, Response: w, Params: ps}
		fn(a)
	}
}
