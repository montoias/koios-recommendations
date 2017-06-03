package api

import (
	"github.com/gin-gonic/gin"
	"github.com/montoias/koios-recommendations/config"
)

// SetupDependencies builds the dependency graph based on the config
func SetupDependencies(config *config.MainConfig) *Dependencies {
	gin.SetMode(config.GetString("app.mode"))

	origins := config.GetString("app.cross_origin_domain")
	defaultTimeout := config.GetInt("app.default_timeout")

	dependencies := NewDependencyGraph()
	dependencies.CrossOriginDomains = origins
	dependencies.DefaultTimeout = defaultTimeout
	dependencies.DefineAll()
	dependencies.ResolveAll()

	return dependencies
}
