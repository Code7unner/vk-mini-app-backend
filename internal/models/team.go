package models

import "github.com/go-pg/pg/v10"

type TeamImpl interface {

}

type Team struct {
	tableName    struct{} `pg:"teams,alias:c"` //nolint
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Tag          string   `json:"tag"`
	Photo100     string   `json:"photo_100"`
	Photo200     string   `json:"photo_200"`
	PhotoMaxOrig string   `json:"photo_max_orig"`
	Rating       int      `json:"rating"`
}

type TeamRepo struct {
	db *pg.DB
}

func NewTeamModel(db *pg.DB) *TeamRepo {
	return &TeamRepo{db}
}
