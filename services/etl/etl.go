package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/stivens13/spotter-assessment/config"
	"github.com/stivens13/spotter-assessment/models"
	youtubeclient "github.com/stivens13/spotter-assessment/services/youtube-client"
	"github.com/stivens13/spotter-assessment/tools/generator"
)

type ETL struct {
	YoutubeClient           *youtubeclient.YoutubeClient
	SpotterCFG 			*config.SpotterAPIConfig
	NewChannelsAmount       int
	AverageVideosPerChannel int
}

func NewETL(
	youtubeClient *youtubeclient.YoutubeClient,
	spotterCfg *config.SpotterAPIConfig,
	newChannelsAmount, averageVideosPerChannel int,
) *ETL {
	return &ETL{
		YoutubeClient:           youtubeClient,
		SpotterCFG: 			spotterCfg,
		NewChannelsAmount:       newChannelsAmount,
		AverageVideosPerChannel: averageVideosPerChannel,
	}
}

func InitServices() (*ETL, error) {
	cfg := config.GetETLConfig()

	youtubeClient := youtubeclient.NewYoutubeClient(cfg.YoutubeConfig)

	spotterCfg := config.GetSpotterAPIConfig()

	etl := NewETL(youtubeClient, spotterCfg, cfg.NewChannelAmount, cfg.AverageVideosPerChannel)

	return etl, nil
}

type channelsRawInput struct {
	Data []string `json:"data"`
}

func (etl *ETL) GetSpotterAPIChannelsQuery() (string) {
	return fmt.Sprintf("http://%s:%s/channels", etl.SpotterCFG.Host, etl.SpotterCFG.Port)
}

func (etl *ETL) GetSpotterAPIVideosQuery() (string) {
	return fmt.Sprintf("http://%s:%s/videos", etl.SpotterCFG.Host, etl.SpotterCFG.Port)
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
	data, err := etl.ExtractData()
	if err != nil {
		log.Fatalf("error extracting data: %v\n", err)
		return
	}
	processedData := etl.TransformData(data)
	etl.LoadVideos(processedData)
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

func (etl *ETL) TransformData(data models.VideoMap) (result models.VideoList) {
	for _, videos := range data.Data {
		result.Data = append(result.Data, videos.Data...)
	}
	return result
}

func (etl *ETL) LoadVideos(data models.VideoList) error {
    jsonData, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error converting video list to JSON:", err)
        return err
    }

    videosQuery := etl.GetSpotterAPIVideosQuery()

    client := resty.New()

    resp, err := client.R().
        SetHeader("Content-Type", "application/json").
        SetBody(jsonData).
        Post(videosQuery)

    if err != nil {
        fmt.Println("Error making HTTP POST request:", err)
        return err
    }

    if resp.StatusCode() != http.StatusCreated {
        fmt.Println("spotter load channels request failed with status code:", resp.StatusCode())
        return fmt.Errorf("request failed with status code: %d", resp.StatusCode())
    }
	fmt.Printf("Loaded %d videos\n", len(data.Data))

    return nil
}

func (etl *ETL) GetNewChannelIDs() []string {
	newChannelIDs := make([]string, 0)
	for range etl.NewChannelsAmount {
		newChannelID := generator.GenerateMockChannelID()
		newChannelIDs = append(newChannelIDs, newChannelID)
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

    channelsQuery := etl.GetSpotterAPIChannelsQuery()

    client := resty.New()

    resp, err := client.R().
        SetHeader("Content-Type", "application/json").
        SetBody(jsonData).
        Post(channelsQuery)

    if err != nil {
        fmt.Println("Error making HTTP POST request:", err)
        return err
    }

    if resp.StatusCode() != http.StatusCreated {
        fmt.Println("spotter load channels request failed with status code:", resp.StatusCode())
        return fmt.Errorf("request failed with status code: %d", resp.StatusCode())
    }
	fmt.Printf("Loaded %d channels\n", len(channelIDs))

    return nil
}

func main() {
	time.Sleep(10 * time.Second) // wait for the spotter-api to start (etl can't connect
								// even when spotter-api is healthy)
	etl, err := InitServices()
	if err != nil {
		log.Fatalf("error initializing ETL services: %v\n", err)
	}

	etl.StartETL()
	fmt.Printf("ETL process completed successfully\n")
}
