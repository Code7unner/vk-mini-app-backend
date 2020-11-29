package models

import (
	"github.com/go-pg/pg/v10"
)

type MatchImpl interface {
	Get(id int) (*Match, bool)
	Create(s *Match) (*Match, error)
	GetALl() ([]Match, error)
}

type Match struct {
	tableName      struct{}  `pg:"matches,alias:c"` //nolint
	ID             int       `json:"id"`
	TeamLeftID     int       `json:"team_left_id"`
	TeamRightID    int       `json:"team_right_id"`
	TimeCreated    string `json:"time_created"`
	TimeStarted    string `json:"time_started"`
	TeamLeftReady  bool      `json:"team_left_ready"`
	TeamRightReady bool      `json:"team_right_ready"`
}

type MatchRepo struct {
	db *pg.DB
}

func NewMatchModel(db *pg.DB) *MatchRepo {
	return &MatchRepo{db}
}

func (r *MatchRepo) Get(id int) (*Match, bool) {
	match := &Match{}
	err := r.db.Model(match).Where("id = ?", id).Select()
	if err != nil {
		return match, false
	}

	return match, true
}

func (r *MatchRepo) Create(m *Match) (*Match, error) {
	_, err := r.db.Model(m).Insert()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *MatchRepo) GetALl() ([]Match, error) {
	matches := []Match{}
	if err := r.db.Model(&matches).Select(); err != nil {
		return matches, err
	}

	return matches, nil
}
