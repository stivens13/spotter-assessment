package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Video struct {
	ID         uuid.UUID `gorm:"type:uuid;primarykey;default:gen_random_uuid()"`
	VideoID    string    `json:"video_id" gorm:"uniqueIndex; type:varchar(11)"`
	ChannelID  string    `json:"channel_id" gorm:"index; type:varchar(24)"`
	VideoTitle string    `json:"video_title" gorm:"type:varchar(255)"`
	UploadDate string    `json:"upload_date" gorm:"type:date; index:,sort:desc"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type VideoList struct {
	Data []Video `json:"data"`
}

type Channel struct {
	ID        uuid.UUID `gorm:"type:uuid;primarykey;default:gen_random_uuid()"`
	ChannelID string    `json:"channel_id" gorm:"uniqueIndex; type:varchar(24)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
