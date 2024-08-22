package youtubeclient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stivens13/spotter-assessment/models"
	"github.com/stivens13/spotter-assessment/config"
)

var FetchVideoMetadataYoututubeURL = "http://%s:%s/api/%s"

type YoutubeClient struct {
	cfg *config.YoutubeConfig
}

func NewYoutubeClient(cfg *config.YoutubeConfig) *YoutubeClient {
	return &YoutubeClient{cfg: cfg}
}

func (yc *YoutubeClient) getVideoMetadataQuery(channelID string) string {
	return fmt.Sprintf(FetchVideoMetadataYoututubeURL, 
		yc.cfg.Host,
		yc.cfg.Port,
		channelID,
	)
}

func (yc *YoutubeClient) FetchVideoMetadataFromYoutube(channelID string) (
	response models.VideoList,
	err error,
) {
	query := yc.getVideoMetadataQuery(channelID)
	resp, err := http.Get(query)
	if err != nil {
		fmt.Println("Error making HTTP GET request:", err)
		return response, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch youtube api channel data with status code: %d, channel: %s, query: %s\n", resp.StatusCode, channelID, query)
		return response, err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return response, err
	}

	return response, nil
}
