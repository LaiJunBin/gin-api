package dao

import (
	"github.com/LaiJunBin/gin-api/internal/model"
	"gorm.io/gorm"
)

func (dao *Dao) CreateUser(name, username, password string) (model.User, error) {
	user := model.User{
		Name: name,
		Username: username,
		Password: password,
	}

	return user.Create(dao.Engine)
}

func (dao *Dao) UpdateUser(id uint, newUser model.User) error {
	user := model.User{
		Model: gorm.Model{
			ID: id,
		},
	}

	return user.Update(dao.Engine, newUser)
}

func (dao *Dao) DeleteUser(id uint) error {
	user := model.User{
		Model: gorm.Model{
			ID: id,
		},
	}

	return user.Delete(dao.Engine)
}

func (dao *Dao) GetUserByUsername(username string) (model.User, error) {
	user := model.User{
		Username: username,
	}

	return user.Get(dao.Engine)
}

func (dao *Dao) GetUserByID(id uint) (model.User, error) {
	user := model.User{
		Model: gorm.Model{
			ID: id,
		},
	}

	return user.Get(dao.Engine)
}