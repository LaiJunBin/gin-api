package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func BindAndValidation(c *gin.Context, v interface{}) (bool, []string) {
	var errors []string

	if err := c.ShouldBind(v); err != nil {
		translatorVal := c.Value("translator")
		translator, _ := translatorVal.(ut.Translator)
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errors
		}

		for _, value := range validationErrors.Translate(translator) {
			errors = append(errors, value)
		}

		return false, errors
	}

	return true, nil
}