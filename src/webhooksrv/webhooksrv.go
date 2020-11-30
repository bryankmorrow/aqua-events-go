package webhooksrv

import (
	"github.com/BryanKMorrow/aqua-events-go/src/webhooksrv/router"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	// URL is the local address and port to listen
	URL = "0.0.0.0:8000"
)

// Server represents the webhooksrv
type Server struct {
	URL string
}

// NewServer instantiates a Server
func NewServer() Server {
	return Server{
		URL: URL,
	}
}

// Start the server
func (s *Server) Start() {
	log.Println("Starting server on :", URL)
	r := router.NewRouter()
	r.Init()
	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r.Router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)
	newServer := &http.Server{
		Handler:      handler,
		Addr:         URL,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	log.Fatal(newServer.ListenAndServe())
}
