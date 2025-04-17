package repositories

import (
	"errors"
	"gorm.io/gorm"
)

type RepositoryImpl[T any] interface {
	Create(db *gorm.DB, payload map[string]interface{}) error
	CreateBatch(db *gorm.DB, payload []map[string]interface{}) error
	Update(db *gorm.DB, payload map[string]interface{}) error
	UpdateBatch(db *gorm.DB, updates []map[string]interface{}) error
	Delete(db *gorm.DB, id any) error
	GetByID(db *gorm.DB, id any) (*T, error)
	GetAll(db *gorm.DB, page, pageSize int) ([]T, int64, error)
	DynamicQuery(db *gorm.DB, payload map[string]string) ([]T, error)
	Search(db *gorm.DB, payload map[string]string, page, pageSize int) ([]T, int64, error)
}

type Repository[T any] struct {
}

func NewRepository[T any]() RepositoryImpl[T] {
	return &Repository[T]{}
}

func (r *Repository[T]) Create(db *gorm.DB, payload map[string]interface{}) error {
	return db.Model(new(T)).Create(payload).Error
}

// create function generic insert batch
func (r *Repository[T]) CreateBatch(db *gorm.DB, payload []map[string]interface{}) error {
	return db.Model(new(T)).Create(payload).Error
}

func (r *Repository[T]) Update(db *gorm.DB, payload map[string]interface{}) error {
	return db.Model(new(T)).Where("id", payload["id"]).Updates(payload).Error
}

func (r *Repository[T]) UpdateBatch(db *gorm.DB, updates []map[string]interface{}) error {
	for _, data := range updates {
		id, ok := data["id"]
		if !ok {
			return errors.New("failed to update order: missing id field")
		}
		err := db.Model(new(T)).Where("id = ?", id).Updates(data).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository[T]) Delete(db *gorm.DB, id any) error {
	return db.Model(new(T)).Where("id", id).Delete(new(T)).Error
}

func (r *Repository[T]) GetByID(db *gorm.DB, id any) (*T, error) {
	var entity T
	err := db.Model(new(T)).Where("id", id).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *Repository[T]) GetAll(db *gorm.DB, page, pageSize int) ([]T, int64, error) {
	var entities []T
	var total int64
	err := db.Model(new(T)).Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&entities).Error
	if err != nil {
		return nil, 0, err
	}
	return entities, total, nil
}

func (r *Repository[T]) DynamicQuery(db *gorm.DB, payload map[string]string) ([]T, error) {
	// example payload: map[string]string{"username": "admin"}
	var entities []T
	err := db.Model(new(T)).Where(payload).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *Repository[T]) Search(db *gorm.DB, payload map[string]string, page, pageSize int) ([]T, int64, error) {
	var entities []T
	var total int64

	query := db.Model(new(T))
	for key, value := range payload {
		query = query.Where(key+" LIKE ?", "%"+value+"%")
	}

	err := query.Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&entities).Error
	if err != nil {
		return nil, 0, err
	}

	return entities, total, nil
}
