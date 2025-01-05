package config

import (
	"os"

	"accounting-api-with-go/internal/utils"
	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DatabaseDSN string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		utils.Log.Warn().Err(err).Msg(utils.ErrEnvFileNotFound.String())
	}

	config := &Config{
		Port:       getEnv("PORT", "4520"),
		DatabaseDSN: getEnv("DATABASE_DSN", "<username>:<password>@tcp(localhost:3306)/bank_app"),
	}

	utils.Log.Info().Msg(utils.SuccessConfigLoaded.String())
	return config
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		utils.Log.Info().Msgf("%s environment variable loaded successfully", key)
		return value
	}
	utils.Log.Warn().Msgf("%s environment variable not set, using fallback value: %s", key, fallback)
	return fallback
}