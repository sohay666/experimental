package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"toktok/utils"
)

type CfgPaxel struct {
	Keyword string
	Type    string // search / popular
}

type VideoFiles struct {
	ID       int    `json:"id"`
	Quality  string `json:"quality"`
	FileType string `json:"file_type"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Link     string `json:"link"`
	Size     int    `json:"size"`
}
type Videos struct {
	ID         int          `json:"id"`
	Tags       []any        `json:"tags"`
	URL        string       `json:"url"`
	VideoFiles []VideoFiles `json:"video_files"`
}
type PaxelResponse struct {
	Status       int      `json:"status,omitempty"`
	Code         string   `json:"code,omitempty"`
	Page         int      `json:"page"`
	PerPage      int      `json:"per_page"`
	Videos       []Videos `json:"videos"`
	TotalResults int      `json:"total_results"`
	NextPage     string   `json:"next_page"`
	URL          string   `json:"url"`
}

func (cfg CfgPaxel) Search() (res PaxelResponse, err error) {
	sizeStr := fmt.Sprintf("%d", utils.Config.Integration.Pexels.Size)
	URL := utils.Config.Integration.Pexels.Host + "/videos/popular?per_page=" + sizeStr
	if len(cfg.Keyword) > 0 {
		URL = utils.Config.Integration.Pexels.Host + "/videos/search?query=" + cfg.Keyword + "&per_page=" + sizeStr
	}
	req, err := http.NewRequest("GET", URL, nil)

	req.Header.Set("Authorization", utils.Config.Integration.Pexels.SecretKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	if res.Status != 0 && res.Status != http.StatusOK {
		err = fmt.Errorf(res.Code)
		return
	}

	if len(res.Videos) > utils.Config.Integration.Pexels.Size {
		res.Videos = res.Videos[:utils.Config.Integration.Pexels.Size]
	}

	return
}

type PaxelSuggestionAttributes struct {
	Prefix      string   `json:"prefix"`
	Suggestions []string `json:"suggestions"`
}
type PaxelSuggestionData struct {
	Attributes PaxelSuggestionAttributes `json:"attributes"`
}
type PaxelSuggestionsResponse struct {
	Data PaxelSuggestionData `json:"data"`
}

func (cfg CfgPaxel) Suggestions() (res PaxelSuggestionsResponse, err error) {
	Host := strings.Replace(utils.Config.Integration.Pexels.Host, "api.", "", -1)
	URL := Host + "/en-us/api/v3/search/suggestions/" + cfg.Keyword

	req, err := http.NewRequest("GET", URL, nil)

	req.Header.Set("Authorization", utils.Config.Integration.Pexels.SecretKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	return
}
