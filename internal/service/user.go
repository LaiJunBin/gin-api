package service

import (
	"github.com/LaiJunBin/gin-api/internal/model"
	"github.com/LaiJunBin/gin-api/internal/requests"
	"github.com/LaiJunBin/gin-api/pkg/app"
	"github.com/LaiJunBin/gin-api/pkg/utils"
)

func (s Service) CreateUser(params *requests.RegisterUserRequest) (model.User, error) {
	hashedPassword, _ := utils.HashPassword(params.Password)
	return s.Dao.CreateUser(params.Name, params.Username, hashedPassword)
}

func (s Service) UpdateUser(id uint, params *requests.UpdateUserRequest) error {
	newUser := model.User{}

	if params.Name != "" {
		newUser.Name = params.Name
	}

	if params.Username != "" {
		newUser.Username = params.Username
	}

	if params.Password != "" {
		newUser.Password, _ = utils.HashPassword(params.Password)
	}

	return s.Dao.UpdateUser(id, newUser)
}

func (s Service) DeleteUser(id uint) error {
	return s.Dao.DeleteUser(id)
}

func (s Service) GetUserById(id uint) (model.User, error) {
	return s.Dao.GetUserByID(id)
}

func (s Service) GetUserByUsername(username string) (model.User, error) {
	user, err := s.Dao.GetUserByUsername(username)
	return user, err
}

func (s Service) CheckPermission(id uint, claims *app.Claims) bool {
	user, err := s.GetUserById(id)
	if err != nil {
		return false
	}

	return user.Username == claims.Username
}