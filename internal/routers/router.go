package routers

import (
	"github.com/LaiJunBin/gin-api/internal/routers/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	user := api.NewUser()
	router.POST("/register", user.Register)
	router.POST("/login", user.Login)

	return router
}