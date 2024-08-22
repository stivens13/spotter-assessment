package usecase

import (
	"errors"

	"github.com/stivens13/spotter-assessment/app/helper/constants"
	"github.com/stivens13/spotter-assessment/app/models"
	"github.com/stivens13/spotter-assessment/app/repository"
)

type VideoInteractor struct {
	VideoRepo   *repository.VideoRepository
	YoutubeRepo *repository.YoutubeRepository
}

func NewVideoInteractor(
	videoRepo *repository.VideoRepository,
	youtubeRepo *repository.YoutubeRepository,
) *VideoInteractor {
	return &VideoInteractor{
		VideoRepo:   videoRepo,
		YoutubeRepo: youtubeRepo,
	}
}

func (vi *VideoInteractor) FetchLatestVideosByChannelID(channel_id string) (models.VideoList, error) {
	response, err := vi.VideoRepo.FetchLatestVideosByChannelID(channel_id, constants.LATEST_VIDEO_LIMIT)
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

	response, err := vi.VideoRepo.Create(video)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (vi *VideoInteractor) CreateBatch(videos models.VideoList) error {

	err := vi.VideoRepo.CreateBatch(videos)
	if err != nil {
		return err
	}

	return nil
}
