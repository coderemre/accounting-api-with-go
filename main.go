package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"accounting-api-with-go/internal/tracing"

	"accounting-api-with-go/internal/config"
	"accounting-api-with-go/internal/database"
	"accounting-api-with-go/internal/middlewares"
	"accounting-api-with-go/internal/utils"
	"accounting-api-with-go/routes"
)

func startMetricsServer(port string) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":" + port, nil)
		if err != nil {
			panic("Metrics server failed: " + err.Error())
		}
	}()
}

func main() {
	cfg := config.LoadConfig()
	// Prepare migration source URL
	// sourceURL := "./migrations"
	// databaseURL := os.Getenv("DATABASE_DSN")
	// if !strings.Contains(databaseURL, "://") {
	// 	databaseURL = "mysql://" + databaseURL
	// }
	// m, err := migrate.New(sourceURL, databaseURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatal(err)
	// }

	ctx := context.Background()
	shutdown := tracing.InitTracer(ctx, "accounting-api")
	defer shutdown(ctx)

	startMetricsServer(cfg.MetricsPort)

	utils.InitLogger(cfg.Port)
	utils.Log.Info().Msg(utils.SuccessLoggerInitialized.String())

	var db = database.Connect()

	router := routes.SetupRoutes(db)
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