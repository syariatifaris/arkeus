package net

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Router Contract
type Router interface {
	GetHandler() http.Handler
	PathPrefix(string) *mux.Route //as mux *Route
}

func NewRouter() Router {
	return &muxRouteImpl{
		router: mux.NewRouter(),
	}
}

type muxRouteImpl struct {
	router *mux.Router
}

func (m *muxRouteImpl) GetHandler() http.Handler {
	return m.router
}

func (m *muxRouteImpl) PathPrefix(path string) *mux.Route {
	return m.router.PathPrefix(path)
}
