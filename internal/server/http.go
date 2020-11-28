package server

import (
	"fmt"
	"github.com/code7unner/vk-mini-app-backend/config"
	"github.com/code7unner/vk-mini-app-backend/internal/app"
	"github.com/code7unner/vk-mini-app-backend/internal/auth"
	"github.com/code7unner/vk-mini-app-backend/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(a app.Application, cfg *config.Config) *echo.Echo {
	url := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{"*"}}))

	authorize := auth.New(cfg.SteamToken)

	userHandler := handlers.NewUserHandler(a, authorize)
	u := e.Group("/user")
	u.POST("/register", userHandler.RegisterUser)
	u.GET("/login", userHandler.Login)

	teamHandler := handlers.NewTeamHandler(a, authorize)
	t := e.Group("/team")
	t.GET("/", teamHandler.GetTeam)
	t.POST("/create", teamHandler.CreateTeam)

	steamHandler := handlers.NewSteamHandler(a, url, cfg.SteamToken)
	s := e.Group("/steam")
	s.GET("/login", steamHandler.Login)
	s.GET("/callback", steamHandler.Callback)
	s.GET("/user", steamHandler.GetSteamUser)

	matchHandler := handlers.NewMatchHandler(a)
	m := e.Group("/match")
	m.GET("/", matchHandler.GetMatch)
	m.POST("/create", matchHandler.CreateMatch)

	return e
}
