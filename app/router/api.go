package router

import (
	"go_plate/app/module/account"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App           fiber.Router
	AccountRouter *account.Router
}

func NewRouter(fiber *fiber.App, accountRouter *account.Router) *Router {
	return &Router{
		App:           fiber,
		AccountRouter: accountRouter,
	}
}

// Register routes
func (r *Router) Register() {
	// Test Routes
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Register routes of modules
	r.AccountRouter.RegisterRoutes()
}
