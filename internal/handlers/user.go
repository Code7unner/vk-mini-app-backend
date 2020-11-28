package handlers

import (
	"github.com/code7unner/vk-mini-app-backend/internal/app"
	"github.com/code7unner/vk-mini-app-backend/internal/auth"
	"github.com/code7unner/vk-mini-app-backend/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	app       app.Application
	authorize auth.Auth
}

func NewUserHandler(a app.Application, authorize auth.Auth) *UserHandler {
	return &UserHandler{app: a, authorize: authorize}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	token, err := h.authorize.SetUserToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	u, err := h.app.GetUser(user.ID)
	switch err {
	case app.ErrUserNotFound:
		u, err = h.app.CreateUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		}
		c.SetCookie(h.authorize.NewCookie("user_id", token))
		return c.JSON(http.StatusOK, u)
	case nil:
		c.SetCookie(h.authorize.NewCookie("user_id", token))
		return c.JSON(http.StatusOK, u)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse("unexpected error"))
	}
}

func (h *UserHandler) Login(c echo.Context) error {
	cookie, err := c.Cookie("user_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}
	userID, err := h.authorize.GetUserID(cookie)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}
	u, err := h.app.GetUser(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, u)
}
