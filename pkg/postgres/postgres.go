package postgres

import (
	"log"

	"github.com/ansxy/nagabelajar-be-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresClient(conf *config.Config) (*gorm.DB, error) {
	// logLevel := logger.Error
	// switch conf.
	connection, err := gorm.Open(postgres.Open(conf.PostgresConfig.URI), &gorm.Config{})

	if err != nil {
		log.Fatal("[postgres-client] error: ", err)
		return nil, err
	}

	return connection, nil
}
