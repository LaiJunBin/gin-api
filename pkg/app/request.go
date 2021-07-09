package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidation(c *gin.Context, v interface{}) (bool, []string) {
	var errors []string

	if err := c.ShouldBind(v); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}
		return false, errors
	}

	return true, nil
}