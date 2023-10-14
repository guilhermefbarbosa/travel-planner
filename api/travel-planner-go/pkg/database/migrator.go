package database

import (
	"gorm.io/gorm"
)

type Migrator struct {
	conn *gorm.DB
}

type Tables interface {
	TableName() string
}

func NewMigrator(conn *gorm.DB) *Migrator {
	return &Migrator{
		conn,
	}
}

func (m Migrator) Migrate(tables ...Tables) error {
	return m.conn.AutoMigrate(
		toAnySlice(tables)...,
	)
}

func (m Migrator) Drop(tables ...Tables) error {
	return m.conn.Migrator().DropTable(
		toAnySlice(tables)...,
	)
}

func (m Migrator) Exec(query string) error {
	return m.conn.Exec(query).Error
}

func (m Migrator) TableExists(table string) bool {
	return m.conn.Migrator().HasTable(table)
}

func (m Migrator) ColumnExists(dst interface{}, column string) bool {
	return m.conn.Migrator().HasColumn(dst, column)
}

func toAnySlice[T any](tables []T) []interface{} {
	var slc []any
	for _, table := range tables {
		slc = append(slc, table)
	}
	return slc
}
