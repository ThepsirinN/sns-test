package config

import (
	"time"
)

type User struct {
	JWT  JWT
	Hash Hash
}

type JWT struct {
	Secret      string        `env:"SECRET.USER.JWT.SECRET"`
	ExpDuration time.Duration `env:"SECRET.USER.JWT.Duration"`
}

type Hash struct {
	Secret string `env:"SECRET.USER.HASH.SECRET"`
	Cost   int    `env:"SECRET.USER.HASH.COST"`
}
