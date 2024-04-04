package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gestaltintel/itachinft/internal/validator"
	"github.com/gestaltintel/itachinft/itachi"
	"github.com/go-chi/chi/v5"
)

func (svc *service) CreateNFTHandler() http.HandlerFunc {
	type input struct {
		Hash    string `json:"hash"`
		Address string `json:"address"`
		Asset   string `json:"asset"`
		Owner   string `json:"owner"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var in input
		err := ReadJSON(r, &in)
		if err != nil {
			svc.Send400Error(w, r, err)
			return
		}

		v := validator.New()
		v.Check(in.Hash != "", "hash", "must be provided")
		v.Check(len(in.Address) != 42, "address", "must be 42 character long")
		v.Check(in.Asset != "", "asset", "must be provided")
		v.Check(in.Owner != "", "owner", "must be provided")

		if !v.Valid() {
			svc.Send422Error(w, r, v.Errors)
			return
		}

		err = SendJSON(w, http.StatusOK, nil, capsule{"user_input": in})
		if err != nil {
			svc.Send500Error(w, r, err)
		}
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
			svc.Send500Error(w, r, err)
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
