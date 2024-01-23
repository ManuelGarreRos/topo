package sessions

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewSessionRepository(db *gorm.DB, log *zap.Logger) *SessionRepository {
	return &SessionRepository{
		db:  db,
		log: log,
	}
}

type SessionRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func (sr SessionRepository) ByID(id string) (*SessionEntity, error) {
	s := &SessionEntity{}

	if err := sr.db.Model(s).Where("id = ?", id).First(s).Error; err != nil {
		return nil, err
	}

	return s, nil
}

func (sr SessionRepository) Create(s *SessionEntity) error {
	if err := sr.db.Create(s).Error; err != nil {
		return err
	}

	return nil
}

func (sr SessionRepository) Delete(id string) error {
	if err := sr.db.Where("id = ?", id).Delete(&SessionEntity{}).Error; err != nil {
		return err
	}

	return nil
}
