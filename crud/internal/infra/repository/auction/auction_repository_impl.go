package auction_repository_impl

import (
	"auction_golang/internal/entity/auction_entity"
	"fmt"
)

var auctions map[string]*auction_entity.Auction = make(map[string]*auction_entity.Auction)

type AuctionReposityImpl struct{}

func (r *AuctionReposityImpl) CreateAuction(auction *auction_entity.Auction) error {

	auctions[auction.Id] = auction
	fmt.Print(len(auctions))
	return nil
}
