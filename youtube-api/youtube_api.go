package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/stivens13/spotter-assessment/tools/generator"
)

type YoutubeHandler struct {
}

func NewYoutubeHandler(
	e *echo.Echo,
) *YoutubeHandler {
	youtubeHandler := &YoutubeHandler{
	}

	e.GET("/api:channel_id", youtubeHandler.FetchVideoMetadataByChannel)
	return youtubeHandler
}


func (yh *YoutubeHandler) FetchVideoMetadataByChannel(c echo.Context) error {
	channelID := c.Param("channel_id")
	videos := generator.GenerateMockVideosMetadata(channelID)
	return c.JSON(200, videos)
}

type Services struct {
	Server *echo.Echo
}

func InitServices() *Services {
	echo := echo.New()
	return &Services{
		Server: echo,
	}
}

func main() {
	fmt.Println("Youtube Mock API starting...")
	services := InitServices()
	fmt.Println("Services successfully initialized")
	server := services.Server
	fmt.Println("Starting server")
	server.Logger.Fatal(server.Start(":9000"))
}
