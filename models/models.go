package models

import "commerce/db"

type ListingDetailArgs struct {
	Listing       db.Listing
	Owner         db.User
	Category      db.Category
	Bids          []db.Bid
	MaxBid        db.Bid
	IsInWatchlist bool
	IsActive      bool
	Comments      []db.Comment
}
