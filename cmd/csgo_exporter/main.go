package main

import (
	"fmt"
	"net/http"

	"github.com/namsral/flag"

	playerCollector "github.com/kinduff/csgo_exporter/internal/collector"
	"github.com/kinduff/csgo_exporter/internal/handlers"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"
)

// Log internal request to stdout
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	var (
		config   string
		httpHost string
		httpPort int
	)

	flag.StringVar(&config, "config", "", "path to config file")
	flag.StringVar(&httpHost, "h", "0.0.0.0", "HTTP host")
	flag.IntVar(&httpPort, "p", 9009, "HTTP port")

	flag.Parse()

	registry := prometheus.NewRegistry()
	registry.MustRegister(playerCollector.NewPlayerCollector())

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.Handle("/metrics", handler)

	log.Infof("Listening on %s:%d", httpHost, httpPort)
	httpErr := http.ListenAndServe(
		fmt.Sprintf("%s:%d", httpHost, httpPort),
		logRequest(http.DefaultServeMux),
	)
	if httpErr != nil {
		log.Fatal(httpErr)
	}
}
