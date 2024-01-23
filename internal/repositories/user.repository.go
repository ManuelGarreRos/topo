package repositories

import (
	"TOPO/internal/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB, log *zap.Logger) *UserRepository {
	return &UserRepository{
		db:  db,
		log: log,
	}
}

type UserRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func (ur UserRepository) ByID(id string) (*models.UserEntity, error) {
	u := &models.UserEntity{}

	if err := ur.db.First(u, id).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (ur UserRepository) PaginatedList(q *models.UserQuery) ([]models.UserEntity, error) {
	var ul []models.UserEntity

	if err := ur.db.
		Where(ur.query(q)).
		Limit(q.Pager.Limit).
		Offset(q.Pager.Offset).
		Find(&ul).Error; err != nil {
		return nil, err
	}

	return ul, nil
}

func (ur UserRepository) Create(u *models.UserEntity) error {
	if err := ur.db.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) Delete(id uuid.UUID) error {
	if err := ur.db.Delete(&models.UserEntity{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) Update(u *models.UserEntity) error {
	if err := ur.db.Save(u).Error; err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) query(q *models.UserQuery) *gorm.DB {
	db := ur.db

	if q.ID != uuid.Nil {
		db = db.Where("id = ?", q.ID)
	}

	if q.Name != "" {
		db = db.Where("name = ?", q.Name)
	}

	if q.LastName != "" {
		db = db.Where("last_name = ?", q.LastName)
	}

	if q.Email != "" {
		db = db.Where("email = ?", q.Email)
	}

	return db
}
