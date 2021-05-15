package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kinduff/csgo_exporter/config"
	"github.com/kinduff/csgo_exporter/internal/collector"
	"github.com/kinduff/csgo_exporter/internal/metrics"
	"github.com/kinduff/csgo_exporter/internal/server"

	log "github.com/sirupsen/logrus"
)

var (
	s *server.Server
)

func main() {
	conf := config.Load()

	metrics.Init()

	client := collector.NewCollector(conf)
	go client.Scrape()

	initHTTPServer(conf.HTTPPort)

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
