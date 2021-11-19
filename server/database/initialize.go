package database

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

type Database struct {
	Handle *pg.DB
	config Configuration
}

func Initialize(conf Configuration) {
	var db Database

	db.Handle = pg.Connect(&pg.Options{
		Database: conf.Database,
		User:     conf.User,
		Password: conf.Password,
	})
	defer db.Handle.Close()

	m, err := migrate.New(
		"file://server/database/migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", conf.User, conf.Password, conf.Host, conf.Port, conf.Database))
	if err != nil {
		log.Fatalf("Failed to create new migrate instance: %v", err)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to migrate: %v", err)
	}
}
