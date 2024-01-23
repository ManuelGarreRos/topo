package fixtures

import (
	"TOPO/appctr"
	"TOPO/internal/models"
	"TOPO/internal/repositories"
	"TOPO/internal/services"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type userFixture struct {
	us *services.UserService

	log *zap.Logger

	users map[string]models.User
}

var userFxtr = userFixture{}

func MakeUsers() {
	db := appctr.DB()
	log := appctr.Log()
	ur := repositories.NewUserRepository(db, log)

	userFxtr.us = services.NewUserService(ur, log)
	userFxtr.log = log
	userFxtr.DoUserFixtures()
}

func (uF userFixture) DoUserFixtures() {
	uF.log.Info("Doing user fixtures")
	uF.createUsers()
}

func (uF userFixture) createUsers() {
	uF.log.Info("Creating users")
	user1 := models.NewUserModel(
		uuid.New(),
		"Juan",
		"Perez",
		"juanperez@gmail.com",
		"Av. Siempre Viva 123",
	)
	user2 := models.NewUserModel(
		uuid.New(),
		"Pedro",
		"Picapiedra",
		"pedropicapiedra@gmail.com",
		"Av. Siempre Viva 1234",
	)
	user3 := models.NewUserModel(
		uuid.New(),
		"Vilma",
		"Picapiedra",
		"vilapicapiedra@gmail.com",
		"Av. Siempre Viva 12345",
	)
	user4 := models.NewUserModel(
		uuid.New(),
		"Pablo",
		"Marmol",
		"pablomarmol@gmail.com",
		"Av. Siempre Viva 123456",
	)
	user5 := models.NewUserModel(
		uuid.New(),
		"Betty",
		"Marmol",
		"bettymarmol@gmail.com",
		"Av. Siempre Viva 1234567",
	)

	uF.users = map[string]models.User{
		"juan":  *user1,
		"pedro": *user2,
		"vilma": *user3,
		"pablo": *user4,
		"betty": *user5,
	}

	for _, u := range uF.users {
		if err := uF.us.Create(&u); err != nil {
			uF.log.Fatal("Error creating user", zap.Error(err))
		}
	}
}
