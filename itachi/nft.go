package itachi

import (
	"context"
	"time"
)

// NFT represents a token which may be claimed or unclaimed
type NFT struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ID            string
	Hash          string
	Address       string
	Asset         string
	OriginalOwner string
	CurrentOwner  string
}

// NFTService represents the data layer to manage NFTs
type NFTService interface {
	FindNFTByID(ctx context.Context, id string) (*NFT, error)
	FindNFTs(ctx context.Context, filter NFTFilter) ([]*NFT, error)
	CreateNFT(ctx context.Context, nft *NFT) error
	UpdateNFT(ctx context.Context, id string, update NFTUpdate) (*NFT, error)
	DeleteNFT(ctx context.Context, id string) error
}

type NFTFilter struct {
	ID            *string
	Hash          *string
	Address       *string
	OriginalOwner *string
	CurrentOwner  *string

	Offset int
	Limit  int
}

type NFTUpdate struct {
	CurrentOwner *string
}
