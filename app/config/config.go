package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/stivens13/spotter-assessment/app/helper/constants"
)

type SpotterAPIConfig struct {
	YoutubeConfig *YoutubeConfig
	DBConfig      *DBConfig
}

type ETLConfig struct {
	YoutubeConfig           *YoutubeConfig
	NewChannelAmount        int
	AverageVideosPerChannel int
}

type YoutubeConfig struct {
	APIKey string
	Host   string
	Port   string
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func (c *DBConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Los_Angeles",
		c.Host,
		c.User,
		c.Password,
		c.Database,
		c.Port,
	)
}

func GetETLConfig() *ETLConfig {
	newChannelAmountStr := os.Getenv("NEW_CHANNEL_AMOUNT")
	newChannelAmount, err := strconv.Atoi(newChannelAmountStr)
	if err != nil {
		newChannelAmount = constants.DEFAULT_NEW_CHANNELS
	}

	averageVideosPerChannelStr := os.Getenv("AVERAGE_VIDEOS_PER_CHANNEL")
	averageVideosPerChannel, err := strconv.Atoi(averageVideosPerChannelStr)
	if err != nil {
		averageVideosPerChannel = constants.DEFAULT_AVG_VIDEOS_PER_CHANNEL
	}

	return &ETLConfig{
		YoutubeConfig: &YoutubeConfig{
			APIKey: os.Getenv("YOUTUBE_API_KEY"),
			Host:   os.Getenv("YOUTUBE_HOST"),
			Port:   os.Getenv("YOUTUBE_PORT"),
		},
		NewChannelAmount:        newChannelAmount,
		AverageVideosPerChannel: averageVideosPerChannel,
	}
}

func GetSpotterAPIConfig() *SpotterAPIConfig {
	return &SpotterAPIConfig{
		YoutubeConfig: &YoutubeConfig{
			APIKey: os.Getenv("YOUTUBE_API_KEY"),
			Host:   os.Getenv("YOUTUBE_HOST"),
			Port:   os.Getenv("YOUTUBE_PORT"),
		},
		DBConfig: &DBConfig{
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			Database: os.Getenv("POSTGRES_DB"),
		},
	}
}
