package handlers

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/samerzmd/commandarr/internal/clients"
	"github.com/samerzmd/commandarr/internal/config"
)

func HandleMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, cfg config.Config) {
	text := strings.ToLower(msg.Text)
	args := strings.SplitN(text, " ", 2)

	switch {
	case strings.HasPrefix(text, "/start"):
		sendReply(bot, msg.Chat.ID, "Welcome to Commandarr 🚀")

	case strings.HasPrefix(text, "/search_movie"):
		if len(args) < 2 {
			sendReply(bot, msg.Chat.ID, "Usage: /search_movie <movie-name>")
			return
		}
		query := args[1]
		movies, err := clients.SearchMovie(cfg, query)
		if err != nil || len(movies) == 0 {
			sendReply(bot, msg.Chat.ID, "No movies found or an error occurred.")
			return
		}
		response := formatMovies(movies)
		sendReply(bot, msg.Chat.ID, response)

	case strings.HasPrefix(text, "/search_series"):
		if len(args) < 2 {
			sendReply(bot, msg.Chat.ID, "Usage: /search_series <series-name>")
			return
		}
		query := args[1]
		series, err := clients.SearchSeries(cfg, query)
		if err != nil || len(series) == 0 {
			sendReply(bot, msg.Chat.ID, "No series found or an error occurred.")
			return
		}
		reply := formatSeries(series)
		sendReply(bot, msg.Chat.ID, reply)

	case strings.HasPrefix(text, "/add_movie"):
		if len(args) < 2 {
			sendReply(bot, msg.Chat.ID, "Usage: /add_movie <movie-name>")
			return
		}
		query := args[1]
		movies, err := clients.SearchMovie(cfg, query)
		if err != nil || len(movies) == 0 {
			sendReply(bot, msg.Chat.ID, "Movie not found.")
			return
		}
		movie := movies[0]

		qualityProfileID := 1         // Confirm or update this from your Radarr settings
		rootFolder := "/media/movies" // ✅ Updated with correct path

		err = clients.AddMovie(cfg, movie, qualityProfileID, rootFolder)
		if err != nil {
			if strings.Contains(err.Error(), "MovieExistsValidator") {
				sendReply(bot, msg.Chat.ID, fmt.Sprintf("⚠️ Movie '%s' already exists.", movie.Title))
			} else {
				sendReply(bot, msg.Chat.ID, fmt.Sprintf("Failed to add movie: %v", err))
			}
			return
		}
		sendReply(bot, msg.Chat.ID, fmt.Sprintf("🎬 Movie '%s' added successfully!", movie.Title))

	case strings.HasPrefix(text, "/add_series"):
		if len(args) < 2 {
			sendReply(bot, msg.Chat.ID, "Usage: /add_series <series-name>")
			return
		}
		query := args[1]
		seriesList, err := clients.SearchSeries(cfg, query)
		if err != nil || len(seriesList) == 0 {
			sendReply(bot, msg.Chat.ID, "Series not found.")
			return
		}
		series := seriesList[0]

		qualityProfileID := 1     // confirm from Sonarr settings
		rootFolder := "/media/tv" // replace with your actual path

		err = clients.AddSeries(cfg, series, qualityProfileID, rootFolder)
		if err != nil {
			if strings.Contains(err.Error(), "SeriesExistsValidator") {
				sendReply(bot, msg.Chat.ID, fmt.Sprintf("⚠️ Series '%s' already exists.", series.Title))
			} else {
				sendReply(bot, msg.Chat.ID, fmt.Sprintf("Failed to add series: %v", err))
			}
			return
		}
		sendReply(bot, msg.Chat.ID, fmt.Sprintf("📺 Series '%s' added successfully!", series.Title))

	case strings.HasPrefix(text, "/my_id"):
		sendReply(bot, msg.Chat.ID, fmt.Sprintf("Your chat ID is: %d", msg.Chat.ID))
	default:
		sendReply(bot, msg.Chat.ID, "Unknown command. Please try again.")
	}
}

func sendReply(bot *tgbotapi.BotAPI, chatID int64, reply string) {
	msg := tgbotapi.NewMessage(chatID, reply)
	bot.Send(msg)
}

// Format Radarr movie results
func formatMovies(movies []clients.RadarrMovie) string {
	var result strings.Builder
	result.WriteString("🎬 Movies found:\n\n")
	for i, movie := range movies {
		result.WriteString(fmt.Sprintf("%d. %s (%d)\n", i+1, movie.Title, movie.Year))
		if i >= 4 { // limit to first 5 results
			break
		}
	}
	return result.String()
}

// Format Sonarr series results
func formatSeries(series []clients.SonarrSeries) string {
	var result strings.Builder
	result.WriteString("📺 Series found:\n\n")
	for i, s := range series {
		result.WriteString(fmt.Sprintf("%d. %s (%d)\n", i+1, s.Title, s.Year))
		if i >= 4 { // limit to first 5 results
			break
		}
	}
	return result.String()
}
