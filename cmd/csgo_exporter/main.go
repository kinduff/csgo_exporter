package main

import (
	"fmt"
	"net/http"
	"os"

	playerCollector "github.com/kinduff/csgo_exporter/internal/collector"
	"github.com/kinduff/csgo_exporter/internal/handlers"
	"github.com/kinduff/csgo_exporter/internal/model"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Log internal request to STDOUT.
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	config := model.Config{}

	config.HttpPort = getEnv("HTTP_PORT", "7355")
	config.ApiKey = getEnv("STEAM_API_KEY", "")
	config.SteamID = getEnv("STEAM_ID", "")
	config.SteamName = getEnv("STEAM_NAME", "")

	if (config.SteamID == "" && config.SteamName == "") || config.ApiKey == "" {
		log.Fatal("Please provide an STEAM_API_KEY, and a STEAM_ID or STEAM_NAME")
		os.Exit(1)
	}

	registry := prometheus.NewRegistry()
	registry.MustRegister(playerCollector.NewPlayerCollector(&config))

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.Handle("/metrics", handler)

	log.Infof("Listening on http://localhost:%s", config.HttpPort)

	httpErr := http.ListenAndServe(
		fmt.Sprintf(":%s", config.HttpPort),
		logRequest(http.DefaultServeMux),
	)
	if httpErr != nil {
		log.Fatal(httpErr)
	}
}
