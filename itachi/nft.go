package itachi

import (
	"context"
	"time"
)

// NFT represents a token which may be claimed or unclaimed
type NFT struct {
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ID            string    `json:"id"`
	Hash          string    `json:"hash"`
	Address       string    `json:"address"`
	Asset         string    `json:"asset"`
	OriginalOwner string    `json:"original_owner"`
	CurrentOwner  string    `json:"current_owner"`
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
