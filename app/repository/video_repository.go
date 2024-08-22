package repository

import (
	"github.com/stivens13/spotter-assessment/app/models"
	"gorm.io/gorm"
)

type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepository {
	return &VideoRepository{
		db: db,
	}
}

func (r *VideoRepository) Create(video *models.Video) (*models.Video, error) {
	result := r.db.Create(video)
	if result.Error != nil {
		return nil, result.Error
	}
	return video, nil
}

func (r *VideoRepository) FetchLatestVideosByChannelID(channelID string, limit int) (videos models.VideoList, err error) {
	if err := r.db.
		Model(&models.Video{}).
		Where("channel_id = ?", channelID).
		Order("created_at desc").
		Limit(limit).
		Find(&videos.Data); err != nil {
		return videos, err.Error
	}
	return videos, nil
}

func (r *VideoRepository) Update(video *models.Video) error {
	result := r.db.Save(video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *VideoRepository) Delete(id int) error {
	result := r.db.Delete(&models.Video{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
