package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"accounting-api-with-go/internal/config"
	"accounting-api-with-go/internal/database"
	"accounting-api-with-go/internal/utils"
	"accounting-api-with-go/middlewares"
	"accounting-api-with-go/routes"
)

func main() {
	database.RunMigrations()
	cfg := config.LoadConfig()

	utils.InitLogger(cfg.Port)
	utils.Log.Info().Msg(utils.SuccessLoggerInitialized.String())

	database.Connect()

	router := routes.SetupRoutes()
	router.Use(middlewares.Logger)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		utils.Log.Info().Msg(utils.SuccessServerRunning.String())
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Log.Fatal().Err(err).Msg(utils.ErrServerListenFailed.String())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utils.Log.Info().Msg(utils.SuccessServerShutdown.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		utils.Log.Fatal().Err(err).Msg(utils.ErrServerShutdownFailed.String())
	}

	if err := database.Close(); err != nil {
		utils.Log.Warn().Err(err).Msg(utils.ErrDatabaseCloseFailed.String())
	}

	utils.Log.Info().Msg(utils.SuccessServerExited.String())
}