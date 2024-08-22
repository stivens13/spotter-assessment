package models

type Video struct {
	Base
	VideoMetadata
}

func NewVideo(videoID, channelID, title, date string) *Video {
	return &Video{
		VideoMetadata: VideoMetadata{
			VideoID:    videoID,
			ChannelID:  channelID,
			VideoTitle: title,
			UploadDate: NewDate(date),
		},
	}
}

type VideoList struct {
	Data []*Video `json:"data"`
}

func (vl *VideoList) Make(videos []*Video) *VideoList {
	return &VideoList{
		Data: videos,
	}
}

type VideoMetadata struct {
	VideoID    string `json:"video_id" gorm:"uniqueIndex; type:varchar(11)"`
	ChannelID  string `json:"channel_id" gorm:"index; type:varchar(24)"`
	VideoTitle string `json:"video_title" gorm:"type:varchar(255)"`
	UploadDate Date   `json:"upload_date" gorm:"type:date; index:,sort:desc"`
}

// type VideoMetadataList struct {
// 	Data []*VideoMetadata `json:"data"`
// }

type VideoMap struct {
	Data map[string]VideoList `json:"data"`
}
