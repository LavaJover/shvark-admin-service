package handlers

import (
	"net/http"
	"strconv"

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

// @Summary Register new trader team
// @Description Register new trader team by credentials
// @Tags traders
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body dto.RegisterTraderRequest true "trader credentials"
// @Success 201 {object} dto.RegisterTraderResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 502 {object} dto.ErrorResponse
// @Router /admin/traders/register [post]
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

func (h *TraderHandler) GetTraders(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	traders, totalPages, err := h.traderUsecase.GetTraders(int64(page), int64(limit))
	if err != nil {
		c.JSON(http.StatusBadGateway, dto.ErrorResponse{Error: err.Error()})
		return
	}

	var tradersResponse []*dto.Trader
	for _, trader := range traders {
		tradersResponse = append(tradersResponse, &dto.Trader{
			ID: trader.ID,
			Username: trader.Username,
			Login: trader.Login,
			Password: trader.Password,
		})
	}
	
	c.JSON(http.StatusOK, dto.GetTradersResponse{
		Traders: tradersResponse,
		TotalPages: totalPages,
	})
}