package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stivens13/spotter-assessment/tools/generator"
)

type YoutubeAPI struct {
}

func NewYoutubeAPI(
	e *echo.Echo,
) *YoutubeAPI {
	youtubeHandler := &YoutubeAPI{
	}

	e.GET("/api/:channel_id", youtubeHandler.FetchVideoMetadataByChannel)
	e.GET("/health", Healthcheck)
	return youtubeHandler
}

func Healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}


func (yh *YoutubeAPI) FetchVideoMetadataByChannel(c echo.Context) error {
	channelID := c.Param("channel_id")
	videos := generator.GenerateMockVideosMetadata(channelID)
	return c.JSON(200, videos)
}

type Services struct {
	Server *echo.Echo
	YoutubeAPI *YoutubeAPI
}

func InitServices() *Services {
	echo := echo.New()
	youtubeAPI := NewYoutubeAPI(echo)
	return &Services{
		Server: echo,
		YoutubeAPI: youtubeAPI,
	}
}

func main() {
	fmt.Println("Youtube API starting...")
	services := InitServices()
	fmt.Println("Services successfully initialized")
	server := services.Server
	fmt.Println("Starting server")
	server.Logger.Fatal(server.Start(":9000"))
}
