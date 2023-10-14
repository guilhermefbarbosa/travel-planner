package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseType string

const (
	Postgres DatabaseType = "postgres"
	Mysql    DatabaseType = "mysql"
	Mongo    DatabaseType = "mongo"
	Redis    DatabaseType = "redis"
)

type Config struct {
	Host     string       `env:"DB_HOST,required"`
	User     string       `env:"DB_USER,required"`
	Password string       `env:"DB_PASSWORD,required"`
	Name     string       `env:"DB_NAME,required"`
	Port     string       `env:"DB_PORT,required"`
	Type     DatabaseType `env:"DB_TYPE,required"`
}

type Database struct {
	conn *gorm.DB
}

func NewDatabase(c Config) *Database {
	dialector, err := getDialectorByDBType(c)
	if err != nil {
		panic(err)
	}
	gormConfig := gorm.Config{}
	conn, err := gorm.Open(dialector, &gormConfig)
	if err != nil {
		panic(err)
	}

	return &Database{
		conn: conn,
	}
}

func (d *Database) Conn() *gorm.DB {
	return d.conn
}

func getDialectorByDBType(c Config) (gorm.Dialector, error) {
	switch c.Type {
	case Postgres:
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			c.Host,
			c.User,
			c.Password,
			c.Name,
			c.Port,
		)
		return postgres.Open(dsn), nil
	case Mysql:
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			c.User,
			c.Password,
			c.Host,
			c.Port,
			c.Name,
		)
		return mysql.Open(dsn), nil
	case Mongo:
		return nil, errors.New("not implemented")
	case Redis:
		return nil, errors.New("not implemented")
	default:
		return nil, errors.New("invalid database type")
	}
}
