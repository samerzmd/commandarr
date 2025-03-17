package config

import "os"

type Config struct {
	TelegramToken string
	SonarrURL     string
	SonarrAPIKey  string
	RadarrURL     string
	RadarrAPIKey  string
}

func Load() Config {
	return Config{
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
		SonarrURL:     os.Getenv("SONARR_URL"),
		SonarrAPIKey:  os.Getenv("SONARR_API_KEY"),
		RadarrURL:     os.Getenv("RADARR_URL"),
		RadarrAPIKey:  os.Getenv("RADARR_API_KEY"),
	}
}
