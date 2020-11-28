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
	u.GET("/login/:id", userHandler.Login)
	u.POST("/register", userHandler.RegisterUser)

	teamHandler := handlers.NewTeamHandler(a)
	t := e.Group("/team")
	t.GET("/:id", teamHandler.GetTeam)
	t.POST("/create", teamHandler.CreateTeam)

	steamHandler := handlers.NewSteamHandler(a, host, token)
	s := e.Group("/steam")
	s.GET("/login", steamHandler.Login)
	s.GET("/callback", steamHandler.Callback)
	s.GET("/user/:id", steamHandler.GetSteamUser)

	return e
}
