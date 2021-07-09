package service

import (
	"github.com/LaiJunBin/gin-api/internal/model"
	"github.com/LaiJunBin/gin-api/internal/requests"
	"github.com/LaiJunBin/gin-api/pkg/utils"
)

func (s Service) CreateUser(params *requests.RegisterUserRequest) (model.User, error) {
	hashedPassword, _ := utils.HashPassword(params.Password)
	return s.Dao.CreateUser(params.Name, params.Username, hashedPassword)
}