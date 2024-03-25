package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gestaltintel/itachinft/itachi"
	"github.com/go-chi/chi/v5"
)

func (svc *service) CreateNFTHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "create a new NFT")
	}
}

func (svc *service) ShowNFTHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := readIDParam(r)
		if err != nil {
			svc.logger.Error(err.Error())
			http.NotFound(w, r)
			return
		}

		nft := &itachi.NFT{
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			ID:            id,
			Hash:          "hash string",
			Address:       "address string",
			Asset:         "QmaLZZmsXbfjLMJG9KYZhw7QRMTsJDhxgn14z6snm8jhTg",
			OriginalOwner: "Sanzen",
			CurrentOwner:  "Sanzen",
		}

		err = SendJSON(w, http.StatusOK, nil, capsule{"nft": nft})
		if err != nil {
			svc.SendServerError(w, r, err)
		}
	}
}

func readIDParam(r *http.Request) (string, error) {
	id := chi.URLParam(r, "nftid")

	if id == "" {
		return "", errors.New("no nft id provided")
	}

	return id, nil
}
