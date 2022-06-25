package config

import (
	"github.com/kataras/iris/v12"
	"github.com/B3zaleel/fractage/src/controllers"
)

// Adds all routes to the given iris application.
func AddRoutes(app *iris.Application) {
	app.Get("/palette", controllers.GetPalette)

	app.Get("/cantor-dust", controllers.GetCantorDust)
	app.Get("/cantor-set", controllers.GetCantorSet)
	app.Get("/hopalong", controllers.GetHopalong)
	app.Get("/julia-set", controllers.GetJuliaSet)
	app.Get("/mandelbrot-set", controllers.GetMandelbrotSet)
	app.Get("/sierpinski-carpet", controllers.GetSierpinskiCarpet)
	app.Get("/sierpinski-triangle", controllers.GetSierpinskiTriangle)
}
