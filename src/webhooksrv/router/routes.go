package router

import (
	"github.com/BryanKMorrow/aqua-events-go/src/webhooksrv/handlers"
	"github.com/BryanKMorrow/aqua-events-go/src/webhooksrv/routes"
	"net/http"
)

// Middleware - Main Middleware function
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// GetRoutes - Handle Authentication
func GetRoutes() routes.Routes {
	return routes.Routes{
		routes.Route{"Home", "GET", "/", handlers.IndexHandler},
	}
}
