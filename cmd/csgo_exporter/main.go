package main

import (
	"fmt"
	"net/http"
	"os"

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
		steamID  string
		apiKey   string
		httpHost string
		httpPort int
	)

	flag.StringVar(&config, "config", "", "path to config file")
	flag.StringVar(&httpHost, "host", "0.0.0.0", "HTTP host")
	flag.IntVar(&httpPort, "port", 7355, "HTTP port")
	flag.StringVar(&steamID, "steamid", "", "Your Steam ID")
	flag.StringVar(&apiKey, "apikey", "", "Your Steam API key")
	flag.Parse()

	if steamID == "" || apiKey == "" {
		flag.Usage()
		os.Exit(1)
	}

	registry := prometheus.NewRegistry()
	registry.MustRegister(playerCollector.NewPlayerCollector())

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.Handle("/metrics", handler)

	log.Infof("Listening on http://%s:%d", httpHost, httpPort)
	httpErr := http.ListenAndServe(
		fmt.Sprintf("%s:%d", httpHost, httpPort),
		logRequest(http.DefaultServeMux),
	)
	if httpErr != nil {
		log.Fatal(httpErr)
	}
}
