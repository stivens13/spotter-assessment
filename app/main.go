package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/stivens13/spotter-assessment/app/config"
	"github.com/stivens13/spotter-assessment/app/handler"
	"github.com/stivens13/spotter-assessment/app/helper/constants"
	"github.com/stivens13/spotter-assessment/app/repository"
	"github.com/stivens13/spotter-assessment/app/usecase"
	youtubeclient "github.com/stivens13/spotter-assessment/youtube-client"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func openDBConnection(dbConfig *config.DBConfig) (*gorm.DB, error) {
	dsn := dbConfig.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		CreateBatchSize: constants.DB_CREATE_BATCH_SIZE,
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db, nil
}

type Services struct {
	Server         *echo.Echo
	VideoHandler   *handler.VideoHandler
	ChannelHandler *handler.ChannelHandler
	DB             *gorm.DB
}

func InitServices(cfg *config.SpotterAPIConfig) *Services {
	db, err := openDBConnection(cfg.DBConfig)
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}

	youtubeClient := youtubeclient.NewYoutubeClient(cfg.YoutubeConfig)
	youtubeRepo := repository.NewYoutubeRepository(youtubeClient)
	videoRepo := repository.NewVideoRepository(db)
	videoUsecase := usecase.NewVideoInteractor(videoRepo, youtubeRepo)

	channelRepo := repository.NewChannelRepository(db)
	channelUsecase := usecase.NewChannelInteractor(channelRepo, videoUsecase, youtubeRepo)

	echo := echo.New()

	videoHandler := handler.NewVideoHandler(echo, videoUsecase)
	channelHandler := handler.NewChannelHandler(echo, channelUsecase)

	return &Services{
		DB:             db,
		Server:         echo,
		VideoHandler:   videoHandler,
		ChannelHandler: channelHandler,
	}
}

func main() {
	fmt.Println("Spotter API starting...")
	config := config.GetSpotterAPIConfig()
	fmt.Println("Config successfully loaded")

	services := InitServices(config)
	fmt.Println("Services successfully initialized")
	server := services.Server
	fmt.Println("Starting server")
	server.Logger.Fatal(server.Start(":8080"))
}
