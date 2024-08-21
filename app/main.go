package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stivens13/spotter-assessment/app/models"
	"github.com/stivens13/spotter-assessment/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func openDBConnection(dbConfig *config.DBConfig) (*gorm.DB, error) {
	dsn := dbConfig.GetDSN()

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %w", err)
    }

    return db, nil
}

func getMostRecentVideos(c echo.Context) error {
	return c.JSON(http.StatusOK, []models.Video{})
}

func main() {
	e := echo.New()

	e.GET("/videos/:channel_id", getMostRecentVideos)

	e.Logger.Fatal(e.Start(":8080"))
}
