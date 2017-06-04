package suggestions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/montoias/koios-recommendations/delivery/api/payload"
	"github.com/montoias/koios-recommendations/delivery/api/payload/inbound/users"
	"github.com/montoias/koios-recommendations/delivery/api/payload/outbound/movies"
)

// Suggestions allows to request suggestions
func Suggestions(interactor SuggestionsInteractor) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		body, err := ioutil.ReadAll(ginContext.Request.Body)
		if err != nil {
			ginContext.JSON(http.StatusBadRequest, payload.Error{Message: "unable to read request body"})

			return
		}

		var request users.UsersRequest
		if err = json.Unmarshal(body, &request); err != nil {
			ginContext.JSON(http.StatusBadRequest, payload.Error{Message: "unable to unmarshal request"})

			return
		}

		usersDto := users.ToDto(request)
		if err != nil {
			ginContext.JSON(http.StatusInternalServerError, payload.Error{Message: "unable to process request"})

			return
		}

		suggestedMovies, err := interactor.CreateSuggestions(usersDto)
		if err != nil {
			ginContext.JSON(http.StatusInternalServerError, payload.Error{Message: "unable to create suggestions"})

			return
		}

		moviesResult := movies.ToSchema(suggestedMovies)
		result, err := json.Marshal(moviesResult)
		if err != nil {
			ginContext.JSON(http.StatusInternalServerError, payload.Error{Message: "unable to provide suggestions"})

			return
		}

		ginContext.JSON(http.StatusOK, string(result))
	}
}
