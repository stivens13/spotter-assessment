package usecase

import (
	"errors"
	"fmt"

	"github.com/stivens13/spotter-assessment/models"
	"github.com/stivens13/spotter-assessment/services/spotter-api/repository"
	"gorm.io/gorm"
)

type ChannelInteractor struct {
	ChannelRepo     *repository.ChannelRepository
	VideoInteractor *VideoInteractor
	YoutubeRepo     *repository.YoutubeRepository
}

func NewChannelInteractor(
	channelRepo *repository.ChannelRepository,
	videoInteractor *VideoInteractor,
	youtubeRepo *repository.YoutubeRepository,
) *ChannelInteractor {
	return &ChannelInteractor{
		ChannelRepo:     channelRepo,
		VideoInteractor: videoInteractor,
		YoutubeRepo:     youtubeRepo,
	}
}

func (ci *ChannelInteractor) FetchChannel(channelID string) (response *models.Channel, err error) {
	if channelID == "" {
		return nil, errors.New("channel id cannot be empty")
	}

	if response, err = ci.ChannelRepo.GetChannelByID(channelID); err != nil {
		// check if channel is missing in registry
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// attempt to populate missing channel with videos
			if response, err = ci.PopulateMissingChannelWithVideos(channelID); err != nil {
				return nil, fmt.Errorf("failed to populate missing channel: %w", err)
			}
			return response, nil
		} else {
			return nil, fmt.Errorf("error fetching channel: %w", err)
		}
	}

	return response, nil
}

func (ci *ChannelInteractor) PopulateMissingChannelWithVideos(channelID string) (
	response *models.Channel,
	err error,
) {
	videos, err := ci.YoutubeRepo.FetchVideoMetadataFromYoutube(channelID)
	if err != nil {
		return nil, fmt.Errorf("error on missing channel fallback: channel does not exist: %w", err)
	}
	if err := ci.Create(&models.Channel{ChannelID: channelID}); err != nil {
		return nil, fmt.Errorf("failed to populate missing channel: %w", err)
	}
	if err := ci.VideoInteractor.CreateBatch(videos); err != nil {
		return nil, fmt.Errorf("failed to populate missing channel videos: %w", err)
	}

	return response, nil
}

func (ci *ChannelInteractor) Create(channel *models.Channel) error {
	if channel.ChannelID == "" {
		return errors.New("channel id cannot be empty")
	}

	err := ci.ChannelRepo.Create(channel)
	if err != nil {
		return err
	}

	return nil
}

func (ci *ChannelInteractor) CreateBatch(channelsRaw models.ChannelRawList) error {
	var channels models.ChannelList
	channels.Data = make([]*models.Channel, len(channelsRaw.Data))

	for i, channelID := range channelsRaw.Data {
		if channelID == "" {
			return fmt.Errorf("channel id cannot be empty: %s", channelID)
		}
		channels.Data[i] = &models.Channel{ChannelID: channelID}
	}

	err := ci.ChannelRepo.CreateBatch(channels)
	if err != nil {
		return err
	}

	return nil
}
