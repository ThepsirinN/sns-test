package config

type CacheConfig struct {
	Database int `env:"CONFIG.CACHE.DATABASE"`
}

type CacheSecret struct {
	Host string `env:"SECRET.CACHE.HOST"`
	Port string `env:"SECRET.PORT.PORT"`
}
