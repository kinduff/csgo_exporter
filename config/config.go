package config

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"

	log "github.com/sirupsen/logrus"
)

// Config is the exporter configuration.
type Config struct {
	HTTPPort       string        `config:"http_port,short=p"`
	SteamAPIKey    string        `config:"steam_api_key,required"`
	SteamID        string        `config:"steam_id"`
	SteamName      string        `config:"steam_name"`
	ScrapeInterval time.Duration `config:"scrape_interval,short=i,description=scrape interval in seconds"`
}

func getDefaultConfig() *Config {
	return &Config{
		HTTPPort:       "9617",
		SteamAPIKey:    "",
		SteamID:        "",
		SteamName:      "",
		ScrapeInterval: 30 * time.Second,
	}
}

// Load method loads the configuration by using environment variables.
func Load() *Config {
	loaders := []backend.Backend{
		env.NewBackend(),
		flags.NewBackend(),
	}

	loader := confita.NewLoader(loaders...)

	cfg := getDefaultConfig()
	err := loader.Load(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

// Show method displays all the load configuration
func (c Config) Show() {
	log.Println("=============================================")
	log.Println("         CSGO Exporter Configuration         ")
	log.Println("=============================================")

	val := reflect.ValueOf(&c).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		value := fmt.Sprintf("%v", valueField.Interface())

		if typeField.Name == "SteamAPIKey" {
			value = "[FILTERED]"
		}

		if value != "" {
			log.Printf("%s: %s", typeField.Name, value)
		}
	}

	log.Println("=============================================")
}
