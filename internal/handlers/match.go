package handlers

import (
	"github.com/code7unner/vk-mini-app-backend/internal/app"
	"github.com/labstack/echo/v4"
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
	panic("create match is not exists")
}
