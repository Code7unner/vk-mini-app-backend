package server

import (
	"github.com/code7unner/vk-mini-app-backend/internal/app"
	"github.com/code7unner/vk-mini-app-backend/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(a app.Application, host string, token string) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{"*"}}))

	userHandler := handlers.NewUserHandler(a)
	u := e.Group("/user")
	u.POST("/register", userHandler.RegisterUser)
	u.GET("/login", userHandler.Login)

	teamHandler := handlers.NewTeamHandler(a)
	t := e.Group("/team")
	t.GET("/", teamHandler.GetTeam)
	t.POST("/create", teamHandler.CreateTeam)

	steamHandler := handlers.NewSteamHandler(a, host, token)
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
