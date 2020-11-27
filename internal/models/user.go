package models

import (
	"github.com/go-pg/pg/v10"
)

type UserImpl interface {
	Get(id int) (*User, bool)
	Create(u *User) (*User, error)
	Update(u *User) (*User, error)
	Insert(m *User) error
}

type User struct {
	tableName   struct{} `pg:"users,alias:c"` //nolint
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Lastname    string   `json:"lastname"`
	City        string   `json:"city"`
	Country     string   `json:"country"`
	Sex         int      `json:"sex"`
	Timezone    int      `json:"timezone"`
	PhotoSmall  string   `json:"photo_100"`
	PhotoMedium string   `json:"photo_200"`
	PhotoBig    string   `json:"photo_max_orig"`
	TeamID      int      `json:"team_id"`
}

type UserRepo struct {
	db *pg.DB
}

func NewUserModel(db *pg.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Get(id int) (*User, bool) {
	user := &User{}
	err := r.db.Model(user).Where("id = ?", id).Select()
	if err != nil {
		return user, false
	}

	return user, true
}

func (r *UserRepo) Create(u *User) (*User, error) {
	_, err := r.db.Model(u).Insert()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepo) Update(u *User) (*User, error) {
	_, err := r.db.Model(u).WherePK().UpdateNotZero()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepo) Insert(u *User) error {
	_, err := r.db.Model(u).
		OnConflict("DO NOTHING").
		Insert()
	if err != nil {
		return err
	}

	return nil
}
