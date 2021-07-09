package api

import (
	"github.com/LaiJunBin/gin-api/internal/requests"
	"github.com/LaiJunBin/gin-api/internal/service"
	"github.com/LaiJunBin/gin-api/pkg/app"
	"github.com/LaiJunBin/gin-api/pkg/converts"
	"github.com/LaiJunBin/gin-api/pkg/errors"
	"github.com/LaiJunBin/gin-api/pkg/utils"
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
	params := requests.LoginUserRequest{}
	response := app.NewResponse(c)
	valid, err := app.BindAndValidation(c, &params)

	if !valid {
		errorResponse := errors.InvalidParams.WithDetails(err...)
		response.MakeErrorResponse(errorResponse)
		return
	}

	s := service.New(c.Request.Context())
	user, userErr := s.GetUserByUsername(params.Username)

	if userErr != nil {
		response.MakeErrorResponse(errors.LoginFail)
		return
	}

	if !utils.CheckPasswordHash(params.Password, user.Password) {
		response.MakeErrorResponse(errors.LoginFail)
		return
	}

	token, tokenErr := app.GenerateToken(user.Username)
	if tokenErr != nil {
		response.MakeErrorResponse(errors.UnauthorizedTokenGenerate)
		return
	}

	response.MakeResponse(token)
}

func (u User) Get(c *gin.Context) {
	id := converts.StrTo(c.Param("id")).MustUInt()
	params := requests.GetUserRequest {
		ID: id,
	}
	response := app.NewResponse(c)
	valid, err := app.BindAndValidation(c, &params)

	if !valid {
		errorResponse := errors.InvalidParams.WithDetails(err...)
		response.MakeErrorResponse(errorResponse)
		return
	}

	s := service.New(c.Request.Context())
	user, userErr := s.GetUserById(id)

	if userErr != nil {
		response.MakeErrorResponse(errors.NotFound)
		return
	}

	response.MakeResponse(user)
}

func (u User) Update(c *gin.Context) {
	id := converts.StrTo(c.Param("id")).MustUInt()
	params := requests.UpdateUserRequest {
		ID: id,
	}
	response := app.NewResponse(c)
	valid, err := app.BindAndValidation(c, &params)

	if !valid {
		errorResponse := errors.InvalidParams.WithDetails(err...)
		response.MakeErrorResponse(errorResponse)
		return
	}

	s := service.New(c.Request.Context())
	updateErr := s.UpdateUser(id, &params)

	if updateErr != nil {
		response.MakeErrorResponse(errors.UpdateUserFail)
		return
	}

	response.MakeResponse("")
}

func (u User) Delete(c *gin.Context) {
	s := service.New(c.Request.Context())
	id := converts.StrTo(c.Param("id")).MustUInt()
	claims := c.Keys["claims"].(*app.Claims)

	response := app.NewResponse(c)
	if !s.CheckPermission(id, claims) {
		response.MakeErrorResponse(errors.StatusUnauthorized)
		return
	}

	err := s.DeleteUser(id)
	if err != nil {
		response.MakeErrorResponse(errors.DeleteUserFail)
		return
	}
	response.MakeResponse("")
}
