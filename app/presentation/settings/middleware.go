package settings

import (
	"errors"

	"github.com/gin-gonic/gin"

	errDomain "github.com/yumekumo/sauna-shop/domain/error"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case *errDomain.Error:
				if errors.Is(err, errDomain.NotFoundErr) {
					ReturnNotFound(c, e)
				}
				ReturnStatusBadRequest(c, e)
			default:
				ReturnStatusInternalServerError(c, e)
			}
		}
	}
}
