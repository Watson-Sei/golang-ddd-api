package persistence

import (
	"golang-ddd-api/app/domain/model"
	"golang-ddd-api/app/domain/repository"

	"gorm.io/gorm"
)

type userPersistence struct {
	Conn *gorm.DB
}

func NewUserPersistence(conn *gorm.DB) repository.UserRepository {
	return &userPersistence{Conn: conn}
}

func (up *userPersistence) Search(name string) (*model.User, error) {
	var user model.User
	
	if err := up.Conn.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
