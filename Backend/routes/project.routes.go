package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"mongoCrud/auth"
	m "mongoCrud/models"
	service "mongoCrud/services"
)

func ProjectsRoute(app *fiber.App) {

	project := app.Group("/api/projects")

	//crear proyecto
	project.Post("/new", auth.JWTProtected(), func(c *fiber.Ctx) error {
		p := new(m.Proyecto)
		if err := c.BodyParser(p); err != nil {
			fmt.Println(err)
			return fiber.ErrBadRequest
		}
		u, err := auth.ExtractTokenMetadata(c)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		fmt.Println("el nombre del proyecto", p.Name)

		if err := service.CreateProject(p, u.Email); err != nil {

			fmt.Println(err)
			return fiber.ErrBadRequest
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "correcto"})
	})

	// get proyectos que yo cree
	project.Get("/myprojects", auth.JWTProtected(), func(c *fiber.Ctx) error {

		u, err := auth.ExtractTokenMetadata(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		ps, err := service.GetMyProjects(u.Email)
		fmt.Println(ps)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(ps)
	})

	//get proyectos en los que me han contratodo
	project.Get("/projects", auth.JWTProtected(), func(c *fiber.Ctx) error {
		u, err := auth.ExtractTokenMetadata(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		ps, err := service.GetProjectsWorkingOn(u.Email)
		fmt.Println(ps)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(ps)

	})

	//agregar un trabajador al proyecto
	project.Post("/add-worker", auth.JWTProtected(), func(c *fiber.Ctx) error {
		// body :=struct {
		// 	projectId string
		// 	email string

		// }

		// if err := c.BodyParser(p); err != nil {
		// 	fmt.Println(err)
		// 	return fiber.ErrBadRequest
		// }
		return nil
	})

}
