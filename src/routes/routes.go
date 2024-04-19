package routes

import (
	"github.com/gofiber/fiber/v2"

	"gofiber-mongodb/src/controllers"
)

func Router(c *fiber.App) {

	c.Get("/student", controllers.FindData )
	c.Post("/create-student", controllers.CreateData )
	c.Put("/update-student/:name", controllers.UpdateData )
	c.Delete("/delete-student/:name", controllers.DeleteData )
}
