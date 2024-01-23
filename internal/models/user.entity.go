package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string    `gorm:"not null"`
	LastName string
	Email    string `gorm:"not null"`
	Street   string `gorm:"not null"`
	Column1  string `gorm:"not null;default:'valor_por_defecto'"`
}

func (u *UserEntity) TableName() string {
	return "users"
}
