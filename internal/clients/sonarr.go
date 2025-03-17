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

type SonarrSeries struct {
	Title     string `json:"title"`
	TvdbID    int    `json:"tvdbId"`
	Year      int    `json:"year"`
	TitleSlug string `json:"titleSlug"` // clearly include this
}

func SearchSeries(cfg config.Config, query string) ([]SonarrSeries, error) {
	encodedQuery := url.QueryEscape(query)
	reqURL := fmt.Sprintf("%s/api/v3/series/lookup?term=%s&apikey=%s", cfg.SonarrURL, encodedQuery, cfg.SonarrAPIKey)

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var series []SonarrSeries
	err = json.NewDecoder(resp.Body).Decode(&series)
	if err != nil {
		return nil, err
	}

	return series, nil
}

type AddSeriesRequest struct {
	Title            string `json:"title"`
	QualityProfileID int    `json:"qualityProfileId"`
	TitleSlug        string `json:"titleSlug"`
	TvdbID           int    `json:"tvdbId"`
	RootFolderPath   string `json:"rootFolderPath"`
	Monitored        bool   `json:"monitored"`
	SeasonFolder     bool   `json:"seasonFolder"`
	AddOptions       struct {
		SearchForMissingEpisodes bool `json:"searchForMissingEpisodes"`
	} `json:"addOptions"`
}

func AddSeries(cfg config.Config, series SonarrSeries, qualityProfileID int, rootFolderPath string) error {
	payload := AddSeriesRequest{
		Title:            series.Title,
		QualityProfileID: qualityProfileID,
		TitleSlug:        series.TitleSlug,
		TvdbID:           series.TvdbID,
		RootFolderPath:   rootFolderPath,
		Monitored:        true,
		SeasonFolder:     true,
		AddOptions: struct {
			SearchForMissingEpisodes bool `json:"searchForMissingEpisodes"`
		}{
			SearchForMissingEpisodes: true, // Corrected Field
		},
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err)
	}

	reqURL := fmt.Sprintf("%s/api/v3/series?apikey=%s", cfg.SonarrURL, cfg.SonarrAPIKey)
	resp, err := http.Post(reqURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("sonarr API error (%d): %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
