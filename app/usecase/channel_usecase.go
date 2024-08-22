package usecase

import (
	"errors"
	"fmt"

	"github.com/stivens13/spotter-assessment/app/models"
	"github.com/stivens13/spotter-assessment/app/repository"
)

type ChannelInteractor struct {
	ChannelRepo *repository.ChannelRepository
}

func NewChannelInteractor(channelRepo *repository.ChannelRepository) *ChannelInteractor {
	return &ChannelInteractor{
		ChannelRepo: channelRepo,
	}
}

func (ci *ChannelInteractor) FetchChannel(channelID string) (*models.Channel, error) {
	if channelID == "" {
		return nil, errors.New("channel id cannot be empty")
	}

	return ci.ChannelRepo.GetChannelByID(channelID)
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
