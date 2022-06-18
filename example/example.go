package example

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ztcollazo/vibe"
)

type AppController struct {
	vibe.Controller[*AppController]
	Hello string
}

func (a *AppController) Setup(c *fiber.Ctx) {
	a.Hello = "Hello World!"
}

func (a *AppController) Routes() map[string]vibe.Route {
	return map[string]vibe.Route{
		"Custom": {
			Path: "/custom",
		},
	}
}

func (a *AppController) Index(c *fiber.Ctx) error {
	return c.SendString(a.SomeFuncThatIsNotARoute())
}

func (a *AppController) New(c *fiber.Ctx) error {
	return c.SendString("This is at GET /new")
}

func (a *AppController) Create(c *fiber.Ctx) error {
	return c.SendString("This is at POST /")
}

func (a *AppController) Show(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("This is at GET /:id with id of %s", c.Params("id")))
}

func (a *AppController) Edit(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("This is at GET /:id/edit with id of %s", c.Params("id")))
}

func (a *AppController) Update(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("This is at POST /:id with id of %s", c.Params("id")))
}

func (a *AppController) Destroy(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("This is at DELETE /:id with id of %s", c.Params("id")))
}

func (a *AppController) Custom(c *fiber.Ctx) error {
	return c.SendString("This is a custom route at GET /custom")
}

func (a *AppController) SomeFuncThatIsNotARoute() string {
	return a.Hello
}

func RunApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Route("/", vibe.CreateController(&AppController{}))
	return app
}
