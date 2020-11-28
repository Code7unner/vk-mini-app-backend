package models

import "github.com/go-pg/pg/v10"

type SteamImpl interface {
}

type Steam struct {
	tableName struct{} `pg:"steams,alias:c"` //nolint
	ID        int      `json:"id"`
}

type SteamRepo struct {
	db         *pg.DB
}

func NewSteamModel(db *pg.DB) *SteamRepo {
	return &SteamRepo{db}
}
