package account

import (
	"go_plate/app/module/account/biz"
	"go_plate/app/module/account/data"
	"go_plate/app/module/account/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Router struct {
	App fiber.Router
	s   *service.Service
}

// Register Account Module
var NewAccountModule = fx.Options(
	fx.Provide(data.NewData),
	fx.Provide(service.NewService),
	fx.Provide(biz.NewBiz),

	// Register Router
	fx.Provide(NewRouter),
)

// Router methods
func NewRouter(fiber *fiber.App, svc *service.Service) *Router {
	return &Router{
		App: fiber,
		s:   svc,
	}
}

func (r *Router) RegisterRoutes() {
	// Define routes

}
