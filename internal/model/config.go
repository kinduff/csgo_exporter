package model

import (
	"time"
)

// Config stores the configuration coming from the dotenv file
// or from command-line arguments.
type Config struct {
	HTTPPort       string
	ApiKey         string
	SteamID        string
	SteamName      string
	ScrapeInterval time.Duration
}
