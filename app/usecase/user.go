package usecase

import (
	"golang-ddd-api/app/domain/model"
	"golang-ddd-api/app/domain/repository"
)

type UserUseCase interface {
	Search(name string) (*model.User, error)
}

type userUserCase struct {
	userRepository repository.UserRepository
}

func NewUserUserCase(ur repository.UserRepository) UserUseCase {
	return &userUserCase{
		userRepository: ur,
	}
}

func (uu userUserCase) Search(name string) (user *model.User, err error) {
	result, err := uu.userRepository.Search(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}
