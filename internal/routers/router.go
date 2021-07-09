package routers

import (
	"github.com/LaiJunBin/gin-api/internal/middleware"
	"github.com/LaiJunBin/gin-api/internal/routers/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Translations())

	user := api.NewUser()
	router.POST("/register", user.Register)
	router.POST("/login", user.Login)
	router.GET("/users/:id", user.Get)

	authRouter := router.Group("/")
	authRouter.Use(middleware.JWT())
	authRouter.PATCH("/users/:id", user.Update)
	authRouter.DELETE("/users/:id", user.Delete)

	return router
}