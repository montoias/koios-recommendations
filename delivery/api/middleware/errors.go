package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/montoias/koios-recommendations/delivery/api/payload"
)

// Errors handles errors
func Errors() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		ginContext.Next() // execute all the handlers

		// At this point, all the handlers finished. Let's read last error.
		errorToPrint := ginContext.Errors.Last()
		if errorToPrint != nil {
			ginContext.JSON(-1, payload.NewError(errorToPrint.Error()))
		}
	}
}
