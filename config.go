package main

import (
	"fmt"
)

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (c PostgresConfig) Dialect() string {
	return "postgres"
}

func (c PostgresConfig) ConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "secret",
		Name:     "go_web_dev",
	}
}

// isProd := false
//
// Port = 3000
type Config struct {
	Port int
	Env  string
}

func (c Config) IsProd() bool {
	return c.Env == "prod"
}

func DefaultConfig() Config {
	return Config{
		Port: 3000,
		Env:  "dev",
	}
}

// fmt.Println("Starting the server on :3000...")
// http.ListenAndServe(":3000", csrfMw(userMw.Apply(r)))

// # models/users.go

// userPwPepper  = "secret-random-string"
// hmacSecretKey = "secret-hmac-key"

// # models/services.go

// db, err := gorm.Open("postgres", connectionInfo)
// //....
// db.LogMode(true)
