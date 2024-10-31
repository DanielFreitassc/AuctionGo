package main

import (
	auction_controller "auction_golang/internal/infra/api"
	auction_repository_impl "auction_golang/internal/infra/repository/auction"
	create_auction_use_case "auction_golang/usecase/auction_usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	repository := &auction_repository_impl.AuctionReposityImpl{}
	usecase := create_auction_use_case.NewAuctionUseCase(repository)
	controller := auction_controller.AuctionController{
		CreateAuctionUseCase: usecase,
	}

	router.POST("/auction", controller.Create)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
