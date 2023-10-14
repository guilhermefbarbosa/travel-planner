package main

import (
	"errors"
	"fmt"
	"github.com/guilhermefbarbosa/travel-planner/api/travel-planner-go/internal/config"
	"github.com/guilhermefbarbosa/travel-planner/api/travel-planner-go/pkg/database"
	"github.com/guilhermefbarbosa/travel-planner/api/travel-planner-go/src/structs"
)

// This file is used to create the database schema.
// As this is a simple project, this is not considered production ready for not being prepared for schema evolution.
func main() {
	cfg := config.NewConfig()
	conn := database.NewDatabase(cfg.Database).Conn()
	migrator := database.NewMigrator(conn)

	_ = prepareUUID(cfg.Database, migrator)

	if err := MigrateAll(migrator); err != nil {
		fmt.Printf("migration failed, err: %s, \ndropping all tables...\n", err)
		if err := DropAll(migrator); err != nil {
			fmt.Printf("drop failed, err: %s\n", err)
			return
		}
		fmt.Printf("succesful table drop")
		return
	}
	fmt.Println("successful migration")

}

func MigrateAll(migrator *database.Migrator) error {
	return migrator.Migrate(
		&structs.User{},
		&structs.Travel{},
		&structs.Address{},
		&structs.Payment{},
		&structs.Journey{},
		&structs.Session{},
		&structs.Accommodation{},
	)
}

func DropAll(migrator *database.Migrator) error {
	return migrator.Drop(
		&structs.Accommodation{},
		&structs.Session{},
		&structs.Journey{},
		&structs.Payment{},
		&structs.Address{},
		&structs.Travel{},
		&structs.User{},
	)
}

func prepareUUID(cfg database.Config, migrator *database.Migrator) error {
	switch cfg.Type {
	case database.Postgres:
		return migrator.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	case database.Mysql:
		return migrator.Exec(`CREATE TYPE 'uuid' AS varchar(36)`)
	default:
		return errors.New("database type not supported")
	}
}
