package api

import (
	"github.com/LaiJunBin/gin-api/internal/requests"
	"github.com/LaiJunBin/gin-api/internal/service"
	"github.com/LaiJunBin/gin-api/pkg/app"
	"github.com/LaiJunBin/gin-api/pkg/errors"
	"github.com/gin-gonic/gin"
)

type User struct {}

func NewUser() User {
	return User{}
}

func (u User) Register(c *gin.Context) {
	params := requests.RegisterUserRequest{}
	response := app.NewResponse(c)
	valid, err := app.BindAndValidation(c, &params)

	if !valid {
		errorResponse := errors.InvalidParams.WithDetails(err...)
		response.MakeErrorResponse(errorResponse)
		return
	}

	s := service.New(c.Request.Context())
	user, userErr := s.CreateUser(&params)

	if userErr != nil {
		response.MakeErrorResponse(errors.CreateUserFail)
		return
	}

	response.MakeResponse(user.ID)
}

func (u User) Login(c *gin.Context) {

}