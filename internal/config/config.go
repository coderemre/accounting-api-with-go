package config

import (
	"os"
	"strconv"

	"accounting-api-with-go/internal/utils"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	MetricsPort       string
	RedisAddress	string
	DatabaseDSN string
	TransactionLimits GlobalTransactionLimit
}

type GlobalTransactionLimit struct {
	MaxAmount float64
	MaxCount  int
	Period    string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		utils.Log.Warn().Err(err).Msg(utils.ErrEnvFileNotFound.String())
	}

	config := &Config{
		Port:       getEnv("PORT", "8080"),
		MetricsPort:       getEnv("METRICS_PORT", "2112"),
		RedisAddress:       getEnv("REDIS_PORT", "6379"),
		DatabaseDSN: getEnv("DATABASE_DSN", "<username>:<password>@tcp(localhost:3306)/bank_app"),
		TransactionLimits: LoadGlobalTransactionLimit(),
	}

	utils.Log.Info().Msg(utils.SuccessConfigLoaded.String())
	return config
}

func LoadGlobalTransactionLimit() GlobalTransactionLimit {
	maxAmountStr := getEnv("MAX_TX_AMOUNT", "1000")
	maxCountStr := getEnv("MAX_TX_COUNT", "5")
	period := getEnv("TX_PERIOD", "daily")

	maxAmount, _ := strconv.ParseFloat(maxAmountStr, 64)
	maxCount, _ := strconv.Atoi(maxCountStr)

	return GlobalTransactionLimit{
		MaxAmount: maxAmount,
		MaxCount:  maxCount,
		Period:    period,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		utils.Log.Info().Msgf("%s environment variable loaded successfully", key)
		return value
	}
	utils.Log.Warn().Msgf("%s environment variable not set, using fallback value: %s", key, fallback)
	return fallback
}