package app

import (
	"github.com/code7unner/vk-mini-app-backend/internal/models"
)

type Application interface {
	GetUser(id int) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)

	GetTeam(id int) (*models.Team, error)
	CreateTeam(team *models.Team) (*models.Team, error)
}

type App struct {
	userModel models.UserImpl
	teamModel models.TeamImpl
}

func New(user models.UserImpl, team models.TeamImpl) Application {
	return &App{
		userModel: user,
		teamModel: team,
	}
}

func (a App) GetUser(id int) (*models.User, error) {
	user, ok := a.userModel.Get(id)
	if !ok {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (a App) CreateUser(user *models.User) (*models.User, error) {
	u, err := a.userModel.Create(user)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a App) GetTeam(id int) (*models.Team, error) {
	team, ok := a.teamModel.Get(id)
	if !ok {
		return nil, ErrTeamNotFound
	}

	return team, nil
}

func (a App) CreateTeam(team *models.Team) (*models.Team, error) {
	t, err := a.teamModel.Create(team)
	if err != nil {
		return nil, err
	}

	return t, nil
}
