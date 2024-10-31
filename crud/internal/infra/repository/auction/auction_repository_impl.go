package auction_repository_impl

import "auction_golang/internal/entity/auction_entity"

var auctions map[string]*auction_entity.Auction

type AuctionReposityImpl struct{}

func (r *AuctionReposityImpl) CreateAuction(auction *auction_entity.Auction) error {

	auctions[auction.Id] = auction
	return nil
}
