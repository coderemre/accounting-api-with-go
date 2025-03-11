package database

import (
	"database/sql"
	"time"

	"accounting-api-with-go/internal/config"
	"accounting-api-with-go/internal/utils"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() *sql.DB {
	var err error

	cfg := config.LoadConfig()

	DB, err = sql.Open("mysql", cfg.DatabaseDSN)
	if err != nil {
		utils.Log.Fatal().Err(err).Msg(utils.ErrDatabaseConnectionFailed.String())
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	err = DB.Ping()
	if err != nil {
		utils.Log.Fatal().Err(err).Msg(utils.ErrDatabasePingFailed.String())
	}

	utils.Log.Info().Msg(utils.SuccessDatabaseConnected.String())

	return DB
}

func Close() error {
	if DB != nil {
		utils.Log.Info().Msg(utils.SuccessDatabaseDisconnected.String())
		return DB.Close()
	}
	return nil
}