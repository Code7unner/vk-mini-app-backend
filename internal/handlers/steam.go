package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/code7unner/vk-mini-app-backend/client"
	"github.com/code7unner/vk-mini-app-backend/internal/app"
	"github.com/code7unner/vk-mini-app-backend/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/yohcop/openid-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type SteamHandler struct {
	app            app.Application
	host           string
	steamToken     string
	httpClientPool *client.HttpClientPool
}

func NewSteamHandler(a app.Application, h string, s string) *SteamHandler {
	return &SteamHandler{
		app:            a,
		host:           h,
		steamToken:     s,
		httpClientPool: client.NewHttpClientPool(),
	}
}

func (h *SteamHandler) Login(c echo.Context) error {
	const id = "http://steamcommunity.com/openid"
	realm := fmt.Sprintf("http://%s/", h.host)
	callbackURL := fmt.Sprintf("http://%s/steam/callback", h.host)
	url, err := openid.RedirectURL(id, callbackURL, realm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}

	userID := c.Param("user_id")

	cookie := &http.Cookie{
		Name:    "user_id",
		Value:   userID,
		Expires: <-time.After(time.Hour * 48),
	}
	c.SetCookie(cookie)

	return c.Redirect(http.StatusSeeOther, url)
}

func (h *SteamHandler) Callback(c echo.Context) error {
	uri := fmt.Sprintf("http://%s%s", h.host, c.Request().URL.String())
	steamURL, err := openid.Verify(
		uri,
		openid.NewSimpleDiscoveryCache(),
		openid.NewSimpleNonceStore(),
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}

	urlPaths := strings.Split(steamURL, "/")
	id := urlPaths[len(urlPaths)-1]

	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", h.steamToken, id)
	cl := h.httpClientPool.Get()
	resp, err := cl.Get(url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}
	h.httpClientPool.Put(cl)
	defer resp.Body.Close()

	var steamData models.SteamData
	if err := json.NewDecoder(resp.Body).Decode(&steamData); err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}

	player := steamData.Response.Players[0]
	playerID, _ := strconv.Atoi(player.SteamID)
	steam := &models.Steam{
		ID:                       playerID,
		CommunityVisibilityState: player.CommunityVisibilityState,
		ProfileState:             player.ProfileState,
		PersonaName:              player.PersonaName,
		CommentPermission:        player.CommentPermission,
		ProfileURL:               player.ProfileURL,
		Avatar:                   player.Avatar,
		AvatarMedium:             player.AvatarMedium,
		AvatarFull:               player.AvatarFull,
		AvatarHash:               player.AvatarHash,
		LastLogoff:               player.LastLogoff,
		PersonaState:             player.PersonaState,
		RealName:                 player.RealName,
		PrimaryClanID:            player.PrimaryClanID,
		TimeCreated:              player.TimeCreated,
		PersonaStateFlags:        player.PersonaStateFlags,
		LocCountryCode:           player.LocCountryCode,
	}

	cookie, err := c.Cookie("user_id")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}
	userID, _ := strconv.Atoi(cookie.Value)

	s, err := h.app.GetSteamUser(playerID)
	switch err {
	case app.ErrSteamUserNotFound:
		s, err = h.app.CreateSteamUser(steam, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, s)
	case nil:
		return c.JSON(http.StatusOK, s)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse("unexpected error"))
	}
}

func (h *SteamHandler) GetSteamUser(c echo.Context) error {
	cookie, err := c.Cookie("steam_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	id, err := strconv.Atoi(cookie.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	u, err := h.app.GetSteamUser(id)
	switch err {
	case app.ErrSteamUserNotFound:
		return c.JSON(http.StatusInternalServerError, errorResponse("could not find user"))
	case nil:
		return c.JSON(http.StatusOK, u)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse("unexpected error"))
	}
}
