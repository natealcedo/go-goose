package repository

import (
	"gorm.io/gorm"
)

type Repository[T any] interface {
	GetAll() ([]T, error)
	GetByID(id string) (T, error)
	Create(entity T) (T, error)
	Update(entity T) error
	DeleteByID(id string) error
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

func (r *GormRepository[T]) GetByID(id string) (T, error) {
	var entity T
	result := r.db.Where("id = ?", id).Find(&entity)
	if result.RowsAffected == 0 {
		return entity, gorm.ErrRecordNotFound
	}
	return entity, result.Error
}

func (r *GormRepository[T]) Create(entity T) (T, error) {
	result := r.db.Create(&entity)
	if result.Error != nil {
		return entity, result.Error
	}
	return entity, nil
}

func (r *GormRepository[T]) Update(entity T) error {
	result := r.db.Save(&entity)
	return result.Error
}

func (r *GormRepository[T]) DeleteByID(id string) error {
	var entity T
	result := r.db.Where("id = ?", id).Delete(&entity, id)
	return result.Error
}
