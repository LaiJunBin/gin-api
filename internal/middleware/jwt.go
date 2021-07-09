package middleware

import (
	"github.com/LaiJunBin/gin-api/pkg/app"
	"github.com/LaiJunBin/gin-api/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var errResponse *errors.Error
		authorization := ctx.GetHeader("Authorization")
		splitAuthorization := strings.SplitN(authorization, "Bearer ", 2)
		if len(splitAuthorization) != 2{
			errResponse = errors.InvalidParams
		} else {
			token := strings.TrimSpace(splitAuthorization[1])
			if token == "" {
				errResponse = errors.InvalidParams
			} else {
				claims, err := app.ParseToken(token)
				if err != nil {
					switch err.(*jwt.ValidationError).Errors {
					case jwt.ValidationErrorExpired:
						errResponse = errors.UnauthorizedTokenTimeout
						break
					default:
						errResponse = errors.UnauthorizedTokenError
						break
					}
				}

				ctx.Set("claims", claims)
			}
		}


		if errResponse != nil {
			response := app.NewResponse(ctx)
			response.MakeErrorResponse(errResponse)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
