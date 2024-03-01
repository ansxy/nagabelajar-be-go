package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresConfig PostgresConfig
	XenditConfig   XenditConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	SSLMode  string
	URI      string
}

type XenditConfig struct {
	SecretKey          string
	SuccessRedirectURL string
	FailureRedirectURL string
}

func SetConfig() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}

	v := viper.GetViper()
	viper.AutomaticEnv()
	return &Config{
		PostgresConfig: PostgresConfig{
			Host:     "localhost",
			Port:     "5432",
			Database: "nagabelajar",
			User:     "ansar",
			Password: "ansar123",
			SSLMode:  "disable",
			URI:      "postgresql://ansar:ansar123@localhost:5432/nagabelajar?sslmode=disable",
		},
		XenditConfig: XenditConfig{
			SecretKey:          v.GetString("XENDIT_SECRET_KEY"),
			SuccessRedirectURL: "http://localhost:3000/success",
			FailureRedirectURL: "http://localhost:3000/failure",
		},
	}
}
