package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stivens13/spotter-assessment/app/usecase"
)

type VideoHandler struct {
	VideoUsecase *usecase.VideoInteractor
}

func NewVideoHandler(
	e *echo.Echo,
	videoUsecase *usecase.VideoInteractor,
) *VideoHandler {
	videoHandler := &VideoHandler{
		VideoUsecase: videoUsecase,
	}
	e.GET("/videos/:channel_id", videoHandler.GetMostRecentVideos)

	return &VideoHandler{
		VideoUsecase: videoUsecase,
	}
}

func (vh *VideoHandler) GetMostRecentVideos(c echo.Context) error {
	channelID := c.Param("channel_id")
	videos, err := vh.VideoUsecase.FetchLatestVideosByChannelID(channelID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, videos)
}