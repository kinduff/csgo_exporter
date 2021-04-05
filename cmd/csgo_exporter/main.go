package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/namsral/flag"

	playerCollector "github.com/kinduff/csgo_exporter/internal/collector"
	"github.com/kinduff/csgo_exporter/internal/handlers"
	"github.com/kinduff/csgo_exporter/internal/model"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"
)

// Log internal request to STDOUT.
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	config := model.Config{}

	flag.StringVar(&config.HttpHost, "host", "0.0.0.0", "HTTP host")
	flag.IntVar(&config.HttpPort, "port", 7355, "HTTP port")
	flag.StringVar(&config.SteamID, "steamid", "", "Your Steam ID")
	flag.StringVar(&config.SteamName, "steamname", "", "Your Steam name")
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, assume env variables are set.")
	}

	config.ApiKey = os.Getenv("STEAM_API_KEY")

	if (config.SteamID == "" && config.SteamName == "") || config.ApiKey == "" {
		flag.Usage()
		os.Exit(1)
	}

	registry := prometheus.NewRegistry()
	registry.MustRegister(playerCollector.NewPlayerCollector(&config))

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.Handle("/metrics", handler)

	log.Infof("Listening on http://%s:%d", config.HttpHost, config.HttpPort)

	httpErr := http.ListenAndServe(
		fmt.Sprintf("%s:%d", config.HttpHost, config.HttpPort),
		logRequest(http.DefaultServeMux),
	)
	if httpErr != nil {
		log.Fatal(httpErr)
	}
}
