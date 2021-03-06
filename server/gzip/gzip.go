package gzip

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/SimpleStore/helper"
)

// Response is a Struct for manipulating io writer
type Response struct {
	io.Writer
	http.ResponseWriter
}

func (res Response) Write(b []byte) (int, error) {
	if "" == res.Header().Get("Content-Type") {
		// If no content type, apply sniffing algorithm to un-gzipped body.
		res.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return res.Writer.Write(b)
}

// Middleware force - bool, whether or not to force Gzip regardless of the sent headers.
func Middleware(fn httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, pm httprouter.Params) {
		a := helper.RouterArgs{Request: req, Response: res, Params: pm}
		res.Header().Set("Server", "Golang")
		if !strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			fn(a.Response, a.Request, a.Params)
			return
		}
		res.Header().Set("Vary", "Accept-Encoding")
		res.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(res)
		defer gz.Close()
		gzr := Response{Writer: gz, ResponseWriter: a.Response}
		a = helper.RouterArgs{Request: req, Response: gzr, Params: pm}
		fn(a.Response, a.Request, a.Params)
	}
}
