package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		v := validator.New()
		c.Set("validate", v)
		c.Next()
	}
}
