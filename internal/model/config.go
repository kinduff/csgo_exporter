package model

// Config stores the configuration coming from the dotenv file
// or from command-line arguments.
type Config struct {
	ApiKey    string
	SteamName string
	SteamID   string
	HttpHost  string
	HttpPort  int
}
