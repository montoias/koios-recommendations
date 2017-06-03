package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/montoias/koios-recommendations/delivery/api/payload"
)

// Ping allows to perform health check
func Ping() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		ginContext.JSON(http.StatusOK, payload.NewResponse("pong"))
	}
}
