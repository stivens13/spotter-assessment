package generator

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jaswdr/faker/v2"
	"github.com/mazen160/go-random"
	"github.com/stivens13/spotter-assessment/app/helper/constants"
	"golang.org/x/exp/rand"
)

var (
	fakeit       = gofakeit.New(0)
	fakerr       = faker.New()
	timeNow      = time.Now()
	timeMonthAgo = timeNow.AddDate(0, -1, 0)
	letter       = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

const (
	ChannelIDLength   = 24
	VideoIDLength     = 11
	MaxVideoWordCount = 7
)

type MockVideo struct {
	VideoID      string
	ChannelID    string
	VideoTitle   string
	UploadedDate string
}

type MockChannel struct {
	ChannelID string
}

func GenerateMockChannel() *MockChannel {
	return &MockChannel{
		ChannelID: RandomString(ChannelIDLength),
	}
}

func GenerateMockVideo(channelID string) *MockVideo {
	return &MockVideo{
		VideoID:      RandomString(VideoIDLength),
		ChannelID:    channelID,
		VideoTitle:   GenerateMockSentenceFakeit(GenerateRandomSmallInt()),
		UploadedDate: GenerateMockDate(),
	}
}

func GenerateMockVideos(channelID string) (videos []*MockVideo) {
	for _ = range(10) {
		videos = append(videos, GenerateMockVideo(channelID))
	}
	return videos
}

func GenerateSecureID(n int) string {
	data, err := random.String(n)
	if err != nil {
		fmt.Printf("failed to generate secure string: %v", err)
	}
	return data
}

func GenerateRandomSmallInt() int {
	data, err := random.GetInt(MaxVideoWordCount)
	if err != nil {
		fmt.Printf("failed to generate secure string: %v", err)
	}
	return data
}

func GenerateMockDate() string {
	return fakeit.DateRange(timeMonthAgo, timeNow).Format(constants.VIDEO_DATE_FORMAT)
}

func GenerateMockSentenceFakeit(n int) string {
	return fakeit.Sentence(n)
}

func GenerateYoutubeVideoIDFaker() string {
	return fakerr.YouTube().GenerateVideoID()
}

func RandomString(n int) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
