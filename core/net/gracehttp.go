package net

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
)

//Serve serves servers using facebook gracehttp
func Serve(servers ...*http.Server) error {
	return gracehttp.Serve(servers...)
}
