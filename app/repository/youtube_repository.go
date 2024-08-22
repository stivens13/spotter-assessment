package repository

import (
	"fmt"

	"github.com/stivens13/spotter-assessment/app/models"
	youtubeclient "github.com/stivens13/spotter-assessment/youtube-client"
)

var FetchVideoMetadataYoututubeURL = "http://localhost:9000/videos"

type YoutubeRepository struct {
	YoutubeClient *youtubeclient.YoutubeClient
}

func NewYoutubeRepository(client *youtubeclient.YoutubeClient) *YoutubeRepository {
	return &YoutubeRepository{YoutubeClient: client}
}

func (yr *YoutubeRepository) FetchVideoMetadataFromYoutube(channelID string) (response models.VideoList, err error) {

	response, err = yr.YoutubeClient.FetchVideoMetadataFromYoutube(channelID)
	if err != nil {
		return response, fmt.Errorf("error fetching video metadata from youtube client: %w", err)
	}

	return response, nil
}
