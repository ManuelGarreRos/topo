package services

import (
	"TOPO/internal/models"
	"TOPO/internal/repositories"
	"go.uber.org/zap"
)

func NewUserService(userRepository *repositories.UserRepository, log *zap.Logger) *UserService {
	return &UserService{
		userRepository: userRepository,
		log:            log,
	}
}

type UserService struct {
	userRepository *repositories.UserRepository
	log            *zap.Logger
}

func (us UserService) ByID(id string) (*models.User, error) {
	u, err := us.userRepository.ByID(id)
	if err != nil {
		return nil, err
	}

	return models.ToUserModel(u), nil
}

func (us UserService) PaginatedList(q *models.UserQuery) ([]models.User, error) {
	ul, err := us.userRepository.PaginatedList(q)
	if err != nil {
		return nil, err
	}

	var uml []models.User
	for _, u := range ul {
		uml = append(uml, *models.ToUserModel(&u))
	}

	return uml, nil
}

func (us UserService) Create(u *models.User) error {
	if err := us.userRepository.Create(u.ToEntity()); err != nil {
		return err
	}

	return nil
}

func (us UserService) Delete(u *models.User) error {
	if err := us.userRepository.Delete(u.ID); err != nil {
		return err
	}

	return nil
}

func (us UserService) Update(u *models.User) error {
	if err := us.userRepository.Update(u.ToEntity()); err != nil {
		return err
	}

	return nil
}
