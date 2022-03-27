package routes

import(
	"github.com/gominima/minima"
	ctrl "github.com/hexacry/venebook/controllers"
)

func V1A() *minima.Group {
	var V1G *minima.Group = minima.NewGroup("/api/v1")
	V1G.Get("/", ctrl.Home)

	return V1G
}