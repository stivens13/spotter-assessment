package models

type Channel struct {
	Base
	ChannelID string    `json:"channel_id" gorm:"uniqueIndex; type:varchar(24)"`
}

type ChannelRawList struct {
	Data []string `json:"data"`
}

type ChannelList struct {
	Data []*Channel `json:"data"`
}
