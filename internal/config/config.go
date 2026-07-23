package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	DBSSLMode  string
}

func Load() (Config, error) {
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// O .env é opcional: em produção (ex: Docker) as envs vêm do ambiente do
	// processo. Só tentamos ler o arquivo se ele realmente existir.
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			return Config{}, err
		}
	}

	cfg := Config{
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBSSLMode:  viper.GetString("DB_SSLMODE"),
	}

	return cfg, nil
}

func (c Config) DatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSLMode,
	)
}
