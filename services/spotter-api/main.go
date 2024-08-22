package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stivens13/spotter-assessment/config"
	"github.com/stivens13/spotter-assessment/helper/constants"
	"github.com/stivens13/spotter-assessment/services/spotter-api/handler"
	"github.com/stivens13/spotter-assessment/services/spotter-api/repository"
	"github.com/stivens13/spotter-assessment/services/spotter-api/usecase"
	youtubeclient "github.com/stivens13/spotter-assessment/services/youtube-client"
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

func InitServices() *Services {
	dbCfg := config.GetDBConfig()
	db, err := openDBConnection(dbCfg)
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}


	youtubeCfg := config.GetYoutubeConfig()
	youtubeClient := youtubeclient.NewYoutubeClient(youtubeCfg)
	youtubeRepo := repository.NewYoutubeRepository(youtubeClient)
	videoRepo := repository.NewVideoRepository(db)
	videoUsecase := usecase.NewVideoInteractor(videoRepo, youtubeRepo)

	channelRepo := repository.NewChannelRepository(db)
	channelUsecase := usecase.NewChannelInteractor(channelRepo, videoUsecase, youtubeRepo)

	echo := echo.New()

	videoHandler := handler.NewVideoHandler(echo, videoUsecase)
	channelHandler := handler.NewChannelHandler(echo, channelUsecase)

	echo.GET("/health", Healthcheck)


	return &Services{
		DB:             db,
		Server:         echo,
		VideoHandler:   videoHandler,
		ChannelHandler: channelHandler,
	}
}

func Healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func main() {
	fmt.Println("Spotter API starting...")

	services := InitServices()
	fmt.Println("Services successfully initialized")
	server := services.Server
	fmt.Println("Starting server")
	server.Logger.Fatal(server.Start(":8080"))
}
