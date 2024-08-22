package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stivens13/spotter-assessment/app/models"
	"github.com/stivens13/spotter-assessment/app/usecase"
)

type ChannelHandler struct {
	ChannelUsecase *usecase.ChannelInteractor
}

func NewChannelHandler(
	e *echo.Echo,
	channelUsecase *usecase.ChannelInteractor,
) *ChannelHandler {
	channelHandler := &ChannelHandler{
		ChannelUsecase: channelUsecase,
	}
	e.GET("/channels/:channel_id", channelHandler.FetchChannel)
	e.POST("/channel", channelHandler.CreateChannel)
	e.POST("/channels", channelHandler.CreateChannels)

	return channelHandler
}

func (vh *ChannelHandler) FetchChannel(c echo.Context) error {
	channelID := c.Param("channel_id")
	channels, err := vh.ChannelUsecase.FetchChannel(channelID)
	if err != nil {
		if err.Error() == "no channels found" {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, channels)
}

func (vh *ChannelHandler) CreateChannel(c echo.Context) error {
	var channel models.Channel
	if err := c.Bind(&channel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := vh.ChannelUsecase.Create(&channel); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

func (vh *ChannelHandler) CreateChannels(c echo.Context) error {
	var channels models.ChannelRawList
	if err := c.Bind(&channels); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := vh.ChannelUsecase.CreateBatch(channels); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
