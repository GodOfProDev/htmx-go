package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"htmx-go/internal/handlers"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Listen() {
	h := handlers.NewHandlers()

	engine := html.New("./views/", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("user/index", fiber.Map{})
	})

	app.Get("/ping", h.HandlePing)

	app.Listen("localhost:8080")
}
