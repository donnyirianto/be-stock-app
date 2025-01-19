package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppStruct struct {
	Port    int
	Version string
	Prefork bool
}

type DatabaseStruct struct {
	Host           string
	Name           string
	Username       string
	Port           int
	Password       string
	DbMaxIdle      int
	DbMaxOpenConns int
	DbMaxLifetime  int
}

type Config struct {
	App                        AppStruct
	Database                   DatabaseStruct
	JWTSecret                  string
	JWTSecretRefresh           string
	TokenExpiryDuration        int
	TokenRefreshExpiryDuration int
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("unable to read config file: %v", err)
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %v", err)
	}

	cfg.App = AppStruct{
		Version: viper.GetString("VERSION"),
		Port:    viper.GetInt("PORT"),
		Prefork: viper.GetBool("PREFORK"),
	}

	cfg.Database = DatabaseStruct{
		Host:           viper.GetString("DB_HOST"),
		Username:       viper.GetString("DB_USERNAME"),
		Password:       viper.GetString("DB_PASSWORD"),
		Port:           viper.GetInt("DB_PORT"),
		Name:           viper.GetString("DB_NAME"),
		DbMaxIdle:      viper.GetInt("DB_MAX_IDLE"),
		DbMaxOpenConns: viper.GetInt("DB_POOLING_LIMIT"),
	}

	cfg.JWTSecret = viper.GetString("JWT_SECRET")
	cfg.JWTSecretRefresh = viper.GetString("JWT_REFRESH_SECRET")
	cfg.TokenExpiryDuration = viper.GetInt("JWT_ACCESS_EXPIRATION")
	cfg.TokenRefreshExpiryDuration = viper.GetInt("JWT_REFRESH_EXPIRATION")

	return &cfg, nil
}
