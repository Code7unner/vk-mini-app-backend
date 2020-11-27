package main

import (
	"context"
	"fmt"
	"github.com/code7unner/vk-mini-app-backend/config"
	"github.com/code7unner/vk-mini-app-backend/internal/app"
	"github.com/code7unner/vk-mini-app-backend/internal/db"
	"github.com/code7unner/vk-mini-app-backend/internal/models"
	"github.com/code7unner/vk-mini-app-backend/internal/server"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func init() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Errorf("No .env file found: %s", err.Error())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Init configs
	cfg, err := config.New()
	if err != nil {
		logrus.Fatal(err)
		return
	}

	// Connect to db
	d, err := db.Connect(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	defer d.Close()

	uModel := models.NewUserModel(d)
	tModel := models.NewTeamModel(d)

	// Starting server
	srv := server.New(app.New(uModel, tModel))
	go func() {
		if err := srv.Start(fmt.Sprintf(":%s", cfg.ServerPort)); err != nil {
			srv.Logger.Info("shutting down the server")
		}
	}()

	// Wait for terminate server
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt)
	<-terminate
	signal.Stop(terminate)
	if err := srv.Shutdown(ctx); err != nil {
		srv.Logger.Fatal(err)
	}
	cancel()
}
