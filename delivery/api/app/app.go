package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pingHandler "github.com/montoias/koios-recommendations/delivery/api/handlers/ping"
	"github.com/montoias/koios-recommendations/delivery/api/middleware"
)

// App is the API for the Microservice
type App struct {
	router *gin.Engine
}

// New returns new instance of the MicroService App
func New(crossOriginDomains string) *App {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.Errors())
	router.Use(middleware.Cors(crossOriginDomains))

	// Ping-Pong endpoint
	router.GET("/ping", pingHandler.Ping())

	return &App{
		router: router,
	}
}

// Run starts this Microservice API server
func (app *App) Run(addr ...string) error {
	return app.router.Run(addr...)
}

// ServeHTTP calls gin.ServeHTTP(w, req)
func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	app.router.ServeHTTP(w, req)
}
