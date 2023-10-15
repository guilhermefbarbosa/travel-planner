package repositories

import (
	"errors"
	"github.com/google/uuid"
	"github.com/guilhermefbarbosa/travel-planner/api/travel-planner-go/pkg/database"
	"gorm.io/gorm"
)

type Repository struct {
	conn *gorm.DB
}

func NewBaseRepository() *Repository {
	return &Repository{
		conn: database.Get().Conn(),
	}
}

func Create[T database.Tables](obj T) (T, error) {
	return obj, database.Get().Conn().Create(obj).Error
}

func Update[T database.Tables](obj T) (T, error) {
	result := database.Get().Conn().Updates(obj)
	if result.Error != nil {
		return *new(T), result.Error
	}
	if result.RowsAffected == 0 {
		return *new(T), errors.New("not found")
	}
	return obj, nil
}

func Delete[T database.Tables](obj T) error {
	result := database.Get().Conn().Delete(obj)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}

func Get[T database.Tables](id uuid.UUID) (T, error) {
	var obj T
	var zero T
	err := database.Get().Conn().First(&obj, id).Error
	if err == gorm.ErrRecordNotFound {
		return zero, errors.New("not found")
	}
	return obj, database.Get().Conn().First(&obj, id).Error
}

func GetByFilter[T database.Tables](clauses ...database.Clause) (T, error) {
	var obj T
	conn := database.Get().Conn()
	conn = addClauses(conn, clauses...)
	return obj, conn.Find(obj).Error
}

func GetMany[T database.Tables](clauses ...database.Clause) ([]T, error) {
	var objs []T
	conn := database.Get().Conn()
	conn = addClauses(conn, clauses...)
	return objs, conn.Find(objs).Error
}

func addClauses(conn *gorm.DB, clauses ...database.Clause) *gorm.DB {
	for _, clause := range clauses {
		conn = clause.Apply(conn)
	}
	return conn
}
