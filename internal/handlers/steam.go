package handlers

import (
	"fmt"
	"github.com/code7unner/vk-mini-app-backend/internal/app"
	"github.com/labstack/echo/v4"
	"github.com/yohcop/openid-go"
	"net/http"
)

type SteamHandler struct {
	app app.Application
	host string
}

func NewSteamHandler(a app.Application, h string) *SteamHandler {
	return &SteamHandler{app: a, host: h}
}

func (h *SteamHandler) Login(c echo.Context) error {
	const id = "http://steamcommunity.com/openid"
	realm := fmt.Sprintf("http://%s/", h.host)
	callbackURL := fmt.Sprintf("http://%s/steam/callback", h.host)
	url, err := openid.RedirectURL(id, callbackURL,	realm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}

	return c.Redirect(http.StatusSeeOther, url)
}

func (h *SteamHandler) Callback(c echo.Context) error {
	url := fmt.Sprintf("http://%s%s", h.host, c.Request().URL.String())
	id, err := openid.Verify(
		url,
		openid.NewSimpleDiscoveryCache(),
		openid.NewSimpleNonceStore(),
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, id)
}
