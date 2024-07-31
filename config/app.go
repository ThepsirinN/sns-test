package config

type AppConfig struct {
	Name string `env:"CONFIG.APP.NAME"`
}

type AppSecret struct {
	Port int `env:"SECRET.APP.PORT"`
}
