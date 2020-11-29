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
	u.GET("/", userHandler.GetUser)
	u.GET("/all", userHandler.GetAllUsers)
	u.POST("/register", userHandler.RegisterUser)

	teamHandler := handlers.NewTeamHandler(a, authorize)
	t := e.Group("/team")
	t.GET("/:id", teamHandler.GetTeam)
	t.GET("/all", teamHandler.GetAllTeams)
	t.POST("/create/:id", teamHandler.CreateTeam)

	steamHandler := handlers.NewSteamHandler(a, url, cfg.SteamToken)
	s := e.Group("/steam")
	s.GET("/login", steamHandler.Login)
	s.GET("/callback", steamHandler.Callback)
	s.GET("/user", steamHandler.GetSteamUser)

	matchHandler := handlers.NewMatchHandler(a)
	m := e.Group("/match")
	m.GET("/:id", matchHandler.GetMatch)
	m.POST("/create", matchHandler.CreateMatch)
	m.POST("/all", matchHandler.GetAllMatches)

	return e
}
