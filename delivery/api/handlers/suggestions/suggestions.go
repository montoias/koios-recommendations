package suggestions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/montoias/koios-recommendations/infrastructure/tmdb"
)

// Ping allows to perform health check
func Suggestions(client *tmdb.Client) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		ginContext.JSON(http.StatusOK, client.Test())
	}
}
