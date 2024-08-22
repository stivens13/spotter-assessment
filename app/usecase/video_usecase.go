package usecase

import (
	"errors"

	"github.com/stivens13/spotter-assessment/app/helper/constants"
	"github.com/stivens13/spotter-assessment/app/models"
	"github.com/stivens13/spotter-assessment/app/repository"
)

type VideoInteractor struct {
	videoRepo *repository.VideoRepository
}

func NewVideoInteractor(videoRepo *repository.VideoRepository) *VideoInteractor {
	return &VideoInteractor{
		videoRepo: videoRepo,
	}
}

func (vi *VideoInteractor) FetchLatestVideosByChannelID(channel_id string) (models.VideoList, error) {
	response, err := vi.videoRepo.FetchLatestVideosByChannelID(channel_id, constants.LATEST_VIDEO_LIMIT)
	if err != nil {
		return models.VideoList{}, err
	}

	if len(response.Data) == 0 {
		return models.VideoList{}, errors.New("no videos found")
	}

	return response, nil
}

func (vi *VideoInteractor) Create(video *models.Video) (*models.Video, error) {
	if video.VideoID == "" || video.VideoTitle == "" || video.ChannelID == "" {
		return nil, errors.New("title and URL cannot be empty")
	}

	response, err := vi.videoRepo.Create(video)
	if err != nil {
		return nil, err
	}

	return response, nil
}
