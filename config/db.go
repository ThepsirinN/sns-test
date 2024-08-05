package config

import "time"

type DatabaseConfig struct {
	Name            string        `env:"CONFIG.DATABASE.NAME"`
	ConnMaxLifeTime time.Duration `env:"CONFIG.DATABASE.ConnMaxLifeTime"`
	MaxOpenConns    int           `env:"CONFIG.DATABASE.MaxOpenConns"`
	MaxIdleConns    int           `env:"CONFIG.DATABASE.MaxIdleConns"`
}

type DatabaseSecret struct {
	Host     string `env:"SECRET.DATABASE.HOST"`
	Port     string `env:"SECRET.DATABASE.PORT"`
	User     string `env:"SECRET.DATABASE.USER"`
	Password string `env:"SECRET.DATABASE.PASSWORD"`
}
