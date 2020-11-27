package main

import (
	"context"
	"github.com/code7unner/vk-mini-app-backend/config"
	"github.com/code7unner/vk-mini-app-backend/internal/server"
	"github.com/code7unner/vk-mini-app-backend/scrapper"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Errorf("No .env file found: %s", err.Error())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.New()
	if err != nil {
		logrus.Fatal(err)
		return
	}

	// Starting scrapper
	scrap := scrapper.New(time.Duration(cfg.MinutesToScrap))
	scrap.Start(ctx)

	// Starting server
	srv := server.New()
	go func() {
		if err := srv.Start(":8081"); err != nil {
			srv.Logger.Info("shutting down the server")
		}
	}()

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt)
	<-terminate
	signal.Stop(terminate)
	if err := srv.Shutdown(ctx); err != nil {
		srv.Logger.Fatal(err)
	}
	cancel()
}
