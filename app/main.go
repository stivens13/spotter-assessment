package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/stivens13/spotter-assessment/app/handler"
	"github.com/stivens13/spotter-assessment/app/repository"
	"github.com/stivens13/spotter-assessment/app/usecase"
	"github.com/stivens13/spotter-assessment/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func openDBConnection(dbConfig *config.DBConfig) (*gorm.DB, error) {
	dsn := dbConfig.GetDSN()

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    return db, nil
}


type Services struct {
	Server *echo.Echo
	VideoUsecase *usecase.VideoInteractor
	ChannelUsecase *usecase.ChannelInteractor
	VideoHandler *handler.VideoHandler
	DB *gorm.DB
}

func InitServices(cfg *config.Config) (*Services) {
	db, err := openDBConnection(cfg.DBConfig)
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}

	// if err = db.AutoMigrate(&models.Video{}, &models.Channel{}); err != nil {
	// 	log.Fatalf("failed to migrate Video table: %v", err)
	// }

	// db.Migrator().CreateConstraint(&models.Video{}, "Channels")
	// db.Migrator().CreateConstraint(&models.Video{}, "fk_video_channels")

	videoRepo := repository.NewVideoRepository(db)
	videoUsecase := usecase.NewVideoInteractor(videoRepo)

	channelRepo := repository.NewChannelRepository(db)
	channelUsecase := usecase.NewChannelInteractor(channelRepo)

	// if err := channelUsecase.Create(&models.Channel{
	// 	ChannelID: "142",
	// }); err != nil {
	// 	log.Fatalf("failed to create test video: %v", err)
	// }

	// if err = videoUsecase.Create(&models.Video{
	// 	VideoID: "41234",
	// 	ChannelID: "142",
	// 	VideoTitle: "Test Video",
	// 	UploadDate: "2021-01-01",
	// }); err != nil {
	// 	log.Fatalf("failed to create test video: %v", err)
	// }

	echo := echo.New()

	videoHandler := handler.NewVideoHandler(echo, videoUsecase)

	return &Services{
		DB: db,
		Server: echo,
		VideoUsecase: videoUsecase,
		ChannelUsecase: channelUsecase,
		VideoHandler: videoHandler,
	}
}

func main() {
	fmt.Println("Spotter API starting...")
	config := config.InitConfig()
	fmt.Println("Config successfully loaded")

	services := InitServices(config)
	fmt.Println("Services successfully initialized")
	server := services.Server
	fmt.Println("Starting server")
	server.Logger.Fatal(server.Start(":8080"))
}
