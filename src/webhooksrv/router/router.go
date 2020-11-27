package router

import (
	"github.com/BryanKMorrow/aqua-events-go/src/webhooksrv/router/subroutes"
	"github.com/BryanKMorrow/aqua-events-go/src/webhooksrv/routes"
	"github.com/gorilla/mux"
	"log"
)


// Router represents the gorilla mux router
type Router struct {
	Router *mux.Router
}

// Init - Initialize the router and get the route and subroutes
func (r *Router) Init() {
	log.Println("Initializing Router")
	r.Router.Use(Middleware)
	baseRoutes := GetRoutes()
	for _, route := range baseRoutes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	v1SubRoutes := subroutes.GetRoutes()
	for name, pack := range v1SubRoutes {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}
}

// AttachSubRouterWithMiddleware - yes
func (r *Router) AttachSubRouterWithMiddleware(path string, subroutes routes.Routes, middleware mux.MiddlewareFunc) (SubRouter *mux.Router) {
	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(middleware)
	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return
}

// NewRouter - return the router
func NewRouter() (r Router) {
	r.Router = mux.NewRouter().StrictSlash(true).UseEncodedPath()
	return
}

