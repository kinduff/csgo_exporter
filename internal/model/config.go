package model

// Config stores the configuration coming from the dotenv file
// or from command-line arguments.
type Config struct {
	HttpPort  string
	ApiKey    string
	SteamID   string
	SteamName string
}
