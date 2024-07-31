package config

type AppConfig struct {
}

type AppSecret struct {
	Port int `env:"SECRET.APP.PORT"`
}
