package models

import (
	"TOPO/common"
	"github.com/google/uuid"
)

const (
	UserName = "name"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Email    string    `json:"email"`
	Street   string    `json:"street"`
}

func NewUserModel(id uuid.UUID, name string, lastName string, email string, street string) *User {
	um := &User{
		ID:       id,
		LastName: lastName,
		Email:    email,
		Name:     name,
		Street:   street,
	}
	return um
}

func (u *User) ToEntity() *UserEntity {
	ue := &UserEntity{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		LastName: u.LastName,
		Street:   u.Street,
	}
	return ue
}

func ToUserModel(ue *UserEntity) *User {
	um := &User{
		ID:       ue.ID,
		Name:     ue.Name,
		Email:    ue.Email,
		LastName: ue.LastName,
		Street:   ue.Street,
	}
	return um
}

type UserList struct {
	Items []User `json:"items"`
	Count int64  `json:"count"`
}

type UserQuery struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Email    string    `json:"email"`
	common.BaseQuery
}
