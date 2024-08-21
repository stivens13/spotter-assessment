package repository

import (
	"errors"

	"github.com/stivens13/spotter-assessment/app/models"
	"gorm.io/gorm"
)

type ChannelRepository struct {
	db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) *ChannelRepository {
	return &ChannelRepository{
		db: db,
	}
}

func (r *ChannelRepository) Create(channel *models.Channel) error {
	result := r.db.Create(channel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ChannelRepository) GetChannelByID(id int) (*models.Channel, error) {
	channel := &models.Channel{}
	result := r.db.First(channel, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("channel not found")
		}
		return nil, result.Error
	}
	return channel, nil
}

func (r *ChannelRepository) Update(channel *models.Channel) error {
	result := r.db.Save(channel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ChannelRepository) Delete(id int) error {
	result := r.db.Delete(&models.Channel{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

