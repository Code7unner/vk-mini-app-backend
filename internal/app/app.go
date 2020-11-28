package app

import (
	"github.com/code7unner/vk-mini-app-backend/internal/models"
)

type Application interface {
	// Users
	GetUser(id int) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)

	// Teams
	GetTeam(id int) (*models.Team, error)
	CreateTeam(team *models.Team, userID int) (*models.Team, error)

	// Steams
	GetSteamUser(id int) (*models.Steam, error)
	CreateSteamUser(steam *models.Steam, userID int) (*models.Steam, error)
}

type App struct {
	userModel  models.UserImpl
	teamModel  models.TeamImpl
	steamModel models.SteamImpl
	matchModel models.MatchImpl
}

func New(user models.UserImpl, team models.TeamImpl, steam models.SteamImpl, match models.MatchImpl) Application {
	return &App{
		userModel:  user,
		teamModel:  team,
		steamModel: steam,
		matchModel: match,
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

func (a App) CreateTeam(team *models.Team, userID int) (*models.Team, error) {
	t, err := a.teamModel.Create(team)
	if err != nil {
		return nil, err
	}

	if err := a.userModel.SetTeamID(userID, team.ID); err != nil {
		return nil, err
	}

	return t, nil
}

func (a App) GetSteamUser(id int) (*models.Steam, error) {
	steam, ok := a.steamModel.Get(id)
	if !ok {
		return nil, ErrSteamUserNotFound
	}

	return steam, nil
}

func (a App) CreateSteamUser(steam *models.Steam, userID int) (*models.Steam, error) {
	s, err := a.steamModel.Create(steam)
	if err != nil {
		return nil, err
	}

	if err := a.userModel.SetSteamID(userID, steam.ID); err != nil {
		return nil, err
	}

	return s, nil
}
