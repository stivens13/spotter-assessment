package repository_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stivens13/spotter-assessment/app/helper/constants"
	"github.com/stivens13/spotter-assessment/app/models"
	"github.com/stivens13/spotter-assessment/app/repository"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/google/go-cmp/cmp"
)

type VideoRepositorySuite struct {
	suite.Suite
	db   *gorm.DB
	repo *repository.VideoRepository
}

func (s *VideoRepositorySuite) SetupTest() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Los_Angeles",
		"localhost",
		"postgres",
		"postgres",
		"postgres",
		"5432",
	)
	var err error
	s.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	s.repo = repository.NewVideoRepository(s.db)
}

func TestVideoRepositorySuite(t *testing.T) {
	suite.Run(t, new(VideoRepositorySuite))
}

func (s *VideoRepositorySuite) TestVideoRepository_FetchLatestVideosByChannelID() {
	limit := constants.LATEST_VIDEO_LIMIT
	opts := cmpopts.IgnoreFields(models.Video{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt", "UploadDate")
	// opts := cmp.Options{
	// 	cmpopts.IgnoreFields(models.Video{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
	// }
	tests := map[string]struct {
		name      string
		channelID string
		expected  models.VideoList
		errWanted bool
	}{
		"Valid case - Channel 1": {
			channelID: "UC6qq5ZRn_epjgdKwtgmeSd3",
			expected: models.VideoList{
				Data: []models.Video{
					{VideoID: "aWyckSMztx1", ChannelID: "UC6qq5ZRn_epjgdKwtgmeSd3", VideoTitle: "Example1 Video 1", UploadDate: models.NewDate("2024-03-21")},
					{VideoID: "aWyckSMztx2", ChannelID: "UC6qq5ZRn_epjgdKwtgmeSd3", VideoTitle: "Example1 Video 2", UploadDate: models.NewDate("2024-03-20")},
					{VideoID: "aWyckSMztx3", ChannelID: "UC6qq5ZRn_epjgdKwtgmeSd3", VideoTitle: "Example1 Video 3", UploadDate: models.NewDate("2024-03-19")},
					{VideoID: "aWyckSMztx4", ChannelID: "UC6qq5ZRn_epjgdKwtgmeSd3", VideoTitle: "Example1 Video 4", UploadDate: models.NewDate("2024-03-18")},
					{VideoID: "aWyckSMztx5", ChannelID: "UC6qq5ZRn_epjgdKwtgmeSd3", VideoTitle: "Example1 Video 5", UploadDate: models.NewDate("2024-03-17")},
				},
			},
			errWanted: false,
		},

		"Valid case - Channel 2": {
			channelID: "UC0032Wkd3aCT4rRi1YOVc2d",
			expected: models.VideoList{
				Data: []models.Video{
					{VideoID: "bc92oS2pxg1", ChannelID: "UC0032Wkd3aCT4rRi1YOVc2d", VideoTitle: "Example2 Video 1", UploadDate: models.NewDate("2024-03-19")},
					{VideoID: "bc92oS2pxg2", ChannelID: "UC0032Wkd3aCT4rRi1YOVc2d", VideoTitle: "Example2 Video 2", UploadDate: models.NewDate("2024-03-18")},
					{VideoID: "bc92oS2pxg3", ChannelID: "UC0032Wkd3aCT4rRi1YOVc2d", VideoTitle: "Example2 Video 3", UploadDate: models.NewDate("2024-03-18")},
					{VideoID: "bc92oS2pxg4", ChannelID: "UC0032Wkd3aCT4rRi1YOVc2d", VideoTitle: "Example2 Video 4", UploadDate: models.NewDate("2024-03-17")},
					{VideoID: "bc92oS2pxg5", ChannelID: "UC0032Wkd3aCT4rRi1YOVc2d", VideoTitle: "Example2 Video 5", UploadDate: models.NewDate("2024-03-16")},
				},
			},
			errWanted: false,
		},

		"Valid case - Channel 3": {
			channelID: "UCTRQblH_muP2X68UsLwFm2G",
			expected: models.VideoList{
				Data: []models.Video{
					{VideoID: "zjk3M4hNjG1", ChannelID: "UCTRQblH_muP2X68UsLwFm2G", VideoTitle: "Example3 Video 1", UploadDate: models.NewDate("2024-03-19")},
					{VideoID: "zjk3M4hNjG2", ChannelID: "UCTRQblH_muP2X68UsLwFm2G", VideoTitle: "Example3 Video 2", UploadDate: models.NewDate("2024-03-18")},
					{VideoID: "zjk3M4hNjG3", ChannelID: "UCTRQblH_muP2X68UsLwFm2G", VideoTitle: "Example3 Video 3", UploadDate: models.NewDate("2024-03-18")},
					{VideoID: "zjk3M4hNjG4", ChannelID: "UCTRQblH_muP2X68UsLwFm2G", VideoTitle: "Example3 Video 4", UploadDate: models.NewDate("2024-03-17")},
					{VideoID: "zjk3M4hNjG5", ChannelID: "UCTRQblH_muP2X68UsLwFm2G", VideoTitle: "Example3 Video 5", UploadDate: models.NewDate("2024-03-16")},
				},
			},
			errWanted: false,
		},

		// Repository should return empty resonse with no error
		// Usecase must handle empty response and return 404
		"Valid case - missing channel ID": {
			channelID: "invalid",
			expected:  models.VideoList{Data: []models.Video{}},
			errWanted: false,
		},
	}

	for name, test := range tests {
		s.Run(name, func() {
			got, err := s.repo.FetchLatestVideosByChannelID(test.channelID, limit)

			if test.errWanted {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			diff := cmp.Diff(got.Data, test.expected.Data, opts)
			if diff != "" {
				s.FailNowf("got mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
