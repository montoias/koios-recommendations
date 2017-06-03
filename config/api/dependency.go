package api

import (
	"github.com/karlkfi/inject"
	app "github.com/montoias/koios-recommendations/delivery/api/app"
)

// NewDependencyGraph returns an initialized Dependency Graph.
// This is the common dependency graph for all environments. To inject specific types dependant of the environment
// you should create a new graph that overrides the defaults behaviour.
func NewDependencyGraph() *Dependencies {
	graph := inject.NewGraph()

	return &Dependencies{
		graph: graph,
	}
}

// Dependencies delivers injection providers
type Dependencies struct {
	graph inject.Graph

	// App
	App *app.App

	// configs

	// http configs
	CrossOriginDomains string
	DefaultTimeout     int
}

// ResolveAll resolves all dependencies and initializes the structure
func (dependencies *Dependencies) ResolveAll() {
	dependencies.graph.ResolveAll()
}

// DefineAll defines how all structures should be created
func (dependencies *Dependencies) DefineAll() {
	// app
	dependencies.graph.Define(
		&dependencies.App,
		inject.NewProvider(
			app.New,
			&dependencies.CrossOriginDomains,
		),
	)
}