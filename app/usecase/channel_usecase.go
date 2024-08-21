package usecase

import (
	"errors"

	"github.com/stivens13/spotter-assessment/app/models"
	"github.com/stivens13/spotter-assessment/app/repository"
)

type ChannelInteractor struct {
	ChannelRepo *repository.ChannelRepository
}

func NewChannelInteractor(ChannelRepo *repository.ChannelRepository) *ChannelInteractor {
	return &ChannelInteractor{
		ChannelRepo: ChannelRepo,
	}
}

func (ci *ChannelInteractor) Create(Channel *models.Channel) error {
	if Channel.ChannelID == "" {
		return errors.New("channel id cannot be empty")
	}

	err := ci.ChannelRepo.Create(Channel)
	if err != nil {
		return err
	}

	return nil
}