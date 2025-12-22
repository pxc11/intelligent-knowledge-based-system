package config

import (
	"time"
)

type Config struct {
	App   AppConfig
	JWT   JWTConfig
	MySQL MySQLConfig
	Redis RedisConfig
}

type AppConfig struct {
	Name string
	Env  string
	Port int
}

type JWTConfig struct {
	Secret string
	Expire time.Duration
}

type MySQLConfig struct {
	Host         string
	Port         int
	DB           string
	Password     string
	Username     string
	MaxIdleConns int
	MaxOpenConns int
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}
