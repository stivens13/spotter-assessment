package usecase

import (
	"errors"

	"github.com/stivens13/spotter-assessment/app/models"
	"github.com/stivens13/spotter-assessment/app/repository"
)

const LATEST_VIDEO_LIMIT = 5

type VideoInteractor struct {
	videoRepo *repository.VideoRepository
}

func NewVideoInteractor(videoRepo *repository.VideoRepository) *VideoInteractor {
	return &VideoInteractor{
		videoRepo: videoRepo,
	}
}

func (vi *VideoInteractor) FetchLatestVideosByChannelID(channel_id string) (models.VideoList, error) {
	return vi.videoRepo.FetchLatestVideosByChannelID(channel_id, LATEST_VIDEO_LIMIT)
}

func (vi *VideoInteractor) Create(video *models.Video) error {
	if video.VideoID == "" || video.VideoTitle == "" || video.ChannelID == "" {
		return errors.New("title and URL cannot be empty")
	}

	err := vi.videoRepo.Create(video)
	if err != nil {
		return err
	}

	return nil
}