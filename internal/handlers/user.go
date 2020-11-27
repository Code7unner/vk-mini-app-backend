package handlers

import (
	"github.com/code7unner/vk-mini-app-backend/internal/app"
	"github.com/code7unner/vk-mini-app-backend/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler struct {
	app app.Application
}

func NewUserHandler(a app.Application) *UserHandler {
	return &UserHandler{app: a}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	u, err := h.app.GetUser(user.ID)
	switch err {
	case app.ErrUserNotFound:
		u, err = h.app.CreateUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, u)
	case nil:
		return c.JSON(http.StatusOK, u)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse("unexpected error"))
	}
}

func (h *UserHandler) Login(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}
	u, err := h.app.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, u)
}
