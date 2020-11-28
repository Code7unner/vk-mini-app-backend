package handlers

import (
	"github.com/code7unner/vk-mini-app-backend/internal/app"
	"github.com/code7unner/vk-mini-app-backend/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TeamHandler struct {
	app app.Application
}

func NewTeamHandler(a app.Application) *TeamHandler {
	return &TeamHandler{app: a}
}

func (h *TeamHandler) CreateTeam(c echo.Context) error {
	cookie, err := c.Cookie("user_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}
	userID, _ := strconv.Atoi(cookie.Value)

	team := new(models.Team)
	if err := c.Bind(team); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	t, err := h.app.GetTeam(team.ID)
	switch err {
	case app.ErrTeamNotFound:
		t, err = h.app.CreateTeam(team, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, t)
	case nil:
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
	id, err := strconv.Atoi(cookie.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}
	t, err := h.app.GetTeam(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, t)
}
