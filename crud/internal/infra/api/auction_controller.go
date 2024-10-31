package auction_controller

import (
	create_auction_use_case "auction_golang/usecase/auction_usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuctionController struct {
	CreateAuctionUseCase *create_auction_use_case.AuctionUseCase
}

func (controller *AuctionController) Create(c *gin.Context) {
	var input create_auction_use_case.AuctionInputDto

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(400, err.Error())
		return
	}

	output, err := controller.CreateAuctionUseCase.CreateAuction(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, output)
}
