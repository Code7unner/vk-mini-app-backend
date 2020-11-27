package server

import (
	"github.com/code7unner/vk-mini-app-backend/internal/app"
	"github.com/code7unner/vk-mini-app-backend/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(a app.Application) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{"*"}}))

	userHandler := handlers.NewUserHandler(a)
	u := e.Group("/user")

	u.GET("/login/:id", userHandler.Login)
	u.POST("/register", userHandler.RegisterUser)

	return e
}
