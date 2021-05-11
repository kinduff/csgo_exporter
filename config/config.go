package config

import (
	"fmt"
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
	log.Println("         CSGO Exporter Configuration         ")
	log.Println("=============================================")

	val := reflect.ValueOf(&config).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		value := fmt.Sprintf("%v", valueField.Interface())

		if typeField.Name == "ApiKey" {
			value = maskLeft(value)
		}

		log.Printf("%s: %s", typeField.Name, value)
	}

	log.Println("=============================================")
}

func maskLeft(s string) string {
	rs := []rune(s)
	for i := 6; i < len(rs); i++ {
		rs[i] = '*'
	}
	return string(rs)
}
