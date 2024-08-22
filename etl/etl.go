package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stivens13/spotter-assessment/app/config"
	"github.com/stivens13/spotter-assessment/app/models"
	"github.com/stivens13/spotter-assessment/tools/generator"
	youtubeclient "github.com/stivens13/spotter-assessment/youtube-client"
)

var (
	LoadChannelsURL = "http://localhost:8080/channels"
)

type ETL struct {
	YoutubeClient           *youtubeclient.YoutubeClient
	NewChannelsAmount       int
	AverageVideosPerChannel int
}

func NewETL(
	youtubeClient *youtubeclient.YoutubeClient,
	newChannelsAmount, averageVideosPerChannel int,
) *ETL {
	return &ETL{
		YoutubeClient:           youtubeClient,
		NewChannelsAmount:       newChannelsAmount,
		AverageVideosPerChannel: averageVideosPerChannel,
	}
}

func InitServices() (*ETL, error) {
	cfg := config.GetETLConfig()

	youtubeClient := youtubeclient.NewYoutubeClient(cfg.YoutubeConfig)

	etl := NewETL(youtubeClient, cfg.NewChannelAmount, cfg.AverageVideosPerChannel)

	return etl, nil
}

type channelsRawInput struct {
	Data []string `json:"data"`
}

func (etl *ETL) AggregateVideoMetadataFromYoutube(channelIDs []string) (videos models.VideoMap, err error) {
	videos.Data = make(map[string]models.VideoList)
	for _, channelID := range channelIDs {
		if videos.Data[channelID], err = etl.YoutubeClient.FetchVideoMetadataFromYoutube(channelID); err != nil {
			// in case of errors, log and keep going
			fmt.Printf("error fetching video metadata from youtube client: %v\n", err)
		}
	}
	return videos, nil
}

func (etl *ETL) StartETL() {
	fmt.Println("Starting ETL process...")
	etl.ExtractData()
	etl.TransformData()
	etl.LoadData()
}

func (etl *ETL) ExtractData() (videoMap models.VideoMap, err error) {
	fmt.Println("Generating New Channels...")
	channelIDs := etl.GetNewChannelIDs()
	if err := etl.LoadChannels(channelIDs); err != nil {
		return videoMap, fmt.Errorf("error loading channels: %v", err)
	}

	fmt.Println("Extract Video Metadata from the mock YouTube API...")
	videoMap, err = etl.AggregateVideoMetadataFromYoutube(channelIDs)
	if err != nil {
		return videoMap, fmt.Errorf("error aggregating video metadata: %v", err)
	}
	return videoMap, nil
}

func (etl *ETL) TransformData() {
}

func (etl *ETL) LoadData() {
}

func (etl *ETL) GetNewChannelIDs() []string {
	newChannelIDs := make([]string, etl.NewChannelsAmount)
	for range etl.NewChannelsAmount {
		newChannelIDs = append(newChannelIDs, generator.GenerateMockChannelID())
	}
	return newChannelIDs
}

func (etl *ETL) LoadChannels(channelIDs []string) error {
	channelsRaw := channelsRawInput{Data: channelIDs}
	jsonData, err := json.Marshal(channelsRaw)
	if err != nil {
		fmt.Println("Error converting channelIDs to JSON:", err)
		return err
	}

	resp, err := http.Post(LoadChannelsURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making HTTP POST request:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP POST request failed with status code:", resp.StatusCode)
		return err
	}
	return nil
}

func main() {
}

