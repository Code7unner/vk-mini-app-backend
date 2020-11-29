package handlers

import (
	"github.com/code7unner/vk-mini-app-backend/internal/app"
	"github.com/code7unner/vk-mini-app-backend/internal/models"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type MatchHandler struct {
	app app.Application
}

func NewMatchHandler(a app.Application) *MatchHandler {
	return &MatchHandler{
		app: a,
	}
}

func (h *MatchHandler) CreateMatch(c echo.Context) error {
	match := new(models.Match)
	if err := c.Bind(match); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	match.TimeCreated = time.Now().String()
	match.ID = rand.Int()
	m, err := h.app.GetMatch(match.ID)
	switch err {
	case app.ErrMatchNotFound:
		m, err = h.app.CreateMatch(match)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, m)
	case nil:
		return c.JSON(http.StatusOK, m)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse("unexpected error"))
	}
}

func (h *MatchHandler) GetMatch(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	m, err := h.app.GetMatch(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, m)
}

func (h *MatchHandler) GetAllMatches(c echo.Context) error {
	matches, err := h.app.GetAllMatches()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, matches)
}
