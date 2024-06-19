package routers

import (
	"github.com/Maniexie/crud-fiber-postgres/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterApp(c *fiber.App) {
	c.Get("/mahasiswa", controllers.MahasiswaControllerShow)
	c.Post("/mahasiswa/create", controllers.MahasiswaControllerCreate)
	c.Put("/mahasiswa/:id", controllers.MahasiswaControllerUpdate)
	c.Get("/mahasiswa/:id", controllers.MahasiswaControllerFind)
	c.Delete("/mahasiswa/:id", controllers.MahasiswaControllerDelete)
}
