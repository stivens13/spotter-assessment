package models

type Video struct {
	VideoID    string `json:"video_id"`
	VideoTitle string `json:"video_title"`
	UploadDate string `json:"upload_date"`
}
