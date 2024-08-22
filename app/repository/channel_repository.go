package repository

import (
	"errors"

	"github.com/stivens13/spotter-assessment/app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	ChannelBatchSize = 1000
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
	result := r.db.
		Model(&models.Channel{}).
		Create(channel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ChannelRepository) CreateBatch(channels models.ChannelList) error {
	if err := r.db.
		Model(&models.Channel{}).
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		Create(channels.Data).Error; err != nil {
		return err
	}
	return nil
}

func (r *ChannelRepository) GetChannelByID(channelID string) (*models.Channel, error) {
	response := &models.Channel{}
	if err := r.db.
		Model(&models.Channel{}).
		First(&response, "channel_id", channelID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("channel not found")
		}
		return nil, err
	}
	return response, nil
}

func (r *ChannelRepository) Update(channel *models.Channel) error {
	if err := r.db.
		Model(&models.Channel{}).
		Save(&channel).Error; err != nil {
		return err	
	}
	return nil
}

func (r *ChannelRepository) Delete(id int) error {
	if err := r.db.
		Delete(&models.Channel{}, id).
		Error; err != nil {
		return err
	}
	return nil
}
