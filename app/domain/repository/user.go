package repository

import "golang-ddd-api/app/domain/model"

type UserRepository interface {
	Search(name string) (*model.User, error)
}
