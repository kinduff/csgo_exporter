// Package main takes care of the initialization of the application.
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kinduff/csgo_exporter/config"
	"github.com/kinduff/csgo_exporter/internal/client"
	"github.com/kinduff/csgo_exporter/internal/collector"
	"github.com/kinduff/csgo_exporter/internal/metrics"
	"github.com/kinduff/csgo_exporter/internal/server"

	log "github.com/sirupsen/logrus"
)

var (
	s *server.Server
)

func main() {
	cfg := config.Load()

	if cfg.SteamID == "" {
		client := client.NewClient()
		cfg.SteamID = client.RetrieveSteamID(cfg)
	}

	cfg.Show()

	metrics.Init(cfg)

	client := collector.NewCollector(cfg)
	go client.Scrape()

	initHTTPServer(cfg.HTTPPort)

	handleExitSignal()
}

func initHTTPServer(port string) {
	s = server.NewServer(port)
	go s.ListenAndServe()
}

func handleExitSignal() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	s.Stop()
	log.Fatal("HTTP server stopped")
}
