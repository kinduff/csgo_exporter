package config

import (
	"os"
	"reflect"
	"time"

	"github.com/kinduff/csgo_exporter/internal/model"

	log "github.com/sirupsen/logrus"
)

func Load() model.Config {
	config := model.Config{}

	config.HTTPPort = getEnv("HTTP_PORT", "7355")
	config.ApiKey = getEnv("STEAM_API_KEY", "")
	config.SteamID = getEnv("STEAM_ID", "")
	config.SteamName = getEnv("STEAM_NAME", "")
	interval, _ := time.ParseDuration(getEnv("SCRAPE_INTERVAL", "30s"))
	config.ScrapeInterval = interval

	if (config.SteamID == "" && config.SteamName == "") || config.ApiKey == "" {
		log.Fatal("Please provide a STEAM_API_KEY, and a STEAM_ID or STEAM_NAME")
		os.Exit(1)
	}

	show(config)

	return config
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func show(config model.Config) {
	log.Println("=============================================")
	log.Println("                CSGO Exporter                ")
	log.Println("=============================================")

	val := reflect.ValueOf(&config).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		log.Printf("%s: %s", typeField.Name, valueField.Interface())
	}

	log.Println("=============================================")
}
