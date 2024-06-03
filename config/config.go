package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresConfig      PostgresConfig
	FirebaseConfig      FirebaseConfig
	XenditConfig        XenditConfig
	SmartContractConfig SmartContractConfig
	FRONTENDURL         string
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

type SmartContractConfig struct {
	Dial                 string
	SmartContractAddress string
	Key                  string
}

type FirebaseConfig struct {
	FirebaseStorageBucket string
	ServiceAccountPath    string
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
			Host:     v.GetString("POSTGRES_HOST"),
			Port:     v.GetString("POSTGRES_PORT"),
			Database: v.GetString("POSTGRES_DATABASE"),
			User:     v.GetString("POSTGRES_USER"),
			Password: v.GetString("POSTGRES_PASSWORD"),
			SSLMode:  "disable",
			URI:      v.GetString("POSTGRES_URI"),
		},
		XenditConfig: XenditConfig{
			SecretKey:          v.GetString("XENDIT_SECRET_KEY"),
			SuccessRedirectURL: "http://localhost:3000/success",
			FailureRedirectURL: "http://localhost:3000/failure",
		},
		FirebaseConfig: FirebaseConfig{
			ServiceAccountPath:    v.GetString("FIREBASE_SERVICE_ACCOUNT_PATH"),
			FirebaseStorageBucket: v.GetString("FIREBASE_STORAGE_BUCKET"),
		},
		SmartContractConfig: SmartContractConfig{
			Dial:                 v.GetString("SMART_CONTRACT_DIAL"),
			SmartContractAddress: v.GetString("SMART_CONTRACT_ADDRESS"),
			Key:                  v.GetString("KEYSTORE_BLOCKCHAIN"),
		},
		FRONTENDURL: v.GetString("FRONTEND_URL"),
	}
}
