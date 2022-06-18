package config

import (
	"github.com/kataras/iris/v12"
	"github.com/B3zaleel/fractage/src/controllers"
)

// Adds all routes to the given iris application.
func AddRoutes(app *iris.Application) {
	app.Get("/cantor-set", controllers.GetCantorSet)
	app.Get("/sierpinski-carpet", controllers.GetSierpinskiCarpet)
}
