package handlers

import (
	"net/http"

	"github.com/LavaJover/shvark-admin-service/internal/domain"
	"github.com/LavaJover/shvark-admin-service/internal/http/dto"
	"github.com/gin-gonic/gin"
)

type TraderHandler struct {
	traderUsecase domain.TraderUsecase
}

func NewTraderHandler(traderUsecase domain.TraderUsecase) *TraderHandler{
	return &TraderHandler{
		traderUsecase: traderUsecase,
	}
}

func (h *TraderHandler) RegisterTrader(c *gin.Context) {
	var request dto.RegisterTraderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	trader := &domain.Trader{
		Username: request.Username,
		Login: request.Login,
		Password: request.Password,
	}
	if err := h.traderUsecase.RegisterNewTrader(trader); err != nil {
		c.JSON(http.StatusBadGateway, dto.ErrorResponse{Error: "failed to register new trader"})
		return
	}

	c.JSON(http.StatusCreated, dto.RegisterTraderResponse{})
}