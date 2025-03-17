package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/samerzmd/commandarr/internal/config"
)

type RadarrMovie struct {
	Title     string `json:"title"`
	Year      int    `json:"year"`
	TmdbID    int    `json:"tmdbId"`
	TitleSlug string `json:"titleSlug"`
}

type AddMovieRequest struct {
	Title               string        `json:"title"`
	QualityProfileID    int           `json:"qualityProfileId"`
	TitleSlug           string        `json:"titleSlug"`
	TmdbID              int           `json:"tmdbId"`
	RootFolderPath      string        `json:"rootFolderPath"`
	Monitored           bool          `json:"monitored"`
	MinimumAvailability string        `json:"minimumAvailability"`
	Year                int           `json:"year"`
	Images              []interface{} `json:"images"`
	AddOptions          struct {
		SearchForMovie bool `json:"searchForMovie"`
	} `json:"addOptions"`
}

// SearchMovie searches Radarr for movies matching the query
func SearchMovie(cfg config.Config, query string) ([]RadarrMovie, error) {
	encodedQuery := url.QueryEscape(query)
	url := fmt.Sprintf("%s/api/v3/movie/lookup?term=%s&apikey=%s", cfg.RadarrURL, encodedQuery, cfg.RadarrAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var movies []RadarrMovie
	if err := json.NewDecoder(resp.Body).Decode(&movies); err != nil {
		return nil, err
	}

	return movies, nil
}

// AddMovie adds a movie to Radarr
func AddMovie(cfg config.Config, movie RadarrMovie, qualityProfileID int, rootFolderPath string) error {
	payload := AddMovieRequest{
		Title:               movie.Title,
		QualityProfileID:    qualityProfileID,
		TitleSlug:           movie.TitleSlug,
		TmdbID:              movie.TmdbID,
		RootFolderPath:      rootFolderPath,
		Monitored:           true,
		MinimumAvailability: "released",
		Year:                movie.Year,
		Images:              []interface{}{}, // critical: empty array, not nil
		AddOptions: struct {
			SearchForMovie bool `json:"searchForMovie"`
		}{
			SearchForMovie: true,
		},
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err)
	}

	reqURL := fmt.Sprintf("%s/api/v3/movie?apikey=%s", cfg.RadarrURL, cfg.RadarrAPIKey)
	resp, err := http.Post(reqURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("http post error: %v", err)
	}
	defer resp.Body.Close()

	// Improved error handling to capture detailed response
	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("radarr API error (%d): %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
