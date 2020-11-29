package models

import (
	"github.com/go-pg/pg/v10"
)

type TeamImpl interface {
	GetAll() ([]Team, error)
	SetMatchID(teamID, matchID int) error
	Get(id int) (*Team, bool)
	Create(team *Team) (*Team, error)
}

type Team struct {
	tableName   struct{} `pg:"teams,alias:c"` //nolint
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Tag         string   `json:"tag"`
	PhotoSmall  string   `json:"photo_100"`
	PhotoMedium string   `json:"photo_200"`
	PhotoBig    string   `json:"photo_max_orig"`
	Rating      int      `json:"rating"`
	MatchID     int      `json:"match_id"`
}

type TeamRepo struct {
	db *pg.DB
}

func NewTeamModel(db *pg.DB) *TeamRepo {
	return &TeamRepo{db}
}

func (r *TeamRepo) GetAll() ([]Team, error) {
	teams := []Team{}
	if err := r.db.Model(&teams).Select(); err != nil {
		return teams, err
	}

	return teams, nil
}

func (r *TeamRepo) SetMatchID(teamID, matchID int) error {
	team := &Team{}
	_, err := r.db.Model(team).
		Set("match_id = ?", matchID).
		Where("id = ?", teamID).
		UpdateNotZero()
	if err != nil {
		return err
	}

	return nil
}

func (r *TeamRepo) Get(id int) (*Team, bool) {
	team := &Team{}
	err := r.db.Model(team).Where("id = ?", id).Select()
	if err != nil {
		return team, false
	}

	return team, true
}

func (r *TeamRepo) Create(t *Team) (*Team, error) {
	_, err := r.db.Model(t).Insert()
	if err != nil {
		return nil, err
	}

	return t, nil
}
