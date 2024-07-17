package repository

import (
	"gorm.io/gorm"
)

type Repository[T any] interface {
	GetAll() ([]T, error)
	GetById(id string) (T, error)
	Create(entity T) error
	Update(entity T) error
	Delete(id string) error
}

type GormRepository[T any] struct {
	db *gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}

func (r *GormRepository[T]) GetAll() ([]T, error) {
	var entities []T
	result := r.db.Find(&entities)
	return entities, result.Error
}

func (r *GormRepository[T]) GetById(id string) (T, error) {
	var entity T
	result := r.db.First(&entity, id)
	return entity, result.Error
}

func (r *GormRepository[T]) Create(entity T) error {
	result := r.db.Create(&entity)
	return result.Error
}

func (r *GormRepository[T]) Update(entity T) error {
	result := r.db.Save(&entity)
	return result.Error
}

func (r *GormRepository[T]) Delete(id string) error {
	var entity T
	result := r.db.Delete(&entity, id)
	return result.Error
}
