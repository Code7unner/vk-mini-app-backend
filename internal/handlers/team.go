package handlers

import (
	"github.com/code7unner/vk-mini-app-backend/internal/app"
	"github.com/code7unner/vk-mini-app-backend/internal/auth"
	"github.com/code7unner/vk-mini-app-backend/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TeamHandler struct {
	app       app.Application
	authorize auth.Auth
}

func NewTeamHandler(a app.Application, authorize auth.Auth) *TeamHandler {
	return &TeamHandler{app: a, authorize: authorize}
}

func (h *TeamHandler) CreateTeam(c echo.Context) error {
	cookie, err := c.Cookie("user_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}
	userID, err := h.authorize.GetUserID(cookie)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	team := new(models.Team)
	if err := c.Bind(team); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	token, err := h.authorize.SetUserToken(team.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	t, err := h.app.GetTeam(team.ID)
	switch err {
	case app.ErrTeamNotFound:
		t, err = h.app.CreateTeam(team, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		}
		c.SetCookie(h.authorize.NewCookie("team_id", token))
		return c.JSON(http.StatusOK, t)
	case nil:
		c.SetCookie(h.authorize.NewCookie("team_id", token))
		return c.JSON(http.StatusOK, t)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse("unexpected error"))
	}
}

func (h *TeamHandler) GetTeam(c echo.Context) error {
	cookie, err := c.Cookie("team_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}
	id, err := h.authorize.GetTeamID(cookie)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}
	t, err := h.app.GetTeam(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, t)
}

func (h *TeamHandler) GetAllTeams(c echo.Context) error {
	teams, err := h.app.GetAllTeams()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, teams)
}


