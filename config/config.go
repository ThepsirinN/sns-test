package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sns-barko/constant"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Log      Log
	Database DatabaseConfig
	Cache    CacheConfig
}

type Secret struct {
	App      AppSecret
	Database DatabaseSecret
	Cache    CacheSecret
}

var cfg Config
var secret Secret

func InitConfig() (*Config, *Secret) {
	err := initConfig()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	return &cfg, &secret
}

func initConfig() error {
	configPath, err := filepath.Abs(constant.CONFIG_PATH + constant.CONFIG_FILE)
	if err != nil {
		return err
	}
	secretPath, err := filepath.Abs(constant.SECRET_PATH + constant.SECRET_FILE)
	if err != nil {
		return err
	}

	if err := godotenv.Load(configPath, secretPath); err != nil {
		return err
	}

	if err := env.Parse(&cfg); err != nil {
		return err
	}

	if err := env.Parse(&secret); err != nil {
		return err
	}

	fmt.Println(os.Environ())

	fmt.Println(cfg)
	fmt.Println(secret)

	return nil
}
