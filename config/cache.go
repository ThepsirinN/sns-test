package config

import "time"

type CacheConfig struct {
	Database      int           `env:"CONFIG.CACHE.DATABASE"`
	PoolMaxIdle   int           `env:"CONFIG.CACHE.POOL.MAX.IDLE"`
	PoolMaxActive int           `env:"CONFIG.CACHE.POOL.MAX.ACTIVE"`
	PoolTimeout   time.Duration `env:"CONFIG.CACHE.POOL.TIMEOUT"`
}

type CacheSecret struct {
	Host     string `env:"SECRET.CACHE.HOST"`
	Port     string `env:"SECRET.PORT.PORT"`
	Password string `env:"SECRET.PORT.PASSWORD"`
}
