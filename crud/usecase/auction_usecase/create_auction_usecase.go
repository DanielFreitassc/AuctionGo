package create_auction_use_case

import "auction_golang/internal/entity/auction_entity"

type AuctionInputDto struct {
	ProductName string                          `json:"product_name"`
	Condition   auction_entity.AuctionCondition `json:"condition"`
}

type AuctionOutputDto struct {
	Id          string
	ProductName string
	Condition   auction_entity.AuctionCondition
}

type AuctionUseCase struct {
	AuctionReposity auction_entity.AuctionRepositoryInterface
}

func NewAuctionUseCase(repository auction_entity.AuctionRepositoryInterface) *AuctionUseCase {
	return &AuctionUseCase{
		AuctionReposity: repository,
	}
}

func (u AuctionUseCase) CreateAuction(input *AuctionInputDto) (*AuctionOutputDto, error) {
	a := auction_entity.CreateAuction(input.ProductName, input.Condition)

	if err := u.AuctionReposity.CreateAuction(a); err != nil {
		return nil, err

	}

	return &AuctionOutputDto{
		Id:          a.Id,
		ProductName: a.ProductName,
		Condition:   a.Condition,
	}, nil
}
