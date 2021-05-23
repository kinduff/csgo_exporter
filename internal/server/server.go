// Package server provides an HTTP server to serve the metrics and other endpoints.
package server

import (
	"net/http"
	"time"

	"github.com/kinduff/csgo_exporter/internal/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/context"

	log "github.com/sirupsen/logrus"
)

// Server is the struct for the HTTP server.
type Server struct {
	httpServer *http.Server
}

// NewServer method initializes a new HTTP server instance and associates
// the different routes that will be used by Prometheus (metrics) or for monitoring (health)
func NewServer(port string) *Server {
	mux := http.NewServeMux()
	httpServer := &http.Server{Addr: ":" + port, Handler: mux}

	s := &Server{
		httpServer: httpServer,
	}

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.Handle("/metrics", promhttp.Handler())

	return s
}

// ListenAndServe method serves HTTP requests.
func (s *Server) ListenAndServe() {
	log.Infof("Starting HTTP server on http://localhost%s", s.httpServer.Addr)

	err := s.httpServer.ListenAndServe()
	if err != nil {
		log.Printf("Failed to start serving HTTP requests: %v", err)
	}
}

// Stop method stops the HTTP server (so the exporter become unavailable).
func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s.httpServer.Shutdown(ctx)
}
